package functions

import (
	"github.com/dave/dst"
)

// asId renders an ID from a type expression, for linking within the documentation.
// Returns an empty string if the object does not exist in the package.
func (f *Functions) asId(expr dst.Expr) string {
	switch t := expr.(type) {
	case *dst.Ident:
		if !f.objectExists(t.Name) {
			return ""
		}

		return t.Name
	case *dst.StarExpr:
		return f.asId(t.X)
	case *dst.ArrayType:
		return f.asId(t.Elt)
	case *dst.MapType:
		return f.asId(t.Key) + f.asId(t.Value)
	default:
		return ""
	}
}
