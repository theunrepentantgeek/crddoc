package model

import (
	"strings"

	"github.com/dave/dst"
)

type Property struct {
	Name        string
	Type        TypeReference
	description []string
	required    string
}

func TryNewProperty(name string, field *dst.Field) (*Property, bool) {
	//TODO: Parse tags

	description, commands := parseComments(field.Decs.Start.All())

	// If the first line of the description starts with "<property>: ", remove that prefix
	if len(description) > 0 {
		if s, ok := strings.CutPrefix(description[0], name+": "); ok {
			description[0] = strings.TrimLeft(s, " ")
		}
	}

	result := &Property{
		Name:        name,
		Type:        NewTypeReference(field.Type),
		description: description,
	}

	if commands.Count() > 0 {
		if commands.Exists("kubebuilder", "validation", "Optional") {
			result.required = "Optional"
		} else if commands.Exists("kubebuilder", "validation", "Required") {
			result.required = "Required"
		}
	}

	return result, true
}

func (p *Property) Name() string {
	return p.name
}

func (p *Property) Type() dst.Expr {
	return p.propertyType
}

func (p *Property) Required() string {
	return p.required
}

func (p *Property) Description() []string {
	return p.description
}
