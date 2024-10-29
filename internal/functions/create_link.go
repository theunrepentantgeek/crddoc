package functions

import "github.com/theunrepentantgeek/crddoc/internal/model"

// createLink creates a link to the passed TypeReference, returning an empty string
// if no link is required.
func (f *Functions) createLink(ref *model.TypeReference) string {
	// Types in this package get anchors
	if _, ok := f.pkg.Declaration(ref.Id()); ok {
		return "#" + ref.Id()
	}

	// TODO: External links

	return ""
}
