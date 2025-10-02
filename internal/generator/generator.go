package generator

import (
	"bufio"
	"bytes"
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"text/template"

	sprig "github.com/go-task/slim-sprig/v3"

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
	fns := functions.New(log)
	funcMap := fns.CreateFuncMap()

	tmpl := template.New("crddoc").
		Funcs(sprig.FuncMap()).
		Funcs(funcMap)

	return &Generator{
		cfg:      cfg,
		log:      log,
		fns:      fns,
		template: tmpl,
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
		sourceDescription = "loading templates from folder " + g.cfg.TemplatePath
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

func (g *Generator) GenerateToFile(
	pkg *model.Package,
	outputPath string,
	log logr.Logger,
) error {
	// Render the template and write to output
	f, err := os.Create(outputPath)
	if err != nil {
		return errors.Wrapf(err, "creating output file %q", outputPath)
	}

	defer f.Close()

	log.Info("Writing to", "outputPath", outputPath)

	w := bufio.NewWriter(f)

	err = g.GenerateToWriter(pkg, w)
	if err != nil {
		return errors.Wrapf(err, "generating output to %q", outputPath)
	}

	err = w.Flush()
	if err != nil {
		return errors.Wrapf(err, "flushing output to %q", outputPath)
	}

	return nil
}

func (g *Generator) GenerateToMultipleFiles(
	pkg *model.Package,
	outputDir string,
	log logr.Logger,
) error {
	// Create output directory if it doesn't exist
	if err := os.MkdirAll(outputDir, 0o755); err != nil {
		return errors.Wrapf(err, "creating output directory %q", outputDir)
	}

	log.Info("Writing multiple files to directory", "outputDir", outputDir)

	// Get all declarations and filter for structs (objects and resources)
	declarations := pkg.Declarations(model.OrderAlphabetical)

	structCount := 0

	for _, decl := range declarations {
		// Include both Objects and Resources as they are both structs
		if decl.Kind() == model.ObjectDeclaration || decl.Kind() == model.ResourceDeclaration {
			structCount++
			filename := decl.Name() + ".md"
			outputPath := filepath.Join(outputDir, filename)

			log.Info(
				"Writing struct to file",
				"struct", decl.Name(),
				"kind", decl.Kind(),
				"outputPath", outputPath)

			if err := g.generateObjectToFile(pkg, decl, outputPath); err != nil {
				return errors.Wrapf(
					err,
					"generating output for struct %q to %q",
					decl.Name(),
					outputPath)
			}
		}
	}

	if structCount == 0 {
		log.Info("No structs found to generate files for")
	} else {
		log.Info("Generated multiple files", "count", structCount, "outputDir", outputDir)
	}

	return nil
}

func (g *Generator) generateObjectToFile(
	pkg *model.Package,
	decl model.Declaration,
	outputPath string,
) error {
	// Create the output file
	f, err := os.Create(outputPath)
	if err != nil {
		return errors.Wrapf(err, "creating output file %q", outputPath)
	}

	defer f.Close()

	w := bufio.NewWriter(f)

	// Generate content for the specific declaration
	err = g.generateObjectToWriter(pkg, decl, w)
	if err != nil {
		return errors.Wrapf(err, "generating content for declaration %q", decl.Name())
	}

	err = w.Flush()
	if err != nil {
		return errors.Wrapf(err, "flushing output to %q", outputPath)
	}

	return nil
}

func (g *Generator) generateObjectToWriter(
	pkg *model.Package,
	decl model.Declaration,
	writer io.Writer,
) error {
	g.log.Info(
		"Rendering template for declaration",
		"declaration", decl.Name(),
		"kind", decl.Kind(),
		"package", pkg.Name(),
		"group", pkg.Group(),
		"version", pkg.Version(),
	)

	g.fns.SetPackage(pkg)

	if err := g.fns.SetConfig(g.cfg); err != nil {
		g.log.Error(err, "failed to set function config")

		return errors.Wrap(err, "failed to set function config")
	}

	var raw bytes.Buffer

	// Use the single-object template to render just this declaration
	err := g.template.ExecuteTemplate(
		&raw,
		"single-object",
		decl)
	if err != nil {
		g.log.Error(err, "failed to execute single-object template")

		return errors.Wrap(err, "failed to execute single-object template")
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

	g.log.Info("Declaration template rendered successfully", "declaration", decl.Name())

	return nil
}

func (g *Generator) GenerateToWriter(
	pkg *model.Package,
	writer io.Writer,
) error {
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
