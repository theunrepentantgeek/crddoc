package model

import (
	"reflect"
	"strings"

	"github.com/dave/dst"
)

type Property struct {
	Field       string // Name of the field
	Name        string // Serialized name of the field
	Type        TypeReference
	DeclaredOn  PropertyContainer
	description []string
	required    string
}

func TryNewProperty(name string, field *dst.Field) (*Property, bool) {
	// TODO: Parse tags as well
	description, commands := ParseComments(field.Decs.Start.All())
	description = formatComments(description, name)

	result := NewProperty(
		name,
		NewTypeReferenceFromExpr(field.Type),
		description)

	if commands.Any() {
		if commands.Exists("kubebuilder", "validation", "Optional") {
			result.required = "Optional"
		} else if commands.Exists("kubebuilder", "validation", "Required") {
			result.required = "Required"
		}
	}

	if name, ok := result.tryParseName(field); ok {
		result.Name = name
	}

	return result, true
}

func NewProperty(
	name string,
	ref TypeReference,
	description []string,
) *Property {
	return &Property{
		Name:        name,
		Field:       name,
		Type:        ref,
		description: description,
	}
}

func (p *Property) Required() string {
	return p.required
}

func (p *Property) Description() []string {
	return p.description
}

func (p *Property) tryParseName(field *dst.Field) (string, bool) {
	if field.Tag == nil || field.Tag.Value == "" {
		return "", false
	}

	tag := reflect.StructTag(strings.Trim(field.Tag.Value, "`"))

	// Try to find a name configured with a json tag
	if name, ok := p.tryParseNameFromTag("json", tag); ok {
		return name, true
	}

	// Try to find a name configured with a yaml tag
	if name, ok := p.tryParseNameFromTag("yaml", tag); ok {
		return name, true
	}

	return "", false
}

func (*Property) tryParseNameFromTag(
	tag string,
	tagStruct reflect.StructTag,
) (string, bool) {
	if y, ok := tagStruct.Lookup(tag); ok {
		parts := strings.Split(y, ",")

		name := parts[0]
		if name != "" {
			return name, true
		}
	}

	return "", false
}

func (p *Property) setContainer(container PropertyContainer) {
	p.DeclaredOn = container
}
