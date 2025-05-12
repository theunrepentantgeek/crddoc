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
	description []string
	pkg         *Package
	usage       []PropertyReference // List of other properties that reference this object
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
		if path, ok := importReferences.LookupImportPath(property.Type); ok {
			property.Type.impPath = path
		}
	}

	for _, embed := range o.embeds {
		if path, ok := importReferences.LookupImportPath(embed.Type); ok {
			embed.Type.impPath = path
		}
	}
}

// alphabeticalPropertyComparison does a case insensitive comparison of the names of the
// two properties, allowing them to be sorted.
func alphabeticalPropertyComparison(left *Property, right *Property) int {
	leftName := strings.ToLower(left.Name)
	rightName := strings.ToLower(right.Name)

	return strings.Compare(leftName, rightName)
}
