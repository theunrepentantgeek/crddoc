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
			t.Parallel()
			g := NewGomegaWithT(t)

			// Arrange
			cfg := &Config{
				FileMode: mode,
			}

			// Act
			err := cfg.Validate()

			// Assert
			g.Expect(err).To(MatchError(ContainSubstring("invalid mode")))
			g.Expect(err).To(MatchError(
				ContainSubstring("must be either 'single-file' or 'multiple-file'")))
		})
	}
}

func TestConfig_HasFileMode_ReturnsExpectedResults(t *testing.T) {
	t.Parallel()

	cases := map[string]struct {
		setFileMode  string
		testFileMode string
		expected     bool
	}{
		"single-file": {
			setFileMode:  FileModeSingleFile,
			testFileMode: FileModeSingleFile,
			expected:     true,
		},
		"multiple-file": {
			setFileMode:  FileModeMultipleFile,
			testFileMode: FileModeMultipleFile,
			expected:     true,
		},
		"single-file lowercase": {
			setFileMode:  FileModeSingleFile,
			testFileMode: "single-file",
			expected:     true,
		},
		"multiple-file lowercase": {
			setFileMode:  FileModeMultipleFile,
			testFileMode: "multiple-file",
			expected:     true,
		},
		"single-file-not-multiple-file": {
			setFileMode:  FileModeSingleFile,
			testFileMode: FileModeMultipleFile,
			expected:     false,
		},
		"multiple-file-not-single-file": {
			setFileMode:  FileModeMultipleFile,
			testFileMode: FileModeSingleFile,
			expected:     false,
		},
	}

	for n, c := range cases {
		t.Run(n, func(t *testing.T) {
			t.Parallel()

			g := NewGomegaWithT(t)

			cfg := Standard()
			cfg.SetFileMode(&c.setFileMode)

			g.Expect(cfg.HasFileMode(c.testFileMode)).To(Equal(c.expected))
		})
	}
}
