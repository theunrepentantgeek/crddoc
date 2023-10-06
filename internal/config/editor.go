package config

import (
	"fmt"
	"regexp"
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
}

func (edit *Editor) Validate() error {
	if edit.Context != "" {
		_, err := regexp.Compile(edit.Context)
		if err != nil {
			return fmt.Errorf("unable to compile 'context': %w", err)
		}
	}

	if edit.Search == "" {
		return fmt.Errorf("editor 'search' may not be empty")
	} else {
		_, err := regexp.Compile(edit.Search)
		if err != nil {
			return fmt.Errorf("unable to compile 'search': %w", err)
		}
	}

	return nil
}
