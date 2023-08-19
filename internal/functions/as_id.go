package functions

import (
	"github.com/dave/dst"
)

// asId renders an ID from a type expression, for linking within the documentation.
// Returns an empty string if the object does not exist in the package.
func (f *Functions) asId(expr dst.Expr) string {
	return f.pkg.CreateIdFor(expr)
}
