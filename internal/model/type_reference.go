package model

import (
	"github.com/dave/dst"
)

type TypeReference struct {
	name string
	id   string
}

func NewTypeReference(expr dst.Expr) TypeReference {
	return TypeReference{
		name: createName(expr),
		id:   createId(expr),
	}
}

func (ref *TypeReference) Name() string {
	return ref.name
}

func (ref *TypeReference) Id() string {
	return ref.id
}

// createIdFor renders an ID from a type expression, for linking within the documentation.
// We used to create exclusively lowercase IDs, but some resources have related types
// that differ only in case, so we now preserve case.
func createId(expr dst.Expr) string {
	switch t := expr.(type) {
	case *dst.Ident:
		return t.Name
	case *dst.StarExpr:
		return createId(t.X)
	case *dst.ArrayType:
		return createId(t.Elt)
	case *dst.MapType:
		// For now, just use the value type
		// What should we do if the key type is useful?
		// createIdFor(t.Key, pkg)
		return createId(t.Value)
	default:
		return ""
	}
}

// createName renders a type expression as a string
func createName(expr dst.Expr) string {
	switch t := expr.(type) {
	case *dst.Ident:
		return t.Name
	case *dst.StarExpr:
		return createName(t.X)
	case *dst.SelectorExpr:
		return createName(t.X) + "." + createName(t.Sel)
	case *dst.ArrayType:
		return createName(t.Elt) + "[]"
	case *dst.MapType:
		return "map[" + createName(t.Key) + "]" + createName(t.Value)
	default:
		return ""
	}
}
