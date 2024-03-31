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
		Long:  "Export the default configuration to a file for customization.",
		RunE: func(cmd *cobra.Command, args []string) error {
			return exportConfiguration(args, options)
		},
	}

	return cmd, nil
}

type exportConfigurationOptions struct{}

func exportConfiguration(
	args []string,
	_ *exportConfigurationOptions,
) error {
	if len(args) == 0 {
		return errors.New("no export file supplied")
	}

	if len(args) > 1 {
		return errors.New("too many export filenames supplied")
	}

	// Create our default configuration
	cfg := config.Default()

	// Save the configuration to a file as yaml
	return cfg.Save(args[0])
}
