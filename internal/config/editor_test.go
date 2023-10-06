package config

import (
	"testing"

	"github.com/onsi/gomega"
)

func TestEditor_Validate_WhenContextIsInvalidRegex_ReturnsExpectedError(t *testing.T) {
	t.Parallel()

	g := gomega.NewWithT(t)

	// Arrange
	edit := &Editor{
		Context: "[A-Z",
		Search:  "[A-Za-z0-9]+",
	}

	// Act
	err := edit.Validate()

	// Assert
	g.Expect(err).To(gomega.HaveOccurred())
	g.Expect(err.Error()).To(gomega.ContainSubstring("unable to compile 'context':"))
}

func TestEditor_WhenSearchIsMissing_ReturnsExpectedError(t *testing.T) {
	t.Parallel()

	g := gomega.NewWithT(t)

	// Arrange
	edit := &Editor{}

	// Act
	err := edit.Validate()

	// Assert
	g.Expect(err).To(gomega.HaveOccurred())
	g.Expect(err.Error()).To(gomega.ContainSubstring("editor 'search' may not be empty"))
}

func TestEditor_Validate_WhenSearchIsInvalidRegex_ReturnsExpectedError(t *testing.T) {
	t.Parallel()

	g := gomega.NewWithT(t)

	// Arrange
	edit := &Editor{
		Search: "[A-Z",
	}

	// Act
	err := edit.Validate()

	// Assert
	g.Expect(err).To(gomega.HaveOccurred())
	g.Expect(err.Error()).To(gomega.ContainSubstring("unable to compile 'search':"))
}
