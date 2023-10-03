package config

import (
	"fmt"
	"regexp"

	"github.com/pkg/errors"
)

// Editor represents a point modification to make to exported documentation
type Editor struct {
	// Context is a regex identifying a substring to modify, limiting the scope of the search and
	// replace. If omitted, the entire string is eligible for modification.
	Context string `yaml:"context"`

	// Search is a regex identifying a substring to replace.
	Search string `yaml:"search"`

	// Replace is the string to substitute for the search regex.
	Replacement string `yaml:"replace"`

	contextRegexp *regexp.Regexp
	searchRegexp  *regexp.Regexp
}

func (edit *Editor) Validate() error {
	if edit.Search == "" {
		return fmt.Errorf("editor 'search' may not be empty")
	}

	if err := edit.ensureInitialized(); err != nil {
		return err
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
