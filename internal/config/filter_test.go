package config

import (
	"testing"

	. "github.com/onsi/gomega"
)

func TestFilter_Validate_WhenBlank_ReturnsError(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	// Arrange
	f := &Filter{}

	// Act
	err := f.validate()

	// Assert
	g.Expect(err).To(HaveOccurred())
	g.Expect(err).To(MatchError(ContainSubstring("must specify either")))
}

func TestFilter_Validate_WhenBothIncludeAndExcludeSpecified_ReturnsError(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	// Arrange
	f := &Filter{
		Exclude: "foo",
		Include: "bar",
	}

	// Act
	err := f.validate()

	// Assert
	g.Expect(err).To(HaveOccurred())
	g.Expect(err).To(MatchError(ContainSubstring("cannot specify both")))
}
