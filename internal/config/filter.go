package config

import (
	"regexp"

	"github.com/pkg/errors"
)

type Filter struct {
	Because      string `yaml:"because"`
	Exclude      string `yaml:"exclude"`
	excludeRegex *regexp.Regexp
	Include      string `yaml:"include"`
	includeRegex *regexp.Regexp


}

func (f *Filter) validate() error {
	// Check we only have one kind of filter
	if f.Exclude != "" && f.Include != "" {
		return errors.Errorf(
			"cannot specify both include filter %s and exclude filter %s",
			f.Include,
			f.Exclude)
	}

	// Check that we have at least one kind of filter
	if f.Exclude == "" && f.Include == "" {
		return errors.New("must specify either include or exclude filter")
	}

	return nil
}
