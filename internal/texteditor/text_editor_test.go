package texteditor

import (
	"testing"

	. "github.com/onsi/gomega"

	"github.com/theunrepentantgeek/crddoc/internal/config"
)

func TestEditor_Replace_ReturnsExpectedResult(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	redactVowels, err := NewEditor(
		config.Editor{
			Search:      "[aeiou]",
			Replacement: "x",
		})
	g.Expect(err).To(Succeed())

	redactVowelsInWordsStartingWithM, err := NewEditor(
		config.Editor{
			Context:     "(?i)m\\w*",
			Search:      "[aeiou]",
			Replacement: "x",
		})
	g.Expect(err).To(Succeed())

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
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()
			g := NewWithT(t)

			actual := c.editor.Replace(c.input)
			g.Expect(actual).To(Equal(c.expected))
		})
	}
}

func TestEditor_Replace_WhenInvalid_Panics(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	// Arrange
	edit, err := NewEditor(
		config.Editor{
			Search: "[A-Z",
		})
	g.Expect(err).NotTo(Succeed())

	// Act & Assert
	g.Expect(
		func() {
			edit.Replace("foo")
		}).To(Panic())
}
