package config

import (
	"testing"

	"github.com/onsi/gomega"
)

func TestConfig_Validate_WhenFilterInvalid_ReturnsError(t *testing.T) {
	t.Parallel()
	g := gomega.NewWithT(t)

	// Arrange
	c := &Config{
		TypeFilters: []*Filter{
			{
				Include: "Foo",
				Exclude: "Bar",
			},
		},
	}

	// Act
	err := c.Validate()

	// Assert
	g.Expect(err).To(gomega.HaveOccurred())
}
