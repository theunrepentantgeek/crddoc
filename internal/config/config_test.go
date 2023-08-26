package config

import (
	"testing"

	. "github.com/onsi/gomega"
)

func TestConfig_Filter_WhenNoFilters_ReturnsInclude(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	// Arrange
	c := &Config{}

	// Act
	result := c.Filter("foo")

	// Assert
	g.Expect(result).To(Equal(FilterResultInclude))
}

func TestConfig_Filter_WhenExcludeFilterMatches_ReturnsExclude(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	// Arrange
	c := &Config{
		TypeFilters: []*Filter{
			{
				Exclude: "Foo*",
			},
		},
	}

	// Act
	result := c.Filter("FooBar")

	// Assert
	g.Expect(result).To(Equal(FilterResultExclude))
}

func TestConfig_Filter_WhenIncludeFilterMatches_ReturnsInclude(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	// Arrange
	c := &Config{
		TypeFilters: []*Filter{
			{
				Include: "[0-9]+",
			},
		},
	}

	// Act
	result := c.Filter("12345")

	// Assert
	g.Expect(result).To(Equal(FilterResultInclude))
}

func TestConfig_Filter_WhenExcludeMatchesBeforeInclude_ReturnsExclude(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	// Arrange
	c := &Config{
		TypeFilters: []*Filter{
			{
				Exclude: "Foo*",
			},
			{
				Include: "Foo*",
			},
		},
	}

	// Act
	result := c.Filter("Foot")

	// Assert
	g.Expect(result).To(Equal(FilterResultExclude))
}

func TestConfig_Filter_WhenIncludeMatchesBeforeExclude_ReturnsInclude(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	// Arrange
	c := &Config{
		TypeFilters: []*Filter{
			{
				Include: "[0-9]+",
			},
			{
				Exclude: "[0-9]+",
			},
		},
	}

	// Act
	result := c.Filter("12345")

	// Assert
	g.Expect(result).To(Equal(FilterResultInclude))
}
