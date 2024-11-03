package typefilter

import "github.com/theunrepentantgeek/crddoc/internal/config"

type List struct {
	filters []*TypeFilter
}

func New(cfg *config.Config) *List {
	filters := make([]*TypeFilter, 0, len(cfg.TypeFilters))
	for _, f := range cfg.TypeFilters {
		filters = append(filters, NewTypeFilter(f))
	}

	return &List{
		filters: filters,
	}
}

func (c *List) Filter(name string) Result {
	for _, f := range c.filters {
		if result := f.Applies(name); result != Skipped {
			return result
		}
	}

	return Included
}
