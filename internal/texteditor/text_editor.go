package texteditor

import (
	"regexp"

	"github.com/pkg/errors"
	"github.com/theunrepentantgeek/crddoc/internal/config"
)

// Editor represents a point modification to make to exported documentation
type Editor struct {
	cfg config.Editor

	contextRegexp *regexp.Regexp
	searchRegexp  *regexp.Regexp
}

func NewEditor(cfg config.Editor) (*Editor, error) {
	result := &Editor{
		cfg: cfg,
	}

	if cfg.Context != "" {
		rx, err := regexp.Compile(cfg.Context)
		if err != nil {
			return nil, errors.Wrap(err, "editor unable to compile 'context':")
		}

		result.contextRegexp = rx
	}

	if cfg.Search != "" {
		rx, err := regexp.Compile(cfg.Search)
		if err != nil {
			return nil, errors.Wrap(err, "editor unable to compile 'search':")
		}

		result.searchRegexp = rx
	}

	return result, nil
}

func (edit *Editor) Replace(input string) string {
	if edit.contextRegexp == nil {
		// Simple replace over the whole string
		return edit.searchRegexp.ReplaceAllString(input, edit.cfg.Replacement)
	}

	// Replace only in the context of the context regex
	return edit.contextRegexp.ReplaceAllStringFunc(
		input,
		func(context string) string {
			return edit.searchRegexp.ReplaceAllString(context, edit.cfg.Replacement)
		})
}
