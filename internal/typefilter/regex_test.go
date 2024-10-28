package typefilter

import (
	"testing"

	. "github.com/onsi/gomega"
)

func TestCreateGlobber_GivenGlob_ReturnsExpectedRegex(t *testing.T) {
	t.Parallel()

	cases := map[string]struct {
		glob     string
		expected string
	}{
		"simple word": {
			glob:     "alpha",
			expected: "(?i)alpha",
		},
		"text files": {
			glob:     "*.txt",
			expected: `(?i).*\.txt`,
		},
		"all files": {
			glob:     "*.*",
			expected: `(?i).*\..*`,
		},
	}

	for name, c := range cases {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			g := NewGomegaWithT(t)

			globber := createGlobber(c.glob)
			g.Expect(globber.String()).To(Equal(c.expected))
		})
	}
}
