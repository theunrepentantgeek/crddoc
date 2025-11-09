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

//nolint:funlen // command setup requires multiple steps
func newDocumentCRDsCommand(log logr.Logger) (*cobra.Command, error) {
	// Use local variables to capture command line flags
	var (
		configPath       string
		outputPath       string
		templatePath     string
		classDiagrams    bool
		useGoFieldNames  bool
		fileMode         string
		includeFunctions bool
	)

	cmd := &cobra.Command{
		Use:   "crds",
		Short: "Generate CRD documentation from a package",
		Long:  "Generate CRD documentation from a package.",
		RunE: func(cmd *cobra.Command, args []string) error {
			// Initialize options within the RunE lambda
			options := &documentCRDsOptions{
				configPath:   &configPath,
				outputPath:   &outputPath,
				templatePath: &templatePath,
				fileMode:     &fileMode,
			}

			// Use Changed() method to decide whether to populate options
			if cmd.Flags().Changed("class-diagrams") {
				options.classDiagrams = &classDiagrams
			}

			if cmd.Flags().Changed("use-go-field-names") {
				options.useGoFieldNames = &useGoFieldNames
			}

			if cmd.Flags().Changed("include-functions") {
				options.includeFunctions = &includeFunctions
			}

			return documentCRDs(args, options, log)
		},
	}

	cmd.Flags().StringVarP(
		&configPath,
		"config",
		"c",
		"",
		"Path to a config file")

	cmd.Flags().StringVarP(
		&outputPath,
		"output",
		"o",
		"",
		"Output path: file path for single-file mode, directory path for multiple-file mode")

	if err := cmd.MarkFlagRequired("output"); err != nil {
		return nil, errors.Wrap(err, "setting up --output")
	}

	cmd.Flags().StringVarP(
		&templatePath,
		"template",
		"t",
		"",
		"Path to a folder containing templates to use for rendering the documentation")

	cmd.Flags().BoolVarP(
		&classDiagrams,
		"class-diagrams",
		"",
		false,
		"Generate class diagrams for the CRDs")

	cmd.Flags().BoolVarP(
		&useGoFieldNames,
		"use-go-field-names",
		"",
		false,
		"Use Go field names instead of serialized field names from JSON/YAML tags")

	cmd.Flags().StringVarP(
		&fileMode,
		"file-mode",
		"f",
		"",
		"File mode: 'single-file' (default) or 'multiple-file'")

	cmd.Flags().BoolVarP(
		&includeFunctions,
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
	args []string,
	options *documentCRDsOptions,
	log logr.Logger,
) error {
	if err := options.validate(args); err != nil {
		return err
	}

	cfg, err := loadConfig(options)
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
	options.applyToConfig(cfg)

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
func (options *documentCRDsOptions) applyToConfig(cfg *config.Config) {
	cfg.SetTemplatePath(options.templatePath)
	cfg.SetFileMode(options.fileMode)

	// Only apply boolean flags if they were explicitly set by the user
	// This allows config file values to be preserved when flags aren't specified
	cfg.EnableClassDiagrams(options.classDiagrams)
	cfg.SetUseGoFieldNames(options.useGoFieldNames)
	cfg.SetIncludeFunctions(options.includeFunctions)
}
