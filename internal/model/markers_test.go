package model

import (
	"strings"
	"testing"

	. "github.com/onsi/gomega"
)

func TestMarkers_Lookup_ReturnsExpectedValue(t *testing.T) {
	t.Parallel()

	cases := []struct {
		name     string
		lookup   string
		expected string
	}{
		{
			"Required",
			"kubebuilder:validation",
			"Max=100",
		},
	}

	markers := NewMarkers()
	markers.Add("+kubebuilder:validation:Max=100")

	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()
			g := NewGomegaWithT(t)

			path := strings.Split(c.lookup, ":")
			actual, ok := markers.Lookup(path...)

			g.Expect(ok).To(Equal(c.expected != ""))
			g.Expect(actual).To(Equal(c.expected))
		})
	}
}

func TestMarkers_Exists_ReturnsExpectedValue(t *testing.T) {
	t.Parallel()

	cases := []struct {
		name     string
		lookup   string
		expected bool
	}{
		{
			"Required",
			"kubebuilder:validation:Required",
			true,
		},
		{
			"Optional",
			"kubebuilder:validation:Optional",
			true,
		},
		{
			"Flagged",
			"kubebuilder:validation:Flagged",
			false,
		},
	}

	markers := NewMarkers()
	markers.Add("+kubebuilder:validation:Required")
	markers.Add("+kubebuilder:validation:Optional")

	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()
			g := NewGomegaWithT(t)

			path := strings.Split(c.lookup, ":")
			found := markers.Exists(path...)

			g.Expect(found).To(Equal(c.expected))
		})
	}
}
