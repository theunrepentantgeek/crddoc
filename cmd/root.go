package cmd

import (
	"os"

	"github.com/go-logr/logr"
	"github.com/spf13/cobra"
)

func Execute() {
	log := CreateLogger()
	root, err := newRootCommand(log)
	if err != nil {
		log.Error(err, "failed to create root command")
		os.Exit(1)
	}

	if err := root.Execute(); err != nil {
		log.Error(err, "failed to execute command")
		os.Exit(1)
	}
}

var (
	verbose bool
)

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
	}

	for _, f := range cmds {
		cmd, err := f(log)
		if err != nil {
			return rootCmd, err
		}

		rootCmd.AddCommand(cmd)
	}

	rootCmd.PersistentPreRun = func(cmd *cobra.Command, args []string) {
		if verbose {
			verboseLogging()
		}
	}

	return rootCmd, nil
}
