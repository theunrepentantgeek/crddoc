package model

import (
	"go/token"
	"os"
	"path/filepath"
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

	// Iterate over the file's declarations and add them to this package
	objects := p.findObjects(file.Decls)
	enums := p.findEnums(file.Decls)

	p.addDeclarations(objects)
	p.addDeclarations(enums)

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

// Object returns the object with the given name, if found.
func (p *Package) Object(name string) (*Object, bool) {
	p.lock.Lock()
	defer p.lock.Unlock()

	dec, ok := p.declarations[name]
	if !ok {
		return nil, false
	}

	obj, ok := dec.(*Object)
	return obj, ok
}

// findObjects scans the declarations in a file and returns a slice of objects
func (p *Package) findObjects(decls []dst.Decl) []Declaration {
	var result []Declaration

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
				result = append(result, obj)
			}
		}
	}

	return result
}

// findEnums scans the declarations in a file and returns a slice of enumerations
func (p *Package) findEnums(decls []dst.Decl) []Declaration {

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
				enums[enum.Name()] = enum
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

	var result []Declaration
	for _, enum := range enums {
		result = append(result, enum)
	}

	return result
}

// addDeclarations adds more declarations to the package
func (p *Package) addDeclarations(declarations []Declaration) {
	p.lock.Lock()
	defer p.lock.Unlock()

	for _, dec := range declarations {
		// Skip excluded declarations
		if p.cfg.Filter(dec.Name()) == config.FilterResultExclude {
			continue
		}

		//TODO: Check for name collisions
		// (Should never happen - BUT if it does, we need to know)
		p.declarations[dec.Name()] = dec
	}
}

func (p *Package) catalogCrossReferences() {
	p.lock.Lock()
	defer p.lock.Unlock()

	// Index all usage
	usages := make(map[string][]PropertyReference)
	for _, dec := range p.declarations {
		obj, ok := dec.(*Object)
		if !ok {
			continue
		}

		for _, prop := range obj.properties {
			if t := p.CreateIdFor(prop.Type()); t != "" {
				ref := NewPropertyReference(obj, prop.Name())
				usages[t] = append(usages[t], ref)
			}
		}
	}

	// Update all objects
	for name, usage := range usages {
		if obj, ok := p.declarations[name]; ok {
			slices.SortFunc(usage, ComparePropertyReferences)
			obj.SetUsage(usage)
		}
	}
}

// asId renders an ID from a type expression, for linking within the documentation.
// Returns an empty string if the object does not exist in the package.
func (p *Package) CreateIdFor(expr dst.Expr) string {
	switch t := expr.(type) {
	case *dst.Ident:
		if _, ok := p.declarations[t.Name]; !ok {
			return ""
		}

		return t.Name
	case *dst.StarExpr:
		return p.CreateIdFor(t.X)
	case *dst.ArrayType:
		return p.CreateIdFor(t.Elt)
	case *dst.MapType:
		// TODO: What should we do if both Key and Value are custom types?
		return p.CreateIdFor(t.Key) + p.CreateIdFor(t.Value)
	default:
		return ""
	}
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
