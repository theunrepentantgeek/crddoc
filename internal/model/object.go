package model

import (
	"strings"

	"github.com/dave/dst"
	"golang.org/x/exp/maps"
	"golang.org/x/exp/slices"
)

type Object struct {
	TypeReference
	properties  map[string]*Property
	description []string
	usage       []PropertyReference // List of other properties that reference this object
}

func TryNewObject(spec dst.Spec, comments []string) (*Object, bool) {
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

	ref := NewTypeReference(typeSpec.Name)
	description, _ := parseComments(comments)

	// If the first line of the description starts with "<name>: ", remove that prefix
	if len(description) > 0 {
		if s, ok := strings.CutPrefix(description[0], ref.Name()+": "); ok {
			description[0] = strings.TrimLeft(s, " ")
		}
	}

	result := &Object{
		TypeReference: ref,
		properties:    make(map[string]*Property),
		description:   description,
	}

	result.properties = result.findProperties(structType)

	return result, true
}

func (o *Object) Kind() DeclarationType {
	return ObjectDeclaration
}

func (o *Object) Usage() []PropertyReference {
	return o.usage
}

func (o *Object) SetUsage(uses []PropertyReference) {
	o.usage = uses
}

// Properties returns all the properties of the object, in alphabetical order
func (o *Object) Properties() []*Property {
	result := maps.Values(o.properties)
	slices.SortFunc(result, alphabeticalPropertyComparison)
	return result
}

func (o *Object) Property(name string) (*Property, bool) {
	prop, ok := o.properties[name]
	return prop, ok
}

func (o *Object) Description() []string {
	return o.description
}

func (o *Object) findProperties(structType *dst.StructType) map[string]*Property {
	result := make(map[string]*Property)

	// Iterate over the fields in the struct type and try to create a property for each one
	for _, field := range structType.Fields.List {

		// A single field might contain multiple properties
		for _, name := range field.Names {
			if property, ok := TryNewProperty(name.Name, field); ok {
				result[property.Name] = property
			}
		}
	}

	return result
}

func alphabeticalPropertyComparison(left *Property, right *Property) int {
	if left.name < right.name {
		return -1
	}

	if left.name > right.name {
		return 1
	}

	return 0
}
