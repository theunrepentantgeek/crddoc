package model

import (
	"strings"
	"testing"

	. "github.com/onsi/gomega"
)

func TestMarkerSwitch_WhenEmpty_HasNotBeenSeen(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	mds := MakeMarkerSwitch("test")
	g.Expect(mds.Seen()).To(BeFalse())
}

func TestMarkerSwitch_AfterUpdate_SeenHasExpectedResult(t *testing.T) {
	t.Parallel()

	cases := map[string]struct {
		path         string
		marker       string
		expectedSeen bool
	}{
		"No marker": {
			path:         "groupName",
			marker:       "",
			expectedSeen: false,
		},
		"Matching marker": {
			path:         "kubebuilder:validation:Optional",
			marker:       "+kubebuilder:validation:Optional",
			expectedSeen: true,
		},
		"Matching prefix of marker": {
			path:         "kubebuilder:validation",
			marker:       "+kubebuilder:validation:Optional",
			expectedSeen: true,
		},
		"Non-matching marker": {
			path:         "kubebuilder:validation",
			marker:       "+groupName=cache.azure.com",
			expectedSeen: false,
		},
	}

	for n, c := range cases {
		t.Run(n, func(t *testing.T) {
			t.Parallel()
			g := NewGomegaWithT(t)

			markers := NewMarkers()
			if c.marker != "" {
				markers.Add(c.marker)
			}

			p := strings.Split(c.path, ":")
			value := MakeMarkerSwitch(p...)

			value.Update(markers)
			g.Expect(value.Seen()).To(Equal(c.expectedSeen))
		})
	}
}
