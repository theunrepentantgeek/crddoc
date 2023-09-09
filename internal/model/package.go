package model

import (
	"go/token"
	"os"
	"path/filepath"
	"strings"
	"sync"

	"github.com/dave/dst"
	"github.com/dave/dst/decorator"
	"github.com/go-logr/logr"
	"github.com/pkg/errors"
	"github.com/theunrepentantgeek/crddoc/internal/config"
	"golang.org/x/exp/maps"
	"golang.org/x/exp/slices"
	"golang.org/x/sync/errgroup"
)

// Package is a struct containing all of the declarations found in a package directory
type Package struct {
	name         string
	cfg          *config.Config
	declarations map[string]Declaration // Dictionary of all the objects in the package, keyed by name
	log          logr.Logger
	lock         sync.Mutex
}

type Order string

const (
	OrderAlphabetical = "alphabetical"
)

func NewPackage(cfg *config.Config, log logr.Logger) *Package {
	return &Package{
		cfg:          cfg,
		declarations: make(map[string]Declaration),
		log:          log,
	}
}

func (p *Package) Name() string {
	return p.name
}

// LoadDirectory scans a directory for Go files and loads them into the Package
func (p *Package) LoadDirectory(folder string) error {
	fdr, pkg := filepath.Split(folder)

	p.log.Info(
		"Loading package",
		"name", pkg,
		"folder", fdr)

	p.name = pkg

	files, err := os.ReadDir(folder)
	if err != nil {
		return errors.Wrapf(err, "failed to read directory %s", folder)
	}

	var eg errgroup.Group
	for _, f := range files {
		f := f

		if f.IsDir() {
			continue
		}

		if filepath.Ext(f.Name()) != ".go" {
			continue
		}

		eg.Go(func() error {
			var path = filepath.Join(folder, f.Name())
			return p.LoadFile(path)
		})
	}

	if err := eg.Wait(); err != nil {
		return errors.Wrapf(err, "failed to load directory %s", folder)
	}

	p.catalogCrossReferences()

	return nil
}

func (p *Package) LoadFile(path string) (failure error) {
	defer func() {
		if err := recover(); err != nil {
			failure = errors.Errorf("panic reading %s: %v", path, err)
		}
	}()

	_, f := filepath.Split(path)
	p.log.V(1).Info(
		"Loading source file",
		"file", f)

	// Create a reader for the file
	reader, err := os.Open(path)
	if err != nil {
		return errors.Wrapf(err, "failed to open file %s", path)
	}

	defer reader.Close()

	file, err := decorator.Parse(reader)
	if err != nil {
		return errors.Wrapf(err, "failed to parse file %s", path)
	}

	// Find declarations of interest
	objects := p.findObjects(file.Decls)
	resources := p.findResources(objects)
	enums := p.findEnums(file.Decls)

	// Add them to the package
	addDeclarations(p, resources)
	addDeclarations(p, objects)
	addDeclarations(p, enums)

	return nil
}

func (p *Package) Declarations(order Order) []Declaration {
	p.lock.Lock()
	defer p.lock.Unlock()

	result := maps.Values(p.declarations)

	// Sort the declarations as specified
	switch order {
	case OrderAlphabetical:
		// Sort the objects alphabetically
		slices.SortFunc(result, alphabeticalObjectComparison)
	}

	return result
}

// Declaration returns the declaration with the given name, if found.
func (p *Package) Declaration(name string) (Declaration, bool) {
	p.lock.Lock()
	defer p.lock.Unlock()

	dec, ok := p.declarations[strings.ToLower(name)]
	if !ok {
		return nil, false
	}

	return dec, ok
}

// Object returns the object with the given name, if there is one
func (p *Package) Object(name string) (*Object, bool) {
	dec, ok := p.Declaration(name)
	if !ok {
		return nil, false
	}

	obj, ok := dec.(*Object)
	return obj, ok
}

// findObjects scans the declarations in a file and returns a slice of objects
func (p *Package) findObjects(decls []dst.Decl) map[string]*Object {
	result := make(map[string]*Object, len(decls))

	for _, decl := range decls {
		// Check for a GenDecl containing a TYPE
		gd, ok := decl.(*dst.GenDecl)
		if !ok || gd.Tok != token.TYPE {
			continue
		}

		comments := gd.Decs.Start.All()

		// Iterate over the specs in the GenDecl and try to create an object from each
		for _, spec := range gd.Specs {

			if obj, ok := TryNewObject(spec, comments); ok {
				result[obj.Id()] = obj
			}
		}
	}

	return result
}

func (p *Package) findResources(objects map[string]*Object) map[string]*Resource {
	result := make(map[string]*Resource)

	// Find all the objects that are actually resources
	for _, obj := range objects {
		if resource, ok := TryNewResource(obj); ok {
			result[resource.Id()] = resource
		}
	}

	// Remove those objects so we don't have any name collisions
	for name := range result {
		delete(objects, name)
	}

	return result
}

// findEnums scans the declarations in a file and returns a slice of enumerations
func (p *Package) findEnums(decls []dst.Decl) map[string]*Enum {

	// Collect Enum Types
	enums := make(map[string]*Enum)
	for _, decl := range decls {
		// Check for a GenDecl containing a TYPE
		gd, ok := decl.(*dst.GenDecl)
		if !ok || gd.Tok != token.TYPE {
			continue
		}

		comments := gd.Decs.Start.All()

		// Iterate over the specs in the GenDecl and try to create an enum from each
		for _, spec := range gd.Specs {
			if enum, ok := TryNewEnum(spec, comments); ok {
				enums[enum.Id()] = enum
			}
		}
	}

	// Now that we have all the enums, we can scan the declarations again and add the
	// enum values to the appropriate enum
	for _, decl := range decls {
		// Check for a GenDecl containing a CONST
		gd, ok := decl.(*dst.GenDecl)
		if !ok || gd.Tok != token.CONST {
			continue
		}

		// Iterate over the specs in the GenDecl and try to create an enum value
		for _, spec := range gd.Specs {
			if enumValue, ok := TryNewEnumValue(spec); ok {
				if enum, ok := enums[enumValue.Kind()]; ok {
					enum.AddValue(enumValue)
				}
			}
		}
	}

	return enums
}

func (p *Package) catalogCrossReferences() {
	p.lock.Lock()
	defer p.lock.Unlock()

	usages := p.indexUsage()
	for name, usage := range usages {
		if obj, ok := p.declarations[name]; ok {
			slices.SortFunc(usage, ComparePropertyReferences)
			obj.SetUsage(usage)
		}
	}
}

func (p *Package) indexUsage() map[string][]PropertyReference {
	result := make(map[string][]PropertyReference)
	for name, dec := range p.declarations {
		// Index references from an object
		if host, ok := dec.(PropertyContainer); ok {
			for _, prop := range host.Properties() {
				id := prop.Type.Id()
				if _, ok := p.declarations[id]; ok {
					ref := NewPropertyReference(name, prop.Name)
					result[id] = append(result[id], ref)
				}
			}
		}
	}

	return result
}

func alphabeticalObjectComparison(left Declaration, right Declaration) int {
	if left.Name() < right.Name() {
		return -1
	}

	if left.Name() > right.Name() {
		return 1
	}

	return 0
}

// addDeclarations adds more declarations to the package
func addDeclarations[D Declaration](p *Package, declarations map[string]D) {
	p.lock.Lock()
	defer p.lock.Unlock()

	for name, decl := range declarations {
		// Skip excluded declarations
		if p.cfg.Filter(name) == config.FilterResultExclude {
			continue
		}

		//TODO: Check for name collisions
		// (Should never happen - BUT if it does, we need to know)
		p.declarations[name] = decl
	}
}
