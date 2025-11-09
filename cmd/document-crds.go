package cmd

import (
	"os"

	"github.com/go-logr/logr"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"

	"github.com/theunrepentantgeek/crddoc/internal/config"
	"github.com/theunrepentantgeek/crddoc/internal/generator"
	"github.com/theunrepentantgeek/crddoc/internal/model"
	"github.com/theunrepentantgeek/crddoc/internal/packageloader"
)

func newDocumentCRDsCommand(log logr.Logger) (*cobra.Command, error) {
	options := &documentCRDsOptions{}

	cmd := &cobra.Command{
		Use:   "crds",
		Short: "Generate CRD documentation from a package",
		Long:  "Generate CRD documentation from a package.",
		RunE: func(cmd *cobra.Command, args []string) error {
			return documentCRDs(cmd, args, options, log)
		},
	}

	options.configPath = cmd.Flags().StringP(
		"config",
		"c",
		"",
		"Path to a config file")

	options.outputPath = cmd.Flags().StringP(
		"output",
		"o",
		"",
		"Output path: file path for single-file mode, directory path for multiple-file mode")

	if err := cmd.MarkFlagRequired("output"); err != nil {
		return nil, errors.Wrap(err, "setting up --output")
	}

	options.templatePath = cmd.Flags().StringP(
		"template",
		"t",
		"",
		"Path to a folder containing templates to use for rendering the documentation")

	options.classDiagrams = cmd.Flags().BoolP(
		"class-diagrams",
		"",
		false,
		"Generate class diagrams for the CRDs")

	options.useGoFieldNames = cmd.Flags().BoolP(
		"use-go-field-names",
		"",
		false,
		"Use Go field names instead of serialized field names from JSON/YAML tags")

	options.fileMode = cmd.Flags().StringP(
		"file-mode",
		"f",
		"",
		"File mode: 'single-file' (default) or 'multiple-file'")

	options.includeFunctions = cmd.Flags().BoolP(
		"include-functions",
		"",
		false,
		"Include functions/methods in object documentation and class diagrams")

	return cmd, nil
}

type documentCRDsOptions struct {
	configPath       *string
	outputPath       *string
	templatePath     *string
	classDiagrams    *bool
	useGoFieldNames  *bool
	fileMode         *string
	includeFunctions *bool
}

func documentCRDs(
	cmd *cobra.Command,
	args []string,
	options *documentCRDsOptions,
	log logr.Logger,
) error {
	if err := options.validate(args); err != nil {
		return err
	}

	cfg, err := loadConfig(cmd, options)
	if err != nil {
		return errors.Wrap(err, "loading configuration")
	}

	packageFolder := args[0]
	loader := packageloader.New(cfg, log)

	pkg, err := loader.LoadDirectory(packageFolder)
	if err != nil {
		return errors.Wrapf(err, "loading package from %q", packageFolder)
	}

	return generateCrds(cfg, options, pkg, log)
}

func generateCrds(
	cfg *config.Config,
	options *documentCRDsOptions,
	pkg *model.Package,
	log logr.Logger,
) error {
	gen := generator.New(cfg, log)

	err := gen.LoadTemplates()
	if err != nil {
		return errors.Wrap(err, "loading templates")
	}

	// Choose generation method based on mode
	switch {
	case cfg.HasFileMode(config.FileModeSingleFile):
		err = gen.GenerateToFile(pkg, *options.outputPath, log)
		if err != nil {
			return errors.Wrapf(err, "generating output to %q", *options.outputPath)
		}
	case cfg.HasFileMode(config.FileModeMultipleFile):
		err = gen.GenerateToMultipleFiles(pkg, *options.outputPath, log)
		if err != nil {
			return errors.Wrapf(err, "generating multiple files to %q", *options.outputPath)
		}
	}

	return nil
}

func loadConfig(
	cmd *cobra.Command,
	options *documentCRDsOptions,
) (*config.Config, error) {
	cfg := config.Standard()

	// If we have a config file specified, load it
	if options.configPath != nil && *options.configPath != "" {
		err := cfg.Load(*options.configPath)
		if err != nil {
			return nil, errors.Wrapf(err, "reading config file %q", *options.configPath)
		}
	}

	// Apply overrides from the command line (if any)
	options.applyToConfig(cmd, cfg)

	// Check that our configuration is still valid
	err := cfg.Validate()
	if err != nil {
		return nil, errors.Wrap(err, "validating configuration")
	}

	return cfg, nil
}

func (options *documentCRDsOptions) validate(
	args []string,
) error {
	// Error if package directory missing
	if len(args) == 0 {
		return errors.New("no package directory, expected exactly one")
	}

	// Error if too many package directories supplied
	if len(args) > 1 {
		return errors.New("multiple package directories, expected exactly one")
	}

	// Error if the package directory doesn't exist
	packageFolder := args[0]
	if _, err := os.Stat(packageFolder); err != nil {
		if os.IsNotExist(err) {
			return errors.Errorf("package directory %q does not exist", packageFolder)
		}

		return errors.Wrapf(err, "checking package directory %q", packageFolder)
	}

	// Error if no output path supplied
	if options.outputPath == nil || *options.outputPath == "" {
		return errors.New("no output path supplied")
	}

	return nil
}

// applyToConfig applies options we've received on the command line to the config.
func (options *documentCRDsOptions) applyToConfig(cmd *cobra.Command, cfg *config.Config) {
	cfg.SetTemplatePath(options.templatePath)
	cfg.EnableClassDiagrams(options.classDiagrams)
	cfg.SetUseGoFieldNames(options.useGoFieldNames)
	cfg.SetFileMode(options.fileMode)

	// Only apply the includeFunctions flag if it was explicitly set by the user
	if cmd.Flags().Changed("include-functions") {
		cfg.SetIncludeFunctions(options.includeFunctions)
	}
}
