package templates

import "embed"

var (
    // Default CRD template
    //go:embed *.tmpl
    CRD embed.FS
)
