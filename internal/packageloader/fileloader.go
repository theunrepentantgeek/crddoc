package packageloader

import (
	"go/token"
	"os"
	"path/filepath"

	"github.com/dave/dst"
	"github.com/dave/dst/decorator"
	"github.com/go-logr/logr"
	"github.com/pkg/errors"
	"golang.org/x/exp/maps"

	"github.com/theunrepentantgeek/crddoc/internal/model"
	"github.com/theunrepentantgeek/crddoc/internal/typefilter"
)

type FileLoader struct {
	path             string
	name             string
	importReferences model.ImportReferenceSet
	typeFilters      *typefilter.List
	resources        map[string]*model.Resource
	objects          map[string]*model.Object
	enums            map[string]*model.Enum
	values           map[string][]*model.EnumValue
	packageMarkers   *model.PackageMarkers
	log              logr.Logger
}

func NewFileLoader(
	path string,
	log logr.Logger,
	typeFilters *typefilter.List,
) *FileLoader {
	return &FileLoader{
		name:             filepath.Base(path),
		path:             path,
		importReferences: make(model.ImportReferenceSet),
		typeFilters:      typeFilters,
		resources:        make(map[string]*model.Resource),
		objects:          make(map[string]*model.Object),
		enums:            make(map[string]*model.Enum),
		values:           make(map[string][]*model.EnumValue),
		packageMarkers:   model.NewPackageMarkers(),
		log:              log,
	}
}

func (loader *FileLoader) Load() error {
	loader.log.V(1).Info(
		"Loading source file",
		"file", loader.name)

	file, err := loader.parseFile()
	if err != nil {
		return err
	}

	if err = loader.parseMetadata(file.Decs); err != nil {
		return errors.Wrap(err, "failed to parse metadata")
	}

	loader.parseImports(file.Imports)

	if err = loader.parseDecls(file.Decls); err != nil {
		return errors.Wrap(err, "failed to parse declarations")
	}

	loader.discoverResources()
	loader.assembleEnumerations()

	return nil
}

func (loader *FileLoader) parseDecl(
	gd *dst.GenDecl,
) error {
	switch gd.Tok {
	case token.TYPE:
		comments := gd.Decs.Start.All()
		if err := loader.parseTypes(gd.Specs, comments); err != nil {
			return errors.Wrap(err, "parsing types")
		}
	case token.CONST:
		loader.parseConstants(gd.Specs)
	default:
	}

	return nil
}

// parseTypes iterates through a sequence of dst.Spec declarations trying to parse
// objects and enums.
func (loader *FileLoader) parseTypes(
	specs []dst.Spec,
	comments []string,
) error {
	description, markers := model.ParseComments(comments)
	if err := loader.packageMarkers.Update(markers); err != nil {
		return errors.Wrap(err, "parsing comments for markers")
	}

	// Parse type declarations for objects and enums
	for _, spec := range specs {
		// Try to create an object from this declaration
		if obj, ok := model.TryNewObject(spec, description, loader.importReferences); ok {
			loader.objects[obj.ID()] = obj
		}

		// Try to create an enum from this declaration
		if enum, ok := model.TryNewEnum(spec, description); ok {
			loader.enums[enum.ID()] = enum
		}
	}

	return nil
}

func (loader *FileLoader) parseConstants(
	specs []dst.Spec,
) {
	for _, spec := range specs {
		// Parse constant declarations for enums
		if enumValue, ok := model.TryNewEnumValue(spec); ok {
			kind := enumValue.Kind()
			loader.values[kind] = append(loader.values[kind], enumValue)
		}
	}
}

func (loader *FileLoader) parseDecls(decls []dst.Decl) error {
	for _, decl := range decls {
		if gd, ok := decl.(*dst.GenDecl); ok {
			if err := loader.parseDecl(gd); err != nil {
				return err
			}
		}
	}

	return nil
}

// discoverResources iterates through the objects we've already loaded and identifies any
// that represent resources.
func (loader *FileLoader) discoverResources() {
	// Find all the objects that are actually resources
	for _, obj := range maps.Values(loader.objects) {
		// Try to create a resource from this object
		if resource, ok := model.TryNewResource(obj); ok {
			loader.resources[resource.ID()] = resource
			delete(loader.objects, resource.ID())
		}
	}
}

// assembleEnumerations iterates through all the enums we've found and adds any constant values
// of that type.
func (loader *FileLoader) assembleEnumerations() {
	// Add the values to each enum
	for n, e := range loader.enums {
		if values, ok := loader.values[n]; ok {
			for _, v := range values {
				e.AddValue(v)
			}
		}
	}
}

func (loader *FileLoader) parseFile() (file *dst.File, failure error) {
	defer func() {
		if err := recover(); err != nil {
			failure = errors.Errorf("panic reading %s: %v", loader.path, err)
		}
	}()

	loader.log.V(1).Info(
		"Loading source file",
		"file", loader.name)

	// Create a reader for the file
	reader, err := os.Open(loader.path)
	if err != nil {
		return file, errors.Wrapf(err, "failed to open file %s", loader.path)
	}

	defer reader.Close()

	file, err = decorator.Parse(reader)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to parse file %s", loader.path)
	}

	return file, nil
}

func (loader *FileLoader) Declarations() []model.Declaration {
	expectedDeclarations := len(loader.resources) + len(loader.objects) + len(loader.enums)
	result := make([]model.Declaration, 0, expectedDeclarations)

	result = append(result, filterDeclarations(loader.resources, loader.typeFilters)...)
	result = append(result, filterDeclarations(loader.objects, loader.typeFilters)...)
	result = append(result, filterDeclarations(loader.enums, loader.typeFilters)...)

	return result
}

func filterDeclarations[D model.Declaration](
	decls map[string]D,
	typeFilters *typefilter.List,
) []model.Declaration {
	result := make([]model.Declaration, 0, len(decls))

	for n, d := range decls {
		if typeFilters.Filter(n) == typefilter.Included {
			result = append(result, d)
		}
	}

	return result
}

func (loader *FileLoader) PackageMarkers() *model.PackageMarkers {
	return loader.packageMarkers
}

func (loader *FileLoader) parseMetadata(decs dst.FileDecorations) error {
	leadingMarkers := model.NewMarkers(decs.Start...)
	if err := loader.packageMarkers.Update(leadingMarkers); err != nil {
		return errors.Wrap(err, "failed to update package markers")
	}

	trailingMarkers := model.NewMarkers(decs.End...)
	if err := loader.packageMarkers.Update(trailingMarkers); err != nil {
		return errors.Wrap(err, "failed to update package markers")
	}

	return nil
}

func (loader *FileLoader) parseImports(imports []*dst.ImportSpec) {
	for _, imp := range imports {
		if ref, ok := model.TryNewImportReference(imp); ok {
			loader.importReferences[ref.Name()] = ref
		}
	}
}
