package packageloader

import (
	"go/token"
	"maps"
	"os"
	"path/filepath"

	"github.com/dave/dst"
	"github.com/dave/dst/decorator"
	"github.com/go-logr/logr"
	"github.com/pkg/errors"

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
	interfaces       map[string]*model.Interface
	values           map[string][]*model.EnumValue
	functions        map[string][]*model.Function         // Functions keyed by receiver type ID
	typeAssertions   map[string][]model.TypeAssertionInfo // Type assertions keyed by interface name
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
		interfaces:       make(map[string]*model.Interface),
		values:           make(map[string][]*model.EnumValue),
		functions:        make(map[string][]*model.Function),
		typeAssertions:   make(map[string][]model.TypeAssertionInfo),
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
	case token.VAR:
		loader.parseVarDeclarations(gd.Specs)
	default:
	}

	return nil
}

// parseTypes iterates through a sequence of dst.Spec declarations trying to parse
// objects, enums, and interfaces.
func (loader *FileLoader) parseTypes(
	specs []dst.Spec,
	comments []string,
) error {
	description, markers := model.ParseComments(comments)
	if err := loader.packageMarkers.Update(markers); err != nil {
		return errors.Wrap(err, "parsing comments for markers")
	}

	// Parse type declarations for objects, enums, and interfaces
	for _, spec := range specs {
		// Try to create an object from this declaration
		if obj, ok := model.TryNewObject(spec, description, loader.importReferences); ok {
			loader.objects[obj.ID()] = obj
		}

		// Try to create an enum from this declaration
		if enum, ok := model.TryNewEnum(spec, description); ok {
			loader.enums[enum.ID()] = enum
		}

		// Try to create an interface from this declaration
		if iface, ok := model.TryNewInterface(spec, description, loader.importReferences); ok {
			loader.interfaces[iface.ID()] = iface
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

// parseVarDeclarations parses variable declarations looking for type assertions
// of the form `var _ Interface = &Type{}` or `var _ Interface = Type{}`.
func (loader *FileLoader) parseVarDeclarations(specs []dst.Spec) {
	for _, spec := range specs {
		if assertion, ok := loader.tryParseTypeAssertion(spec); ok {
			loader.typeAssertions[assertion.InterfaceName] = append(
				loader.typeAssertions[assertion.InterfaceName], assertion)
		}
	}
}

// tryParseTypeAssertion tries to parse a type assertion from a var declaration.
// Returns the assertion and true if successful, empty assertion and false otherwise.
//
//nolint:revive,cyclop // Inherent complexity for parsing type assertions
func (loader *FileLoader) tryParseTypeAssertion(spec dst.Spec) (model.TypeAssertionInfo, bool) {
	valueSpec, ok := spec.(*dst.ValueSpec)
	if !ok {
		return model.TypeAssertionInfo{}, false
	}

	// Check for blank identifier (`_`)
	if len(valueSpec.Names) != 1 || valueSpec.Names[0].Name != "_" {
		return model.TypeAssertionInfo{}, false
	}

	// Get the interface type from the type annotation
	if valueSpec.Type == nil {
		return model.TypeAssertionInfo{}, false
	}

	interfaceName := extractTypeName(valueSpec.Type)
	if interfaceName == "" {
		return model.TypeAssertionInfo{}, false
	}

	// Get the implementing type from the value
	if len(valueSpec.Values) != 1 {
		return model.TypeAssertionInfo{}, false
	}

	typeName, isPointer := extractImplementingType(valueSpec.Values[0])
	if typeName == "" {
		return model.TypeAssertionInfo{}, false
	}

	return model.NewTypeAssertionInfo(interfaceName, typeName, isPointer), true
}

// extractTypeName extracts the type name from an expression.
func extractTypeName(expr dst.Expr) string {
	switch t := expr.(type) {
	case *dst.Ident:
		return t.Name
	case *dst.SelectorExpr:
		// For qualified types like pkg.Type, return just the type name
		return t.Sel.Name
	default:
		return ""
	}
}

// extractImplementingType extracts the implementing type name from a composite literal or address expression.
// Returns the type name and whether it's a pointer (address-of expression).
func extractImplementingType(expr dst.Expr) (string, bool) {
	switch t := expr.(type) {
	case *dst.UnaryExpr:
		// Handle &Type{} - address-of expression
		if compLit, ok := t.X.(*dst.CompositeLit); ok {
			return extractTypeName(compLit.Type), true
		}
	case *dst.CompositeLit:
		// Handle Type{} - value expression
		return extractTypeName(t.Type), false
	}

	return "", false
}

// parseFunc parses a function declaration and stores it for later attachment to objects.
func (loader *FileLoader) parseFunc(decl *dst.FuncDecl) {
	// Try to create a function from this declaration
	function, ok := model.TryNewFunction(decl)
	if !ok {
		return
	}

	// Get the base type name (without pointer) from the receiver
	receiverID := function.Receiver.ID()

	// Store the function keyed by receiver type for later attachment
	loader.functions[receiverID] = append(loader.functions[receiverID], function)
}

func (loader *FileLoader) parseDecls(decls []dst.Decl) error {
	for _, decl := range decls {
		if gd, ok := decl.(*dst.GenDecl); ok {
			if err := loader.parseDecl(gd); err != nil {
				return err
			}
		} else if fd, ok := decl.(*dst.FuncDecl); ok {
			loader.parseFunc(fd)
		}
	}

	return nil
}

// discoverResources iterates through the objects we've already loaded and identifies any
// that represent resources.
func (loader *FileLoader) discoverResources() {
	// Find all the objects that are actually resources
	for obj := range maps.Values(loader.objects) {
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

func (loader *FileLoader) Resources() []*model.Resource {
	return filterDeclarations(loader.resources, loader.typeFilters)
}

func (loader *FileLoader) Objects() []*model.Object {
	return filterDeclarations(loader.objects, loader.typeFilters)
}

func (loader *FileLoader) Enums() []*model.Enum {
	return filterDeclarations(loader.enums, loader.typeFilters)
}

func (loader *FileLoader) Interfaces() []*model.Interface {
	return filterDeclarations(loader.interfaces, loader.typeFilters)
}

func (loader *FileLoader) Functions() map[string][]*model.Function {
	return loader.functions
}

func (loader *FileLoader) TypeAssertions() map[string][]model.TypeAssertionInfo {
	return loader.typeAssertions
}

func filterDeclarations[D model.Declaration](
	decls map[string]D,
	typeFilters *typefilter.List,
) []D {
	result := make([]D, 0, len(decls))

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
