package model

import (
	"maps"
	"slices"
	"strings"

	"github.com/dave/dst"
)

type Object struct {
	TypeReference
	properties  map[string]*Property
	embeds      PropertyList
	functions   map[string]*Function
	description []string
	pkg         *Package
	usage       []PropertyReference // List of other properties that reference this object
	interfaces  []*Interface        // List of interfaces this object implements
}

var _ Declaration = &Object{}

func TryNewObject(
	spec dst.Spec,
	comments []string,
	importReferences map[string]ImportReference,
) (*Object, bool) {
	// Check for a TypeSpec ...
	typeSpec, ok := spec.(*dst.TypeSpec)
	if !ok {
		return nil, false
	}

	// ... that contains a StructType ...
	structType, ok := typeSpec.Type.(*dst.StructType)
	if !ok {
		return nil, false
	}

	// ... with at least one field
	if len(structType.Fields.List) == 0 {
		return nil, false
	}

	name := displayOf(typeSpec.Name)
	id := idOf(typeSpec.Name)

	// Clean up the comments
	description, _ := ParseComments(comments)
	description = formatComments(description, name)

	result := NewObjectType(name, id, description)

	result.properties = result.findProperties(structType)
	result.embeds = result.findEmbeddedStructs(structType)
	result.linkImports(importReferences)

	return result, true
}

func NewObjectType(
	name string,
	id string,
	description []string,
) *Object {
	return &Object{
		TypeReference: TypeReference{
			name:    name,
			display: name,
			id:      id,
		},
		properties:  make(map[string]*Property),
		functions:   make(map[string]*Function),
		description: description,
	}
}

func (*Object) Kind() DeclarationType {
	return ObjectDeclaration
}

func (o *Object) Usage() []PropertyReference {
	return o.usage
}

func (o *Object) SetUsage(uses []PropertyReference) {
	o.usage = uses
}

func (o *Object) Package() *Package {
	return o.pkg
}

func (o *Object) SetPackage(p *Package) {
	o.pkg = p
}

// Properties returns all the properties of the object, in alphabetical order.
func (o *Object) Properties() PropertyList {
	result := slices.SortedFunc(
		maps.Values(o.properties),
		alphabeticalPropertyComparison)

	return result
}

// Property returns the property with the given name and true,
// or nil and false if not found.
func (o *Object) Property(name string) (*Property, bool) {
	prop, ok := o.properties[name]

	return prop, ok
}

// Embeds returns all of the embeds of the object, in alphabetical order.
func (o *Object) Embeds() PropertyList {
	result := slices.Clone(o.embeds)
	slices.SortFunc(result, alphabeticalPropertyComparison)

	return result
}

// Embed returns the embed with the given name and true,
// or nil and false if not found.
func (o *Object) Embed(name string) (*Property, bool) {
	for _, embed := range o.embeds {
		if embed.Name == name {
			return embed, true
		}
	}

	return nil, false
}

func (o *Object) Description() []string {
	return o.description
}

// Functions returns all the functions/methods of the object, in alphabetical order.
func (o *Object) Functions() FunctionList {
	result := slices.SortedFunc(
		maps.Values(o.functions),
		alphabeticalFunctionComparison)

	return result
}

// Function returns the function with the given name and true,
// or nil and false if not found.
func (o *Object) Function(name string) (*Function, bool) {
	fn, ok := o.functions[name]

	return fn, ok
}

// AddFunction adds a function to the object.
func (o *Object) AddFunction(fn *Function) {
	if fn == nil {
		return
	}

	fn.setDeclaredOn(o)
	o.functions[fn.Name] = fn
}

// ImplementsInterfaces returns all the interfaces this object implements, in alphabetical order.
func (o *Object) ImplementsInterfaces() []*Interface {
	result := slices.Clone(o.interfaces)
	slices.SortFunc(result, alphabeticalInterfaceComparison)

	return result
}

// AddInterface adds an interface that this object implements.
func (o *Object) AddInterface(iface *Interface) {
	if iface == nil {
		return
	}

	// Check if already added
	for _, existing := range o.interfaces {
		if existing.ID() == iface.ID() {
			return
		}
	}

	o.interfaces = append(o.interfaces, iface)
}

func (o *Object) findProperties(structType *dst.StructType) map[string]*Property {
	result := make(map[string]*Property)

	// Iterate over the fields in the struct type and try to create a property for each one.
	for _, field := range structType.Fields.List {
		// A single field might contain multiple properties.
		for _, name := range field.Names {
			if property, ok := TryNewProperty(name.Name, field); ok {
				property.setContainer(o)
				result[property.Name] = property
			}
		}
	}

	return result
}

func (*Object) findEmbeddedStructs(structType *dst.StructType) PropertyList {
	var result PropertyList

	// Iterate over the fields in the struct type and try to create a property for each one.
	for _, field := range structType.Fields.List {
		if field.Names != nil {
			continue
		}

		// Embedded struct
		if property, ok := TryNewProperty("", field); ok {
			result = append(result, property)
		}
	}

	return result
}

func (o *Object) linkImports(importReferences ImportReferenceSet) {
	for _, property := range o.properties {
		o.linkImportsToType(&property.Type, importReferences)
	}

	for _, embed := range o.embeds {
		o.linkImportsToType(&embed.Type, importReferences)
	}

	for _, function := range o.functions {
		// Link receiver type
		o.linkImportsToType(&function.Receiver, importReferences)

		// Link parameter types
		for i := range function.Parameters {
			o.linkImportsToType(&function.Parameters[i].Type, importReferences)
		}

		// Link result types
		for i := range function.Results {
			o.linkImportsToType(&function.Results[i].Type, importReferences)
		}
	}
}

// linkImportsToType links a single TypeReference to its import path if available.
func (*Object) linkImportsToType(typeRef *TypeReference, importReferences ImportReferenceSet) {
	if path, ok := importReferences.LookupImportPath(*typeRef); ok {
		typeRef.impPath = path
	}
}

// alphabeticalPropertyComparison does a case insensitive comparison of the names of the
// two properties, allowing them to be sorted.
func alphabeticalPropertyComparison(left *Property, right *Property) int {
	leftName := strings.ToLower(left.Name)
	rightName := strings.ToLower(right.Name)

	return strings.Compare(leftName, rightName)
}

// alphabeticalFunctionComparison does a case insensitive comparison of the names of the
// two functions, allowing them to be sorted.
func alphabeticalFunctionComparison(left *Function, right *Function) int {
	leftName := strings.ToLower(left.Name)
	rightName := strings.ToLower(right.Name)

	return strings.Compare(leftName, rightName)
}

// alphabeticalInterfaceComparison does a case insensitive comparison of the names of the
// two interfaces, allowing them to be sorted.
func alphabeticalInterfaceComparison(left *Interface, right *Interface) int {
	leftName := strings.ToLower(left.Name())
	rightName := strings.ToLower(right.Name())

	return strings.Compare(leftName, rightName)
}
