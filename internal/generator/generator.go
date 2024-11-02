package generator

import (
	"bytes"
	"fmt"
	"io"
	"io/fs"
	"os"
	"text/template"

	"github.com/go-logr/logr"
	"github.com/pkg/errors"
	"github.com/shurcooL/markdownfmt/markdown"

	"github.com/theunrepentantgeek/crddoc/internal/config"
	"github.com/theunrepentantgeek/crddoc/internal/functions"
	"github.com/theunrepentantgeek/crddoc/internal/model"
	"github.com/theunrepentantgeek/crddoc/templates"
)

type Generator struct {
	cfg      *config.Config
	log      logr.Logger
	template *template.Template
	fns      *functions.Functions
}

func New(cfg *config.Config, log logr.Logger) *Generator {
	fns := functions.New()
	funcMap := fns.CreateFuncMap()

	return &Generator{
		cfg:      cfg,
		log:      log,
		fns:      fns,
		template: template.New("crddoc").Funcs(funcMap),
	}
}

func (g *Generator) LoadTemplates() error {
	var sourceFS fs.FS
	var sourceDescription string

	if g.cfg.TemplatePath != "" {
		// Load templates from the specified folder
		g.log.Info(
			"Loading templates",
			"path", g.cfg.TemplatePath)
		sourceFS = os.DirFS(g.cfg.TemplatePath)
		sourceDescription = fmt.Sprintf("loading templates from folder %s", g.cfg.TemplatePath)
	} else {
		// Use internal templates
		g.log.Info("Loading internal templates")
		sourceFS = templates.CRD
		sourceDescription = "loading internal templates"
	}

	_, err := g.template.ParseFS(sourceFS, "*.tmpl")
	if err != nil {
		return errors.Wrap(err, sourceDescription)
	}

	return nil
}

func (g *Generator) Generate(pkg *model.Package, writer io.Writer) error {
	g.log.Info(
		"Rendering template",
		"package", pkg.Name(),
		"group", pkg.Group(),
		"version", pkg.Version(),
	)

	g.fns.SetPackage(pkg)
	if err := g.fns.SetConfig(g.cfg); err != nil {
		g.log.Error(err, "failed to set function config")

		return errors.Wrap(err, "failed to set")
	}

	var raw bytes.Buffer
	err := g.template.ExecuteTemplate(
		&raw,
		"crd",
		pkg)
	if err != nil {
		g.log.Error(err, "failed to execute template")

		return errors.Wrap(err, "failed to execute template")
	}

	content := raw.Bytes()

	if g.cfg.PrettyPrint {
		content, err = markdown.Process("", raw.Bytes(), nil)
		if err != nil {
			g.log.Error(err, "failed to tidy markdown")

			return errors.Wrap(err, "failed to tidy markdown")
		}
	}

	_, err = writer.Write(content)
	if err != nil {
		g.log.Error(err, "failed to write markdown")

		return errors.Wrap(err, "failed to write markdown")
	}

	g.log.Info("Template rendered successfully")

	return nil
}
