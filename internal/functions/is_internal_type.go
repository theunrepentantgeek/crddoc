package functions

import "github.com/theunrepentantgeek/crddoc/internal/model"

// isInternalType returns true if the type is defined in the current package, false otherwise.
// Always returns false for primitive types.
func (f *Functions) isInternalType(ref model.TypeReference) bool {
	_, found := f.pkg.Declaration(ref.ID())

	return found
}
