package functions

import "github.com/dave/dst"

// asText renders a type expression as a string
func (f *Functions) asText(expr dst.Expr) string {
	switch t := expr.(type) {
	case *dst.Ident:
		return t.Name
	case *dst.StarExpr:
		return f.asText(t.X)
	case *dst.SelectorExpr:
		return f.asText(t.X) + "." + f.asText(t.Sel)
	case *dst.ArrayType:
		return f.asText(t.Elt) + "[]"
	case *dst.MapType:
		return "map[" + f.asText(t.Key) + "]" + f.asText(t.Value)
	default:
		return "<type>"
	}
}
