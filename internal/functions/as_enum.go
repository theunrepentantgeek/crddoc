package functions

import "github.com/theunrepentantgeek/crddoc/internal/model"

// asEnum converts a declaration to an enum, returning nil if that can't be done
func (f *Functions) asEnum(declaration model.Declaration) *model.Enum {
	if e, ok := declaration.(*model.Enum); ok {
		return e
	}

	return nil
}
