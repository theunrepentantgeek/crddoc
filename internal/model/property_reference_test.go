package model

import (
	"testing"

	. "github.com/onsi/gomega"
)

func TestComparePropertyReferences_GivenReferences_ReturnsExpectedResult(t *testing.T) {
	t.Parallel()

	familyname := NewPropertyReference("person", "person", "familyname")
	knownas := NewPropertyReference("person", "person", "knownas")
	startDate := NewPropertyReference("student", "student", "startDate")

	cases := map[string]struct {
		left     PropertyReference
		right    PropertyReference
		expected int
	}{
		"same host, same property": {
			left:     familyname,
			right:    familyname,
			expected: 0,
		},
		"same host, different property": {
			left:     familyname,
			right:    knownas,
			expected: -1,
		},
		"same host, different property, reversed order": {
			left:     knownas,
			right:    familyname,
			expected: 1,
		},
		"different host": {
			left:     familyname,
			right:    startDate,
			expected: -1,
		},
		"different host, reversed order": {
			left:     startDate,
			right:    familyname,
			expected: 1,
		},
	}

	for name, c := range cases {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			g := NewGomegaWithT(t)

			result := ComparePropertyReferences(c.left, c.right)
			g.Expect(result).To(Equal(c.expected))
		})
	}
}
