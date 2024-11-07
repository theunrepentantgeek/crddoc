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

type Result string

const (
	Included Result = "include" // Type Filter includes the type
	Excluded Result = "exclude" // Type Filter excludes the type
	Skipped  Result = "none"    // Type Filter does not apply to the type
)

func NewTypeFilter(cfg *config.Filter) *TypeFilter {
	return &TypeFilter{
		cfg: cfg,
	}
}

func (f *TypeFilter) Applies(name string) Result {
	f.ensureCompiled()

	if f.includeRegex != nil && f.includeRegex.MatchString(name) {
		return Included
	}

	if f.excludeRegex != nil && f.excludeRegex.MatchString(name) {
		return Excluded
	}

	return Skipped
}

func (f *TypeFilter) ensureCompiled() {
	if f.cfg.Include != "" && f.includeRegex == nil {
		f.includeRegex = createGlobber(f.cfg.Include)
	}

	if f.cfg.Exclude != "" && f.excludeRegex == nil {
		f.excludeRegex = createGlobber(f.cfg.Exclude)
	}
}
