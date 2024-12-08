package model

import (
	"strings"
	"testing"

	. "github.com/onsi/gomega"
)

func TestMarkerValue_WhenEmpty_DoesNotHaveValue(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	value := MakeMarkerValue("test")

	result, ok := value.Value()

	g.Expect(result).To(BeEmpty())
	g.Expect(ok).To(BeFalse())
}

func TestMarkerValue_WhenEmpty_HasNoValue(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	value := MakeMarkerValue("groupName")

	_, ok := value.Value()
	g.Expect(ok).To(BeFalse())
}

func TestMarkerValue_AfterUpdate_ValueHasExpectedResult(t *testing.T) {
	t.Parallel()

	cases := map[string]struct {
		path          string
		marker        string
		expectedValue string
		expectedOk    bool
	}{
		"Matching marker": {
			path:          "groupName",
			marker:        "+groupName=alertsmanagement.azure.com",
			expectedOk:    true,
			expectedValue: "alertsmanagement.azure.com",
		},
		"Non-matching marker": {
			path:       "groupName",
			marker:     "+versionName=v1api20230301storage",
			expectedOk: false,
		},
	}

	for n, c := range cases {
		t.Run(n, func(t *testing.T) {
			t.Parallel()
			g := NewGomegaWithT(t)

			markers := NewMarkers(c.marker)

			p := strings.Split(c.path, ":")
			value := MakeMarkerValue(p...)

			err := value.Update(markers)
			g.Expect(err).To(BeNil())

			result, ok := value.Value()
			g.Expect(ok).To(Equal(c.expectedOk))

			if c.expectedOk {
				g.Expect(result).To(Equal(c.expectedValue))
			}
		})
	}
}

func TestMarkerValue_WhenUpdatedWithSameValue_DoesNotReturnError(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	initMarkers := NewMarkers("+groupName=alertsmanagement.azure.com")

	value := MakeMarkerValue("groupName")
	err := value.Update(initMarkers)
	g.Expect(err).To(BeNil())

	testMarkers := NewMarkers("+groupName=alertsmanagement.azure.com")

	err = value.Update(testMarkers)
	g.Expect(err).To(BeNil())
}

func TestMarkerValue_WhenUpdatedWithDifferentValue_ReturnsError(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	initMarkers := NewMarkers("+groupName=alertsmanagement.azure.com")

	value := MakeMarkerValue("groupName")
	err := value.Update(initMarkers)
	g.Expect(err).To(BeNil())

	testMarkers := NewMarkers("+groupName=network.azure.com")

	err = value.Update(testMarkers)
	g.Expect(err).To(Not(BeNil()))
	g.Expect(err).To(MatchError(ContainSubstring("does not match")))
	g.Expect(err).To(MatchError(ContainSubstring("alertsmanagement.azure.com")))
	g.Expect(err).To(MatchError(ContainSubstring("network.azure.com")))
}

//nolint:funlen, revive // excessive length is acceptable for tests
func TestMarkerValueMerge_givenValue_ReturnsExpectedResult(t *testing.T) {
	t.Parallel()

	cases := map[string]struct {
		initialValue            string
		mergeValue              string
		expectedValue           string
		expectedErrorSubstrings []string
	}{
		"Merge when empty": {
			mergeValue:    "alertsmanagement.azure.com",
			expectedValue: "alertsmanagement.azure.com",
		},
		"Merge of empty": {
			initialValue:  "alertsmanagement.azure.com",
			expectedValue: "alertsmanagement.azure.com",
		},
		"Merge of same value": {
			initialValue:  "alertsmanagement.azure.com",
			mergeValue:    "alertsmanagement.azure.com",
			expectedValue: "alertsmanagement.azure.com",
		},
		"Merge of different value": {
			initialValue: "alertsmanagement.azure.com",
			mergeValue:   "network.azure.com",
			expectedErrorSubstrings: []string{
				"does not match",
				"alertsmanagement.azure.com",
				"network.azure.com",
			},
		},
	}

	for n, c := range cases {
		t.Run(n, func(t *testing.T) {
			t.Parallel()
			g := NewGomegaWithT(t)

			value := MakeMarkerValue("groupName")
			if c.initialValue != "" {
				g.Expect(value.SetValue(c.initialValue)).To(Succeed())
			}

			mergeValue := MakeMarkerValue("groupName")
			if c.mergeValue != "" {
				g.Expect(mergeValue.SetValue(c.mergeValue)).To(Succeed())
			}

			err := value.Merge(mergeValue)
			if len(c.expectedErrorSubstrings) > 0 {
				g.Expect(err).To(Not(BeNil()))

				for _, s := range c.expectedErrorSubstrings {
					g.Expect(err).To(MatchError(ContainSubstring(s)))
				}
			} else {
				g.Expect(err).To(BeNil())

				result, ok := value.Value()

				g.Expect(ok).To(BeTrue())
				g.Expect(result).To(Equal(c.expectedValue))
			}
		})
	}
}
