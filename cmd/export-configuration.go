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

	return cmd, nil
}

// exportConfigurationOptions defines any optional parameters for template export.
// Currently there are none.
type exportConfigurationOptions struct{}

func exportConfiguration(
	args []string,
	_ *exportConfigurationOptions,
	log logr.Logger,
) error {
	if len(args) == 0 {
		return errors.New("no export file supplied")
	}

	if len(args) > 1 {
		return errors.New("too many export filenames supplied")
	}

	// Create our standard configuration
	cfg := config.Standard()

	// Save the configuration to a file as yaml
	file := args[0]
	log.Info("Exporting default configuration", "file", file)

	if err := cfg.Save(file); err != nil {
		return errors.Wrapf(err, "exporting configuration to %s", file)
	}

	return nil
}
