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
	g.Expect(err.Error()).To(gomega.ContainSubstring("compiling editor context expression"))
	g.Expect(err).To(HaveOccurred())
}

func TestEditor_WhenSearchIsMissing_ReturnsExpectedError(t *testing.T) {
	t.Parallel()

	g := NewGomegaWithT(t)

	// Arrange
	edit := &Editor{}

	// Act
	err := edit.Validate()

	// Assert
	g.Expect(err.Error()).To(gomega.ContainSubstring("editor search expression may not be empty"))
	g.Expect(err).To(HaveOccurred())
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
	g.Expect(err.Error()).To(gomega.ContainSubstring("compiling editor search expression"))
	g.Expect(err).To(HaveOccurred())
}
