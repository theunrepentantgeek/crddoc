package typefilter

import (
	"testing"

	"github.com/onsi/gomega"
	"github.com/theunrepentantgeek/crddoc/internal/config"
)

func TestTypeFilter_Applies_WhenExcludeMatchedExactCase_ReturnsExcluded(t *testing.T) {
	t.Parallel()
	g := gomega.NewWithT(t)

	// Arrange
	cfg := &config.Filter{
		Exclude: "Foo*",
	}
	f := NewTypeFilter(cfg)

	// Act
	result := f.Applies("Foot")

	// Assert
	g.Expect(result).To(gomega.Equal(Excluded))
}

func TestTypeFilter_Applies_WhenExcludeMatchedDifferentCase_ReturnsExcluded(t *testing.T) {
	t.Parallel()
	g := gomega.NewWithT(t)

	// Arrange
	cfg := &config.Filter{
		Exclude: "Foo*",
	}
	f := NewTypeFilter(cfg)

	// Act
	result := f.Applies("foot")

	// Assert
	g.Expect(result).To(gomega.Equal(Excluded))
}

func TestTypeFilter_Applies_WhenExcludedUnmatched_ReturnsNone(t *testing.T) {
	t.Parallel()
	g := gomega.NewWithT(t)

	// Arrange
	cfg := &config.Filter{
		Exclude: "Foo*",
	}
	f := NewTypeFilter(cfg)

	// Act
	result := f.Applies("abcde")

	// Assert
	g.Expect(result).To(gomega.Equal(Skipped))
}

func TestTypeFilter_Applies_WhenIncludedMatchedExactCase_ReturnsIncluded(t *testing.T) {
	t.Parallel()
	g := gomega.NewWithT(t)

	// Arrange
	cfg := &config.Filter{
		Include: "Foo*",
	}
	f := NewTypeFilter(cfg)

	// Act
	result := f.Applies("Foot")

	// Assert
	g.Expect(result).To(gomega.Equal(Included))
}

func TestTypeFilter_Applies_WhenIncludedMatchedDifferentCase_ReturnsIncluded(t *testing.T) {
	t.Parallel()
	g := gomega.NewWithT(t)

	// Arrange
	cfg := &config.Filter{
		Include: "Foo*",
	}
	f := NewTypeFilter(cfg)

	// Act
	result := f.Applies("foot")

	// Assert
	g.Expect(result).To(gomega.Equal(Included))
}

func TestTypeFilter_Applies_WhenIncludedUnmatched_ReturnsNone(t *testing.T) {
	t.Parallel()
	g := gomega.NewWithT(t)

	// Arrange
	cfg := &config.Filter{
		Exclude: "Foo*",
	}
	f := NewTypeFilter(cfg)

	// Act
	result := f.Applies("Arm")

	// Assert
	g.Expect(result).To(gomega.Equal(Skipped))
}
