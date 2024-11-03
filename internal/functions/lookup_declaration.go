package functions

import "github.com/theunrepentantgeek/crddoc/internal/model"

func (f *Functions) lookupDeclaration(id string) model.Declaration {
	dec, _ := f.pkg.Declaration(id)

	return dec
}
