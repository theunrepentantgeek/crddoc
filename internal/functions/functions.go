package functions

import (
	"text/template"

	"github.com/go-logr/logr"
	"github.com/pkg/errors"

	"github.com/theunrepentantgeek/crddoc/internal/config"
	"github.com/theunrepentantgeek/crddoc/internal/model"
	"github.com/theunrepentantgeek/crddoc/internal/texteditor"
)

// Functions is a carrier type for all the Functions we provide to templates.
type Functions struct {
	pkg           *model.Package // Current Package used by the functions.
	cfg           *config.Config // Config used by the functions.
	editors       *texteditor.List
	log           logr.Logger
	externalLinks map[string]*template.Template
}

func New(log logr.Logger) *Functions {
	return &Functions{
		externalLinks: make(map[string]*template.Template),
		log:           log,
	}
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
		"externalOnly":            f.externalOnly,
		"includeClassDiagrams":    f.includeClassDiagrams,
		"inlineLinks":             f.inlineLinks,
		"internalOnly":            f.internalOnly,
		"lookupDeclaration":       f.lookupDeclaration,
		"unwrap":                  f.unwrap,
	}
}

func (f *Functions) SetPackage(pkg *model.Package) {
	f.pkg = pkg
}

func (f *Functions) SetConfig(cfg *config.Config) error {
	f.cfg = cfg

	var err error

	f.editors, err = texteditor.New(cfg)
	if err != nil {
		return errors.Wrap(err, "configuring functions")
	}

	for _, link := range cfg.ExternalLinks {
		tmpl, err := template.New(link.ImportPath).Parse(link.URLTemplate)
		if err != nil {
			return errors.Wrapf(err, "parsing external link template for %q", link.ImportPath)
		}

		f.externalLinks[link.ImportPath] = tmpl
	}

	return nil
}
