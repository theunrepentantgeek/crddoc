package functions

import "github.com/theunrepentantgeek/crddoc/internal/model"

func (*Functions) asResource(declaration model.Declaration) *model.Resource {
	if res, ok := declaration.(*model.Resource); ok {
		return res
	}

	return nil
}
