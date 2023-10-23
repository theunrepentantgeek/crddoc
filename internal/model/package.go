package model

import (
	"go/token"
	"os"
	"path/filepath"
	"strings"
	"sync"

	"github.com/theunrepentantgeek/crddoc/internal/config"
	"github.com/theunrepentantgeek/crddoc/internal/typefilter"

	"github.com/dave/dst"
	"github.com/dave/dst/decorator"
	"github.com/go-logr/logr"
	"github.com/pkg/errors"
	"golang.org/x/exp/maps"
	"golang.org/x/exp/slices"
	"golang.org/x/sync/errgroup"
)

// Package is a struct containing all of the declarations found in a package directory
type Package struct {
	name         string
	cfg          *config.Config
	typeFilters  *typefilter.TypeFilterList
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
		typeFilters:  typefilter.New(cfg),
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

// Group returns the group of the package
func (p *Package) Group() string {
	return p.metadata.Group
}

// Version returns the version of the package
func (p *Package) Version() string {
	return p.metadata.Version
}

func (p *Package) catalogCrossReferences() {
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
	for _, dec := range p.declarations {
		// Index references from an object
		if host, ok := dec.(PropertyContainer); ok {
			for _, prop := range host.Properties() {
				id := prop.Type.Id()
				if _, ok := p.declarations[id]; ok {
					ref := NewPropertyReference(dec.Name(), dec.Id(), prop.Name)
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
