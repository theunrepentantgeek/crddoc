package functions

import "github.com/dave/dst"

// renderType renders a type expression as a string
func renderType(expr dst.Expr) string {
	switch t := expr.(type) {
	case *dst.Ident:
		return t.Name
	case *dst.StarExpr:
		return renderType(t.X)
	case *dst.SelectorExpr:
		return renderType(t.X) + "." + renderType(t.Sel)
	case *dst.ArrayType:
		return renderType(t.Elt) + "[]"
	case *dst.MapType:
		return "map[" + renderType(t.Key) + "]" + renderType(t.Value)
	default:
		return "<type>"
	}
}
