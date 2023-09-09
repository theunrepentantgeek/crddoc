package functions

import (
	"text/template"

	"github.com/theunrepentantgeek/crddoc/internal/config"
	"github.com/theunrepentantgeek/crddoc/internal/model"
)

// Functions is a carrier type for all the Functions we provide to templates
type Functions struct {
	pkg *model.Package // Current Package used by the functions
	cfg *config.Config // Config used by the functions
}

func New() *Functions {
	return &Functions{}
}

func (f *Functions) CreateFuncMap() template.FuncMap {
	return template.FuncMap{
		"applyEdits":              f.applyEdits,
		"asEnum":                  f.asEnum,
		"asObject":                f.asObject,
		"asPropertyReferenceList": asList[model.PropertyReference],
		"asPropertyContainer":     f.asPropertyContainer,
		"asResource":              f.asResource,
		"createLink":              f.createLink,
		"lookupDeclaration":       f.lookupDeclaration,
		"unwrap":                  f.unwrap,
	}
}

func (f *Functions) SetPackage(pkg *model.Package) {
	f.pkg = pkg
}

func (f *Functions) SetConfig(cfg *config.Config) {
	f.cfg = cfg
}
