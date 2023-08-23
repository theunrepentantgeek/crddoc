package model

import "github.com/dave/dst"

type EnumValue struct {
	kind        string
	name        string
	value       *dst.BasicLit
	description []string
}

func TryNewEnumValue(spec dst.Spec) (*EnumValue, bool) {
	// Check for a ValueSpec ...
	valueSpec, ok := spec.(*dst.ValueSpec)
	if !ok {
		return nil, false
	}

	// ... that contains a single name ...
	if len(valueSpec.Names) != 1 {
		return nil, false
	}

	name := valueSpec.Names[0]

	// ... and a single value ...
	if len(valueSpec.Values) != 1 {
		return nil, false
	}

	// ... that is a call expression ...
	call, ok := valueSpec.Values[0].(*dst.CallExpr)
	if !ok {
		return nil, false
	}

	// ... to a simple function ...
	fn, ok := call.Fun.(*dst.Ident)
	if !ok {
		return nil, false
	}

	// ... with a single argument ...
	if len(call.Args) != 1 {
		return nil, false
	}

	// ... that is a basic literal ...
	lit, ok := call.Args[0].(*dst.BasicLit)
	if !ok {
		return nil, false
	}

	result := &EnumValue{
		kind:  fn.Name,
		name:  name.Name,
		value: lit,
	}

	return result, true
}

func (v *EnumValue) Kind() string {
	return v.kind
}

func (v *EnumValue) Name() string {
	return v.name
}

func (v *EnumValue) Value() string {
	return v.value.Value
}

func (v *EnumValue) Description() []string {
	return v.description
}
