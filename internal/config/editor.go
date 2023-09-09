package config

import (
	"fmt"
	"regexp"

	"github.com/pkg/errors"
)

type Editor struct {
	Context       string `yaml:"context"` // Context is a regex identifying a substring to modify
	Search        string `yaml:"search"`  // Search is a regex identifying a substring to replace
	Replacement   string `yaml:"replace"` // Replace is the string to replace the search string with
	contextRegexp *regexp.Regexp
	searchRegexp  *regexp.Regexp
}

func (edit *Editor) Validate() error {
	if err := edit.ensureInitialized(); err != nil {
		return err
	}

	if edit.Search == "" {
		return fmt.Errorf("editor 'search' may not be empty")
	}

	return nil
}

func (edit *Editor) Replace(input string) string {
	if err := edit.ensureInitialized(); err != nil {
		// Should never happen under normal circumstances
		// because we'll Validate() first
		panic(err)
	}

	if edit.contextRegexp == nil {
		// Simple replace over the whole string
		return edit.searchRegexp.ReplaceAllString(input, edit.Replacement)
	}

	// Replace only in the context of the context regex
	return edit.contextRegexp.ReplaceAllStringFunc(
		input,
		func(context string) string {
			return edit.searchRegexp.ReplaceAllString(context, edit.Replacement)
		})
}

func (edit *Editor) ensureInitialized() error {
	if edit.Context != "" && edit.contextRegexp == nil {
		rx, err := regexp.Compile(edit.Context)
		if err != nil {
			return errors.Wrap(err, "editor unable to compile 'context':")
		}

		edit.contextRegexp = rx
	}

	if edit.Search != "" && edit.searchRegexp == nil {
		rx, err := regexp.Compile(edit.Search)
		if err != nil {
			return errors.Wrap(err, "editor unable to compile 'search':")
		}

		edit.searchRegexp = rx
	}

	return nil
}
