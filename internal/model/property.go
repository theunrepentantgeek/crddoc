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
	description []string
	required    string
}

func TryNewProperty(name string, field *dst.Field) (*Property, bool) {
	// TODO: Parse tags

	description, commands := ParseComments(field.Decs.Start.All())

	// If the first line of the description starts with "<property>: ", remove that prefix
	if len(description) > 0 {
		if s, ok := strings.CutPrefix(description[0], name+": "); ok {
			description[0] = strings.TrimLeft(s, " ")
		}
	}

	result := &Property{
		Field:       name,
		Name:        name,
		Type:        NewTypeReference(field.Type),
		description: description,
	}

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
	if j, ok := tag.Lookup("json"); ok {
		parts := strings.Split(j, ",")
		name := parts[0]
		if name != "" {
			return name, true
		}
	}

	// Try to find a name configured with a yaml tag
	if y, ok := tag.Lookup("yaml"); ok {
		parts := strings.Split(y, ",")
		name := parts[0]
		if name != "" {
			return name, true
		}
	}

	return "", false
}
