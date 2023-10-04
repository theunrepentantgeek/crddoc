package typefilter

import (
	"regexp"

	"github.com/theunrepentantgeek/crddoc/internal/config"
)

type TypeFilter struct {
	cfg *config.Filter

	excludeRegex *regexp.Regexp
	includeRegex *regexp.Regexp
}

type TypeFilterResult string

const (
	Included TypeFilterResult = "include" // Type Filter includes the type
	Excluded TypeFilterResult = "exclude" // Type Filter excludes the type
	Skipped  TypeFilterResult = "none"    // Type Filter does not apply to the type
)

func NewTypeFilter(cfg *config.Filter) *TypeFilter {
	return &TypeFilter{
		cfg: cfg,
	}
}

func (f *TypeFilter) Applies(name string) TypeFilterResult {
	// ensure our include and exclude filters are compiled
	if f.cfg.Include != "" && f.includeRegex == nil {
		f.includeRegex = createGlobber(f.cfg.Include)
	}

	if f.cfg.Exclude != "" && f.excludeRegex == nil {
		f.excludeRegex = createGlobber(f.cfg.Exclude)
	}

	if f.includeRegex != nil && f.includeRegex.MatchString(name) {
		return Included
	}

	if f.excludeRegex != nil && f.excludeRegex.MatchString(name) {
		return Excluded
	}

	return Skipped
}
