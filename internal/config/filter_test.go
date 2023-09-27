package config

import (
	"testing"

	"github.com/onsi/gomega"
)

func TestFilter_Validate_WhenBlank_ReturnsError(t *testing.T) {
	t.Parallel()
	g := gomega.NewWithT(t)

	// Arrange
	f := NewFilter()

	// Act
	err := f.validate()

	// Assert
	g.Expect(err).To(gomega.HaveOccurred())
	g.Expect(err.Error()).To(gomega.ContainSubstring("must specify either"))
}

func TestFilter_Validate_WhenBothIncludeAndExcludeSpecified_ReturnsError(t *testing.T) {
	t.Parallel()
	g := gomega.NewWithT(t)

	// Arrange
	f := NewFilter()
	f.Exclude = "foo"
	f.Include = "bar"

	// Act
	err := f.validate()

	// Assert
	g.Expect(err).To(gomega.HaveOccurred())
	g.Expect(err.Error()).To(gomega.ContainSubstring("cannot specify both"))
}

func TestFilter_Applies_WhenExcludeMatchedExactCase_ReturnsExcluded(t *testing.T) {
	t.Parallel()
	g := gomega.NewWithT(t)

	// Arrange
	f := NewFilter()
	f.Exclude = "Foo*"
	g.Expect(f.validate()).To(gomega.Succeed())

	// Act
	result := f.Applies("Foot")

	// Assert
	g.Expect(result).To(gomega.Equal(FilterResultExclude))
}

func TestFilter_Applies_WhenExcludeMatchedDifferentCase_ReturnsExcluded(t *testing.T) {
	t.Parallel()
	g := gomega.NewWithT(t)

	// Arrange
	f := NewFilter()
	f.Exclude = "Foo*"
	g.Expect(f.validate()).To(gomega.Succeed())

	// Act
	result := f.Applies("foot")

	// Assert
	g.Expect(result).To(gomega.Equal(FilterResultExclude))
}

func TestFilter_Applies_WhenExcludedUnmatched_ReturnsNone(t *testing.T) {
	t.Parallel()
	g := gomega.NewWithT(t)

	// Arrange
	f := NewFilter()
	f.Exclude = "Fo*"

	// Act
	result := f.Applies("abcde")

	// Assert
	g.Expect(result).To(gomega.Equal(FilterResultNone))
}

func TestFilter_Applies_WhenIncludedMatchedExactCase_ReturnsIncluded(t *testing.T) {
	t.Parallel()
	g := gomega.NewWithT(t)

	// Arrange
	f := NewFilter()
	f.Include = "Fo*"

	// Act
	result := f.Applies("Foot")

	// Assert
	g.Expect(result).To(gomega.Equal(FilterResultInclude))
}

func TestFilter_Applies_WhenIncludedMatchedDifferentCase_ReturnsIncluded(t *testing.T) {
	t.Parallel()
	g := gomega.NewWithT(t)

	// Arrange
	f := NewFilter()
	f.Include = "Fo*"

	// Act
	result := f.Applies("foot")

	// Assert
	g.Expect(result).To(gomega.Equal(FilterResultInclude))
}

func TestFilter_Applies_WhenIncludedUnmatched_ReturnsNone(t *testing.T) {
	t.Parallel()
	g := gomega.NewWithT(t)

	// Arrange
	f := NewFilter()
	f.Include = "Fo*"

	// Act
	result := f.Applies("Arm")

	// Assert
	g.Expect(result).To(gomega.Equal(FilterResultNone))
}
