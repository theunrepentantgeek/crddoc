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

type FilterResult string

const (
	FilterResultInclude FilterResult = "include"
	FilterResultExclude FilterResult = "exclude"
	FilterResultNone    FilterResult = "none"
)

func NewFilter() *Filter {
	return &Filter{}
}

func (f *Filter) Applies(name string) FilterResult {
	// ensure our include and exclude filters are compiled
	if f.Include != "" && f.includeRegex == nil {
		f.includeRegex = createGlobber(f.Include)
	}

	if f.Exclude != "" && f.excludeRegex == nil {
		f.excludeRegex = createGlobber(f.Exclude)
	}

	if f.includeRegex != nil && f.includeRegex.MatchString(name) {
		return FilterResultInclude
	}

	if f.excludeRegex != nil && f.excludeRegex.MatchString(name) {
		return FilterResultExclude
	}

	return FilterResultNone
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

	// If we have an Exclude filter, compile it
	if f.Exclude != "" {
		f.excludeRegex = createGlobber(f.Exclude)
	}

	// If we have an Include filter, compile it
	if f.Include != "" {
		f.includeRegex = createGlobber(f.Include)
	}

	return nil
}
