package cmd

import (
	"github.com/go-logr/logr"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

func Execute(log logr.Logger) error {
	root, err := newRootCommand(log)
	if err != nil {
		return errors.Wrapf(err, "failed to create root command")
	}

	if err := root.Execute(); err != nil {
		return errors.Wrapf(err, "failed to execute command")
	}

	return nil
}

var verbose bool

func newRootCommand(log logr.Logger) (*cobra.Command, error) {
	rootCmd := &cobra.Command{
		Use:              "crddoc",
		Short:            "crddoc is a generator for CRD documentation",
		TraverseChildren: true,
		SilenceErrors:    true, // We show errors ourselves using our logger
		SilenceUsage:     true, // Let users ask for usage themselves
	}

	rootCmd.PersistentFlags().BoolVar(&verbose, "verbose", false, "Enable verbose logging")

	rootCmd.Flags().SortFlags = false

	cmds := []func(logr.Logger) (*cobra.Command, error){
		newDocumentPackageCommand,
		newExportCommand,
	}

	for _, f := range cmds {
		cmd, err := f(log)
		if err != nil {
			return rootCmd, err
		}

		rootCmd.AddCommand(cmd)
	}

	rootCmd.PersistentPreRun = func(
		_ *cobra.Command,
		_ []string,
	) {
		if verbose {
			verboseLogging()
		}
	}

	return rootCmd, nil
}
