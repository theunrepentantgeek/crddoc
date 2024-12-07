package cmd

import (
	"github.com/go-logr/logr"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"

	"github.com/theunrepentantgeek/crddoc/internal/config"
)

func newExportConfigurationCommand(
	log logr.Logger,
) (*cobra.Command, error) {
	options := &exportConfigurationOptions{}
	cmd := &cobra.Command{
		Use:   "configuration",
		Short: "Export standard configuration to a file",
		Long:  "Export default configuration to a file for customization.",
		RunE: func(_ *cobra.Command, args []string) error {
			return exportConfiguration(args, options, log)
		},
	}

	options.outputFile = cmd.Flags().StringP(
		"output",
		"o",
		"",
		"Output file into which configuration will be exported.")

	if err := cmd.MarkFlagRequired("output"); err != nil {
		return nil, errors.Wrap(err, "setting up --output")
	}

	return cmd, nil
}

// exportConfigurationOptions defines any optional parameters for template export.
// Currently there are none.
type exportConfigurationOptions struct {
	outputFile *string
}

func exportConfiguration(
	args []string,
	options *exportConfigurationOptions,
	log logr.Logger,
) error {
	if len(args) > 0 {
		return errors.New("Extra parameters provided")
	}

	// Create our default configuration
	cfg := config.Default()

	// Save the configuration to a file as yaml
	file := *options.outputFile
	log.Info("Exporting default configuration", "file", file)

	if err := cfg.Save(file); err != nil {
		return errors.Wrapf(err, "exporting configuration to %s", file)
	}

	return nil
}
