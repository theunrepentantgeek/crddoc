package functions

import (
	"testing"

	. "github.com/onsi/gomega"

	"github.com/go-logr/logr"

	"github.com/theunrepentantgeek/crddoc/internal/config"
)

func TestIncludeFunctions_WhenFalse_ReturnsFalse(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	// Arrange
	cfg := &config.Config{
		IncludeFunctions: false,
	}
	fns := New(logr.Discard())
	err := fns.SetConfig(cfg)
	g.Expect(err).To(Succeed())

	// Act
	result := fns.includeFunctions()

	// Assert
	g.Expect(result).To(BeFalse())
}

func TestIncludeFunctions_WhenTrue_ReturnsTrue(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	// Arrange
	cfg := &config.Config{
		IncludeFunctions: true,
	}
	fns := New(logr.Discard())
	err := fns.SetConfig(cfg)
	g.Expect(err).To(Succeed())

	// Act
	result := fns.includeFunctions()

	// Assert
	g.Expect(result).To(BeTrue())
}
