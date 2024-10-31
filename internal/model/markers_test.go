package model_test

import (
	"strings"
	"testing"

	"github.com/onsi/gomega"

	"github.com/theunrepentantgeek/crddoc/internal/model"
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
			"kubebuilder:validation:Max",
			"100",
		},
	}

	markers := model.NewMarkers()
	markers.Add("+kubebuilder:validation:Max=100")

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()
			g := gomega.NewWithT(t)

			path := strings.Split(c.lookup, ":")
			actual, ok := markers.Lookup(path...)

			g.Expect(ok).To(gomega.Equal(c.expected != ""))
			g.Expect(actual).To(gomega.Equal(c.expected))
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

	markers := model.NewMarkers()
	markers.Add("+kubebuilder:validation:Required")
	markers.Add("+kubebuilder:validation:Optional")

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()
			g := gomega.NewWithT(t)

			path := strings.Split(c.lookup, ":")
			found := markers.Exists(path...)

			g.Expect(found).To(gomega.Equal(c.expected))
		})
	}
}
