package config

import (
	"testing"

	. "github.com/onsi/gomega"
)

func TestEditor_Validate_WhenContextIsInvalidRegex_ReturnsExpectedError(t *testing.T) {
	t.Parallel()

	g := NewGomegaWithT(t)

	// Arrange
	edit := &Editor{
		Context: "[A-Z",
		Search:  "[A-Za-z0-9]+",
	}

	// Act
	err := edit.Validate()

	// Assert
	g.Expect(err).To(HaveOccurred())
	g.Expect(err).To(MatchError(ContainSubstring("compiling editor context expression")))
}

func TestEditor_WhenSearchIsMissing_ReturnsExpectedError(t *testing.T) {
	t.Parallel()

	g := NewGomegaWithT(t)

	// Arrange
	edit := &Editor{}

	// Act
	err := edit.Validate()

	// Assert
	g.Expect(err).To(HaveOccurred())
	g.Expect(err).To(MatchError(ContainSubstring("editor search expression may not be empty")))
}

func TestEditor_Validate_WhenSearchIsInvalidRegex_ReturnsExpectedError(t *testing.T) {
	t.Parallel()

	g := NewGomegaWithT(t)

	// Arrange
	edit := &Editor{
		Search: "[A-Z",
	}

	// Act
	err := edit.Validate()

	// Assert
	g.Expect(err).To(HaveOccurred())
	g.Expect(err).To(MatchError(ContainSubstring("compiling editor search expression")))
}
