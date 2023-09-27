package config

import (
	"testing"

	"github.com/onsi/gomega"
)

func TestEditor_Replace_ReturnsExpectedResult(t *testing.T) {
	t.Parallel()

	redactVowels := &Editor{
		Search:      "[aeiou]",
		Replacement: "x",
	}

	redactVowelsInWordsStartingWithM := &Editor{
		Context:     "(?i)m\\w*",
		Search:      "[aeiou]",
		Replacement: "x",
	}

	cases := []struct {
		name     string
		editor   *Editor
		input    string
		expected string
	}{
		{
			name:     "no match",
			editor:   redactVowels,
			input:    "xyz",
			expected: "xyz",
		},
		{
			name:     "single match",
			editor:   redactVowels,
			input:    "which",
			expected: "whxch",
		},
		{
			name:     "multiple matches",
			editor:   redactVowels,
			input:    "which one",
			expected: "whxch xnx",
		},
		{
			name:     "Only words starting with M",
			editor:   redactVowelsInWordsStartingWithM,
			input:    "five marvelous monkeys in the moonlight",
			expected: "five mxrvxlxxs mxnkxys in the mxxnlxght",
		},
	}

	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()

			g := gomega.NewWithT(t)
			g.Expect(c.editor.Validate()).To(gomega.Succeed())

			actual := c.editor.Replace(c.input)
			g.Expect(actual).To(gomega.Equal(c.expected))
		})
	}
}

func TestEditor_Replace_WhenInvalid_Panics(t *testing.T) {
	t.Parallel()

	g := gomega.NewWithT(t)

	// Arrange
	edit := &Editor{
		Search: "[A-Z",
	}

	// Act & Assert
	g.Expect(
		func() {
			edit.Replace("foo")
		}).To(gomega.Panic())
}

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
