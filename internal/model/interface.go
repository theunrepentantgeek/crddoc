package model

import (
	"maps"
	"slices"

	"github.com/dave/dst"
)

// Interface represents an interface type declaration.
type Interface struct {
	TypeReference
	methods         map[string]*Function
	embeds          []TypeReference
	description     []string
	pkg             *Package
	usage           []PropertyReference
	implementations []*Object
}

var _ Declaration = &Interface{}

// TryNewInterface attempts to create an Interface from a dst.Spec.
// Returns the interface and true if successful, nil and false otherwise.
func TryNewInterface(
	spec dst.Spec,
	comments []string,
	importReferences map[string]ImportReference,
) (*Interface, bool) {
	// Check for a TypeSpec ...
	typeSpec, ok := spec.(*dst.TypeSpec)
	if !ok {
		return nil, false
	}

	// ... that contains an InterfaceType
	interfaceType, ok := typeSpec.Type.(*dst.InterfaceType)
	if !ok {
		return nil, false
	}

	name := displayOf(typeSpec.Name)
	id := idOf(typeSpec.Name)

	// Clean up the comments
	description, _ := ParseComments(comments)
	description = formatComments(description, name)

	result := &Interface{
		TypeReference: TypeReference{
			name:    name,
			display: name,
			id:      id,
		},
		methods:     make(map[string]*Function),
		description: description,
	}

	result.parseAllMethods(interfaceType, importReferences)

	return result, true
}

// parseAllMethods parses the methods and embedded interfaces from an interface type.
func (i *Interface) parseAllMethods(
	interfaceType *dst.InterfaceType,
	importReferences ImportReferenceSet,
) {
	if interfaceType.Methods == nil {
		return
	}

	for _, field := range interfaceType.Methods.List {
		i.parseMethod(field, importReferences)
	}
}

func (i *Interface) parseMethod(field *dst.Field, importReferences ImportReferenceSet) int {
	// Embedded interfaces appear as fields without names
	// (e.g., `Greeter` in `type MultiTalent interface { Greeter }`).
	if len(field.Names) == 0 {
		// This is an embedded interface
		embedRef := NewTypeReferenceFromExpr(field.Type)
		if path, ok := importReferences.LookupImportPath(embedRef); ok {
			embedRef.impPath = path
		}

		i.embeds = append(i.embeds, embedRef)

		return 1
	}

	// Each name is a method
	for _, nameIdent := range field.Names {
		if fn := i.parseMethodFromField(nameIdent.Name, field, importReferences); fn != nil {
			i.methods[fn.Name] = fn
		}
	}

	return 0
}

// parseMethodFromField creates a Function from a field in an interface.
func (*Interface) parseMethodFromField(
	name string,
	field *dst.Field,
	importReferences ImportReferenceSet,
) *Function {
	// The field type should be a FuncType for interface methods
	funcType, ok := field.Type.(*dst.FuncType)
	if !ok {
		return nil
	}

	// Parse comments from the field
	description, _ := ParseComments(field.Decs.Start.All())
	description = formatComments(description, name)

	// Parse parameters
	var params []Parameter
	if funcType.Params != nil {
		params = parseFieldList(funcType.Params)
	}

	// Parse results
	var results []Parameter
	if funcType.Results != nil {
		results = parseFieldList(funcType.Results)
	}

	fn := &Function{
		Name: name,
		// Interface methods have no receiver - unlike struct methods which have either
		// a value (T) or pointer (*T) receiver, interface methods are just signatures.
		// The zero value TypeReference{} is used to indicate the absence of a receiver.
		Receiver:          TypeReference{},
		IsPointerReceiver: false,
		Parameters:        params,
		Results:           results,
		description:       description,
	}

	// Link import references for parameter and result types
	fn.linkParameters(importReferences)
	fn.linkResults(importReferences)

	return fn
}

// Kind returns the declaration type.
func (*Interface) Kind() DeclarationType {
	return InterfaceDeclaration
}

// Name returns the name of the interface.
func (i *Interface) Name() string {
	return i.display
}

// Description returns the description of the interface.
func (i *Interface) Description() []string {
	return i.description
}

// Usage returns the list of properties that reference this interface.
func (i *Interface) Usage() []PropertyReference {
	return i.usage
}

// SetUsage sets the list of properties that reference this interface.
func (i *Interface) SetUsage(usage []PropertyReference) {
	i.usage = usage
}

// Package returns the package this interface belongs to.
func (i *Interface) Package() *Package {
	return i.pkg
}

// SetPackage sets the package this interface belongs to.
func (i *Interface) SetPackage(pkg *Package) {
	i.pkg = pkg
}

// Methods returns all the methods of the interface, in alphabetical order.
func (i *Interface) Methods() FunctionList {
	result := slices.SortedFunc(
		maps.Values(i.methods),
		alphabeticalFunctionComparison)

	return result
}

// Method returns the method with the given name and true,
// or nil and false if not found.
func (i *Interface) Method(name string) (*Function, bool) {
	fn, ok := i.methods[name]

	return fn, ok
}

// Embeds returns all the embedded interfaces, in the order they were declared.
func (i *Interface) Embeds() []TypeReference {
	return slices.Clone(i.embeds)
}

// Implementations returns all the objects that implement this interface.
func (i *Interface) Implementations() []*Object {
	result := slices.Clone(i.implementations)
	slices.SortFunc(result, alphabeticalObjectComparison)

	return result
}

// AddImplementation adds an object as an implementation of this interface.
func (i *Interface) AddImplementation(obj *Object) {
	if obj == nil {
		return
	}

	// Check if already added
	for _, existing := range i.implementations {
		if existing.ID() == obj.ID() {
			return
		}
	}

	i.implementations = append(i.implementations, obj)
}
