package model

import (
	"regexp"
	"strings"

	"github.com/dave/dst"
)

type ImportReference struct {
	ImportPath string
	Alias      string
}

func TryNewImportReference(imp *dst.ImportSpec) (ImportReference, bool) {
	if imp == nil {
		return ImportReference{}, false
	}

	alias := ""
	if imp.Name != nil {
		alias = imp.Name.Name
	}

	path := strings.Trim(imp.Path.Value, "\"")

	return ImportReference{
		ImportPath: path,
		Alias:      alias,
	}, true
}

var versionRegex = regexp.MustCompile(`v\d+$`)

func (imp *ImportReference) Name() string {
	if imp.Alias != "" {
		return imp.Alias
	}

	if !strings.Contains(imp.ImportPath, "/") {
		return imp.ImportPath
	}

	parts := strings.Split(imp.ImportPath, "/")
	index := len(parts) - 1

	if index > 0 && versionRegex.MatchString(parts[index]) {
		index--
	}

	return parts[index]
}
