package model

import (
	"github.com/dave/dst"
)

type TypeReference struct {
	name    string // name of the type e.g. 'Writer'
	display string // display name for the type e.g. 'io.Writer'
	id      string // unique internal ID for the type
	pkg     string // import reference  for the type e.g. 'io' from 'io.Writer'
	impPath string // import path for the type e.g. 'https://pkg.go.dev/github.com/rotisserie/eris'
}

func NewTypeReferenceFromExpr(expr dst.Expr) TypeReference {
	name := nameOf(expr)
	display := displayOf(expr)
	id := idOf(expr)
	pkg := pkgOf(expr)

	return TypeReference{
		name:    name,
		display: display,
		id:      id,
		pkg:     pkg,
	}
}

func (ref *TypeReference) Name() string {
	return ref.name
}

func (ref *TypeReference) Display() string {
	return ref.display
}

func (ref *TypeReference) ID() string {
	return ref.id
}

func (ref *TypeReference) Package() string {
	return ref.pkg
}

func (ref *TypeReference) ImportPath() string {
	return ref.impPath
}

// createIdFor renders an ID from a type expression, for linking within the documentation.
// We used to create exclusively lowercase IDs, but some resources have related types
// that differ only in case, so we now preserve case.
func idOf(expr dst.Expr) string {
	switch t := expr.(type) {
	case *dst.Ident:
		return t.Name
	case *dst.StarExpr:
		return idOf(t.X)
	case *dst.ArrayType:
		return idOf(t.Elt)
	case *dst.MapType:
		// For now, just use the value type
		// TODO: What should we do if the key type is useful?
		// createIdFor(t.Key, pkg)
		return idOf(t.Value)
	default:
		return ""
	}
}

// nameOf renders a type expression as a string.
func nameOf(expr dst.Expr) string {
	switch t := expr.(type) {
	case *dst.Ident:
		return t.Name
	case *dst.StarExpr:
		return nameOf(t.X)
	case *dst.SelectorExpr:
		return nameOf(t.Sel)
	case *dst.ArrayType:
		return nameOf(t.Elt)
	case *dst.MapType:
		return nameOf(t.Value)
	default:
		return ""
	}
}

// displayOf renders a type expression as a string.
func displayOf(expr dst.Expr) string {
	switch t := expr.(type) {
	case *dst.Ident:
		return t.Name
	case *dst.StarExpr:
		return displayOf(t.X)
	case *dst.SelectorExpr:
		return displayOf(t.X) + "." + displayOf(t.Sel)
	case *dst.ArrayType:
		return displayOf(t.Elt) + "[]"
	case *dst.MapType:
		return "map[" + displayOf(t.Key) + "]" + displayOf(t.Value)
	default:
		return ""
	}
}

func pkgOf(expr dst.Expr) string {
	switch t := expr.(type) {
	case *dst.Ident:
		return ""
	case *dst.StarExpr:
		return pkgOf(t.X)
	case *dst.SelectorExpr:
		return displayOf(t.X)
	case *dst.ArrayType:
		return pkgOf(t.Elt)
	case *dst.MapType:
		// Again we're assuming the Value is the important part
		// TODO: verify this
		return pkgOf(t.Value)
	default:
		return ""
	}
}
