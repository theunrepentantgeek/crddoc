package functions

import "github.com/theunrepentantgeek/crddoc/internal/model"

// asObject converts a declaration to an object, returning nil if that can't be done.
func (*Functions) asObject(declaration model.Declaration) *model.Object {
	if obj, ok := declaration.(*model.Object); ok {
		return obj
	}

	if res, ok := declaration.(*model.Resource); ok {
		return &res.Object
	}

	return nil
}
