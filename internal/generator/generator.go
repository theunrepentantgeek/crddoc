package generator

import (
	"bytes"
	"io"
	"io/fs"
	"text/template"

	"github.com/go-logr/logr"
	"github.com/pkg/errors"
	"github.com/shurcooL/markdownfmt/markdown"
	"github.com/theunrepentantgeek/crddoc/internal/config"
	"github.com/theunrepentantgeek/crddoc/internal/functions"
	"github.com/theunrepentantgeek/crddoc/internal/model"
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

func (g *Generator) LoadTemplates(folder fs.FS) error {
	_, err := g.template.ParseFS(folder, "*.tmpl")
	if err != nil {
		return err
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
	g.fns.SetConfig(g.cfg)

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
