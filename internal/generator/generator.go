package generator

import (
	"io"
	"path/filepath"
	"text/template"

	"github.com/go-logr/logr"
	"github.com/pkg/errors"
	"github.com/theunrepentantgeek/crddoc/internal/functions"
	"github.com/theunrepentantgeek/crddoc/internal/model"
)

type Generator struct {
	log      logr.Logger
	template *template.Template
}

func New(log logr.Logger) *Generator {
	funcMap := functions.CreateFuncMap()
	return &Generator{
		log:      log,
		template: template.New("crddoc").Funcs(funcMap),
	}
}

func (g *Generator) LoadTemplates(folder string) error {
	glob := filepath.Join(folder, "*.tmpl")
	_, err := g.template.ParseGlob(glob)
	if err != nil {
		return err
	}

	return nil
}

func (g *Generator) Generate(pkg *model.Package, writer io.Writer) error {
	err := g.template.ExecuteTemplate(
		writer,
		"crd",
		pkg)

	return errors.Wrap(err, "failed to execute template")
}
