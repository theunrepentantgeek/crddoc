package cmd

import (
	"context"

	"github.com/go-logr/logr"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

func newExportTemplatesCommand(
	log logr.Logger,
) (*cobra.Command, error) {
	options := &exportTemplateOptions{}

	cmd := &cobra.Command{
		Use:   "templates",
		Short: "Export standard templates to a folder",
		Long:  "Export the templates contained within crddoc to a folder for customization.",
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx := cmd.Context()
			return exportTemplates(ctx, args, options, log)
		},
	}

	options.folder = cmd.Flags().StringP(
		"folder",
		"f",
		"",
		"Path to a folder into which templates will be exported")

	return cmd, nil
}

type exportTemplateOptions struct {
	folder *string
}

func exportTemplates(
	ctx context.Context,
	args []string,
	options *exportTemplateOptions,
	log logr.Logger,
) error {
	if err := options.validate(); err != nil {
		return err
	}

	return nil
}

func (options *exportTemplateOptions) validate() error {
	if options.folder == nil || *options.folder == "" {
		return errors.New("no export folder supplied")
	}

	return nil
}
