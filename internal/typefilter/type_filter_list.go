package typefilter

import "github.com/theunrepentantgeek/crddoc/internal/config"

type TypeFilterList struct {
	filters []*TypeFilter
}

func New(cfg *config.Config) *TypeFilterList {
	filters := make([]*TypeFilter, 0, len(cfg.TypeFilters))
	for _, f := range cfg.TypeFilters {
		filters = append(filters, NewTypeFilter(f))
	}

	return &TypeFilterList{
		filters: filters,
	}
}

func (c *TypeFilterList) Filter(name string) TypeFilterResult {
	for _, f := range c.filters {
		if result := f.Applies(name); result != Skipped {
			return result
		}
	}

	return Included
}
