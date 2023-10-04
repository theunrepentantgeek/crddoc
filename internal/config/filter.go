package config

import (
	"github.com/pkg/errors"
)

type Filter struct {
	// Because is an explanation of why this filter is being applied and what it does.
	Because string `yaml:"because"`

	// Exclude is a glob identifying types to exclude from the output.
	Exclude string `yaml:"exclude"`

	// Include is a glob identifying types to include in the output.
	Include string `yaml:"include"`
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
