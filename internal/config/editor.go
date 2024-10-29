package config

import (
	"fmt"
	"regexp"

	"github.com/pkg/errors"
)

// Editor represents a point modification to make to exported documentation.
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
			return errors.Wrap(err, "compiling editor context expression")
		}
	}

	if edit.Search == "" {
		return fmt.Errorf("editor search expression may not be empty")
	}

	_, err := regexp.Compile(edit.Search)
	if err != nil {
		return errors.Wrap(err, "compiling editor search expression")
	}

	return nil
}
