package model

import (
	"strings"

	"github.com/dave/dst"
	"golang.org/x/exp/maps"
	"golang.org/x/exp/slices"
)

type Object struct {
	name        string
	typeSpec    *dst.TypeSpec
	structType  *dst.StructType
	properties  map[string]*Property
	description []string
	usage       []Declaration
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

	name := typeSpec.Name.Name
	description, _ := parseComments(comments)

	// If the first line of the description starts with "<type>: ", remove that prefix
	if len(description) > 0 {
		if s, ok := strings.CutPrefix(description[0], name+": "); ok {
			description[0] = strings.TrimLeft(s, " ")
		}
	}

	result := &Object{
		name:        name,
		typeSpec:    typeSpec,
		structType:  structType,
		properties:  make(map[string]*Property),
		description: description,
	}

	for _, property := range result.findProperties() {
		result.properties[property.Name()] = property
	}

	return result, true
}

func (o *Object) Name() string {
	return o.name
}

func (o *Object) Kind() DeclarationType {
	return ObjectDeclaration
}

func (o *Object) Usage() []Declaration {
	return o.usage
}

func (o *Object) SetUsage(uses []Declaration) {
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

func (o *Object) findProperties() []*Property {
	var result []*Property

	// Iterate over the fields in the struct type and try to create a property for each one
	for _, field := range o.structType.Fields.List {

		// A single field might contain multiple properties
		for _, name := range field.Names {
			if property, ok := TryNewProperty(name.Name, field); ok {
				result = append(result, property)
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
