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
	fns      *functions.Functions
}

func New(log logr.Logger) *Generator {
	fns := functions.New()
	funcMap := fns.CreateFuncMap()
	return &Generator{
		log:      log,
		fns:      fns,
		template: template.New("crddoc").Funcs(funcMap),
	}
}

func (g *Generator) LoadTemplates(folder string) error {
	g.log.Info(
		"Loading templates",
		"templateFolder", folder)
	glob := filepath.Join(folder, "*.tmpl")
	_, err := g.template.ParseGlob(glob)
	if err != nil {
		return err
	}

	return nil
}

func (g *Generator) Generate(pkg *model.Package, writer io.Writer) error {
	g.log.Info(
		"Rendering template",
		"package", pkg.Name(),
	)

	g.fns.SetPackage(pkg)

	err := g.template.ExecuteTemplate(
		writer,
		"crd",
		pkg)

	if err != nil {
		g.log.Error(err, "failed to execute template")
		return errors.Wrap(err, "failed to execute template")
	}

	g.log.Info("Template rendered successfully")
	return nil
}
