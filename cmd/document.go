package cmd

import (
	"github.com/go-logr/logr"
	"github.com/spf13/cobra"
)

func newDocumentCommand(log logr.Logger) (*cobra.Command, error) {
	documentCmd := &cobra.Command{
		Use:   "document",
		Short: "Generate documentation",
		Long:  "Generate documentation",
	}

	cmds := []func(logr.Logger) (*cobra.Command, error){
		newDocumentCRDsCommand,
	}

	for _, f := range cmds {
		cmd, err := f(log)
		if err != nil {
			return cmd, err
		}

		documentCmd.AddCommand(cmd)
	}

	return documentCmd, nil
}
