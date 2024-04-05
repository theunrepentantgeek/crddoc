package cmd

import (
	"github.com/go-logr/logr"
	"github.com/spf13/cobra"
)

func newExportCommand(log logr.Logger) (*cobra.Command, error) {
	exportCmd := &cobra.Command{
		Use:   "export",
		Short: "Export templates or configuration",
		Long:  "Generate documentation for a package.",
	}

	cmds := []func(logr.Logger) (*cobra.Command, error){
		newExportTemplatesCommand,
		newExportConfigurationCommand,
	}

	for _, f := range cmds {
		cmd, err := f(log)
		if err != nil {
			return cmd, err
		}

		exportCmd.AddCommand(cmd)
	}

	return exportCmd, nil
}
