package cmd

import (
	"bufio"
	"context"
	"os"

	"github.com/pkg/errors"

	"github.com/go-logr/logr"
	"github.com/spf13/cobra"
	"github.com/theunrepentantgeek/crddoc/internal/config"
	"github.com/theunrepentantgeek/crddoc/internal/generator"
	"github.com/theunrepentantgeek/crddoc/internal/model"
)

func newDocumentPackageCommand(log logr.Logger) (*cobra.Command, error) {
	options := &packageCommandOptions{}

	cmd := &cobra.Command{
		Use:   "document-package",
		Short: "Generate documentation for a package",
		Long:  "Generate documentation for a package.",
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx := cmd.Context()
			return documentPackage(ctx, args, options, log)
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

	return cmd, nil
}

type packageCommandOptions struct {
	configPath *string
	outputPath *string
}

func documentPackage(
	ctx context.Context,
	args []string,
	options *packageCommandOptions,
	log logr.Logger,
) error {
	if err := validateDocumentPackage(args, options); err != nil {
		return err
	}

	// Load the config file
	cfg := config.Default()
	if options.configPath != nil || *options.configPath != "" {
		err := cfg.Load(*options.configPath)
		if err != nil {
			return errors.Wrapf(err, "reading config file %q", *options.configPath)
		}

		err = cfg.Validate()
		if err != nil {
			return errors.Wrapf(err, "validating config file %q", *options.configPath)
		}
	}

	pkg := model.NewPackage(cfg, log)
	packageFolder := args[0]
	err := pkg.LoadDirectory(packageFolder)
	if err != nil {
		return errors.Wrapf(err, "loading package from %q", packageFolder)
	}

	gen := generator.New(cfg, log)
	templateFolder := "C:\\GitHub\\crddoc\\templates\\crd"
	err = gen.LoadTemplates(templateFolder)
	if err != nil {
		return errors.Wrapf(err, "loading templates from %q", templateFolder)
	}

	// Render the template and write to output
	f, err := os.Create(*options.outputPath)
	if err != nil {
		return errors.Wrapf(err, "creating output file %q", *options.outputPath)
	}

	defer f.Close()

	log.Info("Writing to", "outputPath", *options.outputPath)

	w := bufio.NewWriter(f)

	err = gen.Generate(pkg, w)
	if err != nil {
		return errors.Wrapf(err, "generating output to %q", *options.outputPath)
	}

	w.Flush()
	return nil
}

func validateDocumentPackage(
	args []string,
	options *packageCommandOptions,
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
