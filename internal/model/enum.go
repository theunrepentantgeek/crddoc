package model

import "github.com/dave/dst"

type Enum struct {
	TypeReference
	usage       []PropertyReference
	description []string
	values      []*EnumValue
	base        TypeReference
}

func TryNewEnum(spec dst.Spec, comments []string) (*Enum, bool) {
	// Check for a TypeSpec ...
	typeSpec, ok := spec.(*dst.TypeSpec)
	if !ok {
		return nil, false
	}

	// ... (skipping structs) ...
	if _, ok := typeSpec.Type.(*dst.StructType); ok {
		return nil, false
	}

	// ... that contains an identifier ...
	ident, ok := typeSpec.Type.(*dst.Ident)
	if !ok {
		return nil, false
	}

	result := &Enum{
		TypeReference: NewTypeReference(typeSpec.Name),
		base:          NewTypeReference(ident),
		description:   comments,
	}

	return result, true
}

func (e *Enum) Name() string {
	return e.name
}

func (e *Enum) Kind() DeclarationType {
	return EnumDeclaration
}

func (e *Enum) Usage() []PropertyReference {
	return e.usage
}

func (e *Enum) SetUsage(usage []PropertyReference) {
	e.usage = usage
}

func (e *Enum) Description() []string {
	return e.description
}

func (e *Enum) Values() []*EnumValue {
	return e.values
}

func (e *Enum) AddValue(value *EnumValue) {
	e.values = append(e.values, value)
}
