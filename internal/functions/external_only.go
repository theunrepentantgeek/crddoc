package functions

import "github.com/theunrepentantgeek/crddoc/internal/model"

// exteralOnly selects only properties with types decared in the same package
func (f *Functions) externalOnly(properties []*model.Property) []*model.Property {
	result := make([]*model.Property, 0, len(properties))
	for _, prop := range properties {
		if _, found := f.pkg.Declaration(prop.Type.ID()); found {
			result = append(result, prop)
		}
	}
	return result
}
