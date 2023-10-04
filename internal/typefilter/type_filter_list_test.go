package typefilter

import (
	"testing"

	"github.com/onsi/gomega"
	"github.com/theunrepentantgeek/crddoc/internal/config"
)

func TestTypeFilterList_Filter_WhenIncludeFilterMatches_ReturnsInclude(t *testing.T) {
	t.Parallel()
	g := gomega.NewWithT(t)

	// Arrange
	cfg := &config.Config{
		TypeFilters: []*config.Filter{
			{
				Include: "[0-9]+",
			},
		},
	}
	filters := New(cfg)

	// Act
	result := filters.Filter("12345")

	// Assert
	g.Expect(result).To(gomega.Equal(Included))
}

func TestTypeFilterList_Filter_WhenExcludeMatchesBeforeInclude_ReturnsExclude(t *testing.T) {
	t.Parallel()
	g := gomega.NewWithT(t)

	// Arrange
	cfg := &config.Config{
		TypeFilters: []*config.Filter{
			{
				Exclude: "Foo*",
			},
			{
				Include: "Foo*",
			},
		},
	}
	filters := New(cfg)

	// Act
	result := filters.Filter("Foot")

	// Assert
	g.Expect(result).To(gomega.Equal(Excluded))
}

func TestTypeFilterList_Filter_WhenIncludeMatchesBeforeExclude_ReturnsInclude(t *testing.T) {
	t.Parallel()
	g := gomega.NewWithT(t)

	// Arrange
	cfg := &config.Config{
		TypeFilters: []*config.Filter{
			{
				Include: "[0-9]+",
			},
			{
				Exclude: "[0-9]+",
			},
		},
	}
	filters := New(cfg)

	// Act
	result := filters.Filter("12345")

	// Assert
	g.Expect(result).To(gomega.Equal(Included))
}
