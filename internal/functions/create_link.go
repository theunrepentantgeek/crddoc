package functions

import (
	"bytes"

	"github.com/theunrepentantgeek/crddoc/internal/model"
)

var createLinkLoggedTypes = map[model.TypeReference]bool{}

// createLink creates a link to the passed TypeReference, returning an empty string
// if no link is required.
func (f *Functions) createLink(ref *model.TypeReference) string {
	// Types in this package get anchors
	if _, ok := f.pkg.Declaration(ref.ID()); ok {
		return "#" + ref.ID()
	}

	if ref.ImportPath() != "" {
		return f.createExternalLink(ref)
	}

	return ""
}

func (f *Functions) createExternalLink(ref *model.TypeReference) string {
	// Check if we have a template for this import path
	if tmpl, ok := f.externalLinks[ref.ImportPath()]; ok {
		var buffer bytes.Buffer

		err := tmpl.Execute(&buffer, ref)
		if err != nil {
			f.log.Error(
				err,
				"Failed to generate external link",
				"importPath", ref.ImportPath())

			return ""
		}

		return buffer.String()
	}

	// Log a warning that we don't have configuration to generate an external link, but only once
	// per type. This gives users enough information to configure links if they choose.
	if _, ok := createLinkLoggedTypes[*ref]; !ok {
		createLinkLoggedTypes[*ref] = true

		f.log.V(0).Info(
			"Missing external link configuration",
			"package", ref.Package(),
			"importPath", ref.ImportPath())
	}

	return ""
}
