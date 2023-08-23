package functions

import (
	"text/template"

	"github.com/theunrepentantgeek/crddoc/internal/model"
)

// Functions is a carrier type for all the Functions we provide to templates
type Functions struct {
	pkg *model.Package // Current Package used by the functions
}

func New() *Functions {
	return &Functions{}
}

func (f *Functions) CreateFuncMap() template.FuncMap {
	return template.FuncMap{
		"asEnum":   f.asEnum,
		"asId":     f.asId,
		"asObject": f.asObject,
		"asText":   f.asText,
		"unwrap":   f.unwrap,
	}
}

func (f *Functions) SetPackage(pkg *model.Package) {
	f.pkg = pkg
}

func (f *Functions) objectExists(name string) bool {
	_, ok := f.pkg.Object(name)
	return ok
}
