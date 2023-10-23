package packageloader

import (
	"go/token"
	"os"
	"path/filepath"

	"github.com/dave/dst"
	"github.com/dave/dst/decorator"
	"github.com/go-logr/logr"
	"github.com/pkg/errors"
	"github.com/theunrepentantgeek/crddoc/internal/model"
	"golang.org/x/exp/maps"
)

type FileLoader struct {
	path      string
	name      string
	resources map[string]*model.Resource
	objects   map[string]*model.Object
	enums     map[string]*model.Enum
	values    map[string][]*model.EnumValue
	log       logr.Logger
}

func NewFileLoader(path string, log logr.Logger) *FileLoader {
	return &FileLoader{
		name:      filepath.Base(path),
		path:      path,
		resources: make(map[string]*model.Resource),
		objects:   make(map[string]*model.Object),
		enums:     make(map[string]*model.Enum),
		values:    make(map[string][]*model.EnumValue),
		log:       log,
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

	for _, decl := range file.Decls {
		if gd, ok := decl.(*dst.GenDecl); ok {
			if gd.Tok == token.TYPE {

				comments := gd.Decs.Start.All()
				for _, spec := range gd.Specs {
					// Try to create an object from this declaration
					if obj, ok := model.TryNewObject(spec, comments); ok {
						loader.objects[obj.Id()] = obj
					}

					// Try to create an enum from this declaration
					if enum, ok := model.TryNewEnum(spec, comments); ok {
						loader.enums[enum.Id()] = enum
					}
				}
			}

			if gd.Tok == token.CONST {
				for _, spec := range gd.Specs {
					// Try to create a value from this declaration
					if enumValue, ok := model.TryNewEnumValue(spec); ok {
						loader.values[enumValue.Kind()] = append(loader.values[enumValue.Kind()], enumValue)
					}
				}
			}
		}
	}

	// Find all the objects that are actually resources
	for _, obj := range maps.Values(loader.objects) {
		// Try to create a resource from this object
		if resource, ok := model.TryNewResource(obj); ok {
			loader.resources[resource.Id()] = resource
			delete(loader.objects, resource.Id())
		}
	}

	// Add the values to each enum
	for n, e := range loader.enums {
		if values, ok := loader.values[n]; ok {
			for _, v := range values {
				e.AddValue(v)
			}
		}
	}

	return nil
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
	result := make([]model.Declaration, 0, len(loader.resources)+len(loader.objects)+len(loader.enums))
	for _, r := range loader.resources {
		result = append(result, r)
	}

	for _, o := range loader.objects {
		result = append(result, o)
	}

	for _, e := range loader.enums {
		result = append(result, e)
	}

	return result
}
