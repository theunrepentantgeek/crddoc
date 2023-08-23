package model

import "github.com/dave/dst"

type Enum struct {
	name        string
	base        dst.Expr
	usage       []Declaration
	description []string
	values      []*EnumValue
}

func TryNewEnum(spec dst.Spec, comments []string) (*Enum, bool) {
	// Check for a TypeSpec ...
	typeSpec, ok := spec.(*dst.TypeSpec)
	if !ok {
		return nil, false
	}

	// TEmp - skip structs
	if _, ok := typeSpec.Type.(*dst.StructType); ok {
		return nil, false
	}

	// ... that contains an identifier ...
	ident, ok := typeSpec.Type.(*dst.Ident)
	if !ok {
		return nil, false
	}

	result := &Enum{
		name: typeSpec.Name.Name,
		base: ident,
	}

	return result, true
}

func (e *Enum) Name() string {
	return e.name
}

func (e *Enum) Kind() DeclarationType {
	return EnumDeclaration
}

func (e *Enum) Usage() []Declaration {
	return e.usage
}

func (e *Enum) SetUsage(usage []Declaration) {
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
