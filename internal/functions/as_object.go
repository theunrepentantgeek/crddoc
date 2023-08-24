package functions

import "github.com/theunrepentantgeek/crddoc/internal/model"

// asObject converts a declaration to an object, returning nil if that can't be done
func (f *Functions) asObject(declaration model.Declaration) *model.Object {
	if obj, ok := declaration.(*model.Object); ok {
		return obj
	}

	return nil
}
