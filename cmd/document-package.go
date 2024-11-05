package cmd

import (
	"os"

	"github.com/go-logr/logr"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"

	"github.com/theunrepentantgeek/crddoc/internal/config"
	"github.com/theunrepentantgeek/crddoc/internal/generator"
	"github.com/theunrepentantgeek/crddoc/internal/packageloader"
)

func newDocumentPackageCommand(log logr.Logger) (*cobra.Command, error) {
	options := &documentPackageOptions{}

	cmd := &cobra.Command{
		Use:   "document-package",
		Short: "Generate documentation for a package",
		Long:  "Generate documentation for a package.",
		RunE: func(_ *cobra.Command, args []string) error {
			return documentPackage(args, options, log)
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
		"Write ARM resource CRDs to a single file")

	options.templatePath = cmd.Flags().StringP(
		"template",
		"t",
		"",
		"Path to a folder containing templates to use for rendering the documentation")

	return cmd, nil
}

type documentPackageOptions struct {
	configPath   *string
	outputPath   *string
	templatePath *string
}

func documentPackage(
	args []string,
	options *documentPackageOptions,
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

	gen := generator.New(cfg, log)

	err = gen.LoadTemplates()
	if err != nil {
		return errors.Wrap(err, "loading templates")
	}

	err = gen.GenerateToFile(pkg, *options.outputPath, log)
	if err != nil {
		return errors.Wrapf(err, "generating output to %q", *options.outputPath)
	}

	return nil
}

func loadConfig(
	options *documentPackageOptions,
) (*config.Config, error) {
	cfg := config.Default()

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

func (options *documentPackageOptions) validate(
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
func (options *documentPackageOptions) applyToConfig(cfg *config.Config) {
	cfg.OverrideTemplatePath(options.templatePath)
}
