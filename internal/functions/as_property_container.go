package functions

import "github.com/theunrepentantgeek/crddoc/internal/model"

func (*Functions) asPropertyContainer(declaration model.Declaration) model.PropertyContainer {
	if p, ok := declaration.(model.PropertyContainer); ok {
		return p
	}

	return nil
}
