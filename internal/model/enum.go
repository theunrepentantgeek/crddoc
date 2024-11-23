package model

import "github.com/dave/dst"

type Enum struct {
	TypeReference
	usage       []PropertyReference
	description []string
	values      []*EnumValue
	base        TypeReference
	pkg         *Package
}

func TryNewEnum(spec dst.Spec, comments []string) (*Enum, bool) {
	// Check for a TypeSpec ...
	typeSpec, ok := spec.(*dst.TypeSpec)
	if !ok {
		return nil, false
	}

	// ... (skipping structs) ...
	if _, isStruct := typeSpec.Type.(*dst.StructType); isStruct {
		return nil, false
	}

	// ... that contains an identifier ...
	ident, ok := typeSpec.Type.(*dst.Ident)
	if !ok {
		return nil, false
	}

	result := &Enum{
		TypeReference: NewTypeReferenceFromExpr(typeSpec.Name),
		base:          NewTypeReferenceFromExpr(ident),
		description:   comments,
	}

	return result, true
}

func (e *Enum) Name() string {
	return e.name
}

func (*Enum) Kind() DeclarationType {
	return EnumDeclaration
}

func (e *Enum) Usage() []PropertyReference {
	return e.usage
}

func (e *Enum) SetUsage(usage []PropertyReference) {
	e.usage = usage
}

func (e *Enum) SetPackage(pkg *Package) {
	e.pkg = pkg
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
