package cmd

import (
	"io"
	"io/fs"
	"os"
	"path/filepath"

	"github.com/go-logr/logr"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"

	"github.com/theunrepentantgeek/crddoc/templates"
)

func newExportTemplatesCommand(
	log logr.Logger,
) (*cobra.Command, error) {
	options := &exportTemplateOptions{}

	cmd := &cobra.Command{
		Use:   "templates",
		Short: "Export standard templates to a folder",
		Long:  "Export the templates contained within crddoc to a folder for customization.",
		RunE: func(_ *cobra.Command, args []string) error {
			return exportTemplates(args, options, log)
		},
	}

	options.folder = cmd.Flags().StringP(
		"folder",
		"f",
		"",
		"Path to a folder into which templates will be exported")

	if err := cmd.MarkFlagRequired("folder"); err != nil {
		return nil, errors.Wrap(err, "setting up --folder")
	}

	return cmd, nil
}

type exportTemplateOptions struct {
	folder *string
}

func exportTemplates(
	_ []string,
	options *exportTemplateOptions,
	log logr.Logger,
) error {
	if err := options.validate(); err != nil {
		return errors.Wrapf(err, "invalid options")
	}

	// Iterate through all the files found in templates.CRD and write them to the
	// specified folder.
	err := fs.WalkDir(
		templates.CRD,
		".",
		exportTemplateFiles(templates.CRD, *options.folder, log))
	if err != nil {
		return errors.Wrap(err, "exporting templates")
	}

	return nil
}

func exportTemplateFiles(
	template fs.FS,
	folder string,
	log logr.Logger,
) fs.WalkDirFunc {
	return func(
		path string,
		entry fs.DirEntry,
		_ error,
	) error {
		if entry.IsDir() {
			dir := filepath.Join(folder, path)

			return os.MkdirAll(dir, 0755)
		}

		log.Info(
			"Exporting template file",
			"file", path,
			"folder", folder)

		f, err := template.Open(path)
		if err != nil {
			return errors.Wrap(err, "opening source file")
		}
		defer f.Close()

		return exportTemplateFile(f, filepath.Join(folder, path))
	}
}

func exportTemplateFile(
	src fs.File,
	dst string,
) error {
	file, err := os.Create(dst)
	if err != nil {
		return errors.Wrap(err, "creating destination file")
	}
	defer file.Close()

	_, err = io.Copy(file, src)
	if err != nil {
		return errors.Wrap(err, "copying file")
	}

	return nil
}

func (options *exportTemplateOptions) validate() error {
	if options.folder == nil || *options.folder == "" {
		return errors.New("no export folder supplied")
	}

	return nil
}
