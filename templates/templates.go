package templates

import "embed"

// Default CRD template
//
//go:embed *.tmpl
var CRD embed.FS
