package config

import (
	"os"
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
			g.Expect(err).To(MatchError(ContainSubstring("invalid file mode")))
			g.Expect(err).To(MatchError(ContainSubstring(mode)))
			g.Expect(err).To(MatchError(ContainSubstring(FileModeSingleFile)))
			g.Expect(err).To(MatchError(ContainSubstring(FileModeMultipleFile)))
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

func TestConfig_SetIncludeFunctions_WithNil_DoesNotChangeValue(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	// Arrange
	cfg := Standard()
	cfg.IncludeFunctions = true

	// Act
	cfg.SetIncludeFunctions(nil)

	// Assert
	g.Expect(cfg.IncludeFunctions).To(BeTrue())
}

func TestConfig_SetIncludeFunctions_WithTrue_SetsTrue(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	// Arrange
	cfg := Standard()
	value := true

	// Act
	cfg.SetIncludeFunctions(&value)

	// Assert
	g.Expect(cfg.IncludeFunctions).To(BeTrue())
}

func TestConfig_SetIncludeFunctions_WithFalse_SetsFalse(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	// Arrange
	cfg := Standard()
	cfg.IncludeFunctions = true
	value := false

	// Act
	cfg.SetIncludeFunctions(&value)

	// Assert
	g.Expect(cfg.IncludeFunctions).To(BeFalse())
}

func TestConfig_Load_WithIncludeFunctions_LoadsCorrectly(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	// Create a temporary config file
	tmpDir := t.TempDir()
	configPath := tmpDir + "/test-config.yaml"

	configContent := `includeFunctions: true
prettyPrint: true
`
	err := os.WriteFile(configPath, []byte(configContent), 0o600)
	g.Expect(err).To(Succeed())

	// Arrange
	cfg := Standard()

	// Act
	err = cfg.Load(configPath)

	// Assert
	g.Expect(err).To(Succeed())
	g.Expect(cfg.IncludeFunctions).To(BeTrue())
}
