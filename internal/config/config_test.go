package config

import (
	"testing"

	. "github.com/onsi/gomega"
)

func TestConfig_Validate_WhenFilterInvalid_ReturnsError(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	// Arrange
	cfg := &Config{
		TypeFilters: []*Filter{
			{
				Include: "Foo",
				Exclude: "Bar",
			},
		},
	}

	// Act
	err := cfg.Validate()

	// Assert
	g.Expect(err).To(HaveOccurred())
}

func TestConfig_Validate_WhenModeIsEmpty_SetsDefaultMode(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	// Arrange
	cfg := &Config{
		Mode: "",
	}

	// Act
	err := cfg.Validate()

	// Assert
	g.Expect(err).ToNot(HaveOccurred())
	g.Expect(cfg.Mode).To(Equal("single-file"))
}

func TestConfig_Validate_WhenModeIsValidSingleFile_NormalizesMode(t *testing.T) {
	t.Parallel()

	testCases := []string{
		"single-file",
		"Single-File",
		"SINGLE-FILE",
		"singlefile",
		"SingleFile",
		"SINGLEFILE",
	}

	for _, mode := range testCases {
		t.Run(mode, func(t *testing.T) {
			g := NewGomegaWithT(t)
			
			// Arrange
			cfg := &Config{
				Mode: mode,
			}

			// Act
			err := cfg.Validate()

			// Assert
			g.Expect(err).ToNot(HaveOccurred())
			g.Expect(cfg.Mode).To(Equal("single-file"))
		})
	}
}

func TestConfig_Validate_WhenModeIsValidMultipleFile_NormalizesMode(t *testing.T) {
	t.Parallel()

	testCases := []string{
		"multiple-file",
		"Multiple-File",
		"MULTIPLE-FILE",
		"multiplefile",
		"MultipleFile",
		"MULTIPLEFILE",
		"multiple-files",
		"multiplefiles",
	}

	for _, mode := range testCases {
		t.Run(mode, func(t *testing.T) {
			g := NewGomegaWithT(t)
			
			// Arrange
			cfg := &Config{
				Mode: mode,
			}

			// Act
			err := cfg.Validate()

			// Assert
			g.Expect(err).ToNot(HaveOccurred())
			g.Expect(cfg.Mode).To(Equal("multiple-file"))
		})
	}
}

func TestConfig_Validate_WhenModeIsInvalid_ReturnsError(t *testing.T) {
	t.Parallel()

	testCases := []string{
		"invalid",
		"single",
		"multiple",
		"file",
		"all-files",
	}

	for _, mode := range testCases {
		t.Run(mode, func(t *testing.T) {
			g := NewGomegaWithT(t)
			
			// Arrange
			cfg := &Config{
				Mode: mode,
			}

			// Act
			err := cfg.Validate()

			// Assert
			g.Expect(err).To(HaveOccurred())
			g.Expect(err.Error()).To(ContainSubstring("invalid mode"))
			g.Expect(err.Error()).To(ContainSubstring("must be either 'single-file' or 'multiple-file'"))
		})
	}
}
