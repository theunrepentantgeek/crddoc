package config

import (
	"testing"

	"github.com/onsi/gomega"
)

func TestFilter_Validate_WhenBlank_ReturnsError(t *testing.T) {
	t.Parallel()
	g := gomega.NewWithT(t)

	// Arrange
	f := &Filter{}

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
