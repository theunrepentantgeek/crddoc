package model

import (
	"testing"

	. "github.com/onsi/gomega"
)

func TestPropertyMarkersParse_GivenMarkers_HasExpectedProperties(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	cases := map[string]struct {
		lines            []string
		expectedOptional bool
		expectedRequired bool
	}{
		"Empty": {
			lines:            []string{},
			expectedOptional: false,
			expectedRequired: false,
		},
		"Properties optional specified": {
			lines: []string{
				"// +kubebuilder:validation:Optional",
			},
			expectedOptional: true,
		},
		"Properties required specified": {
			lines: []string{
				"// +kubebuilder:validation:Required",
			},
			expectedRequired: true,
		},
	}

	for n, c := range cases {
		t.Run(n, func(t *testing.T) {
			t.Parallel()

			m := NewPropertyMarkers()
			markers := NewMarkers(c.lines...)
			err := m.Parse(markers)
			g.Expect(err).NotTo(HaveOccurred())

			g.Expect(m.Optional()).To(Equal(c.expectedOptional))
			g.Expect(m.Required()).To(Equal(c.expectedRequired))
		})
	}
}
