package functions

import (
	"testing"

	. "github.com/onsi/gomega"
)

func TestInlineLinks_GivenLines_ReturnsExpectedLines(t *testing.T) {
	t.Parallel()

	cases := []struct {
		name     string
		lines    []string
		expected []string
	}{
		{
			name: "Original Bug",
			lines: []string{
				"KubeletConfig defines the supported subset of kubelet configurations",
				"for nodes in pools. See also [AKS doc], [K8s doc].",
				"[AKS doc]: https://learn.microsoft.com/azure/aks/custom-node-configuration",
				//nolint:revive // disable long-line-length because test cases are long
				"[K8s doc]: https://kubernetes.io/docs/reference/config-api/kubelet-config.v1beta1/",
			},
			expected: []string{
				"KubeletConfig defines the supported subset of kubelet configurations",
				//nolint:revive // disable long-line-length because test cases are long
				"for nodes in pools. See also [AKS doc](https://learn.microsoft.com/azure/aks/custom-node-configuration), [K8s doc](https://kubernetes.io/docs/reference/config-api/kubelet-config.v1beta1/).",
			},
		},
		{
			name: "Preceding Reference",
			lines: []string{
				"[foo]: /url",
				"[foo]",
			},
			expected: []string{
				"[foo](/url)",
			},
		},
		{
			name: "Succeeding Reference",
			lines: []string{
				"[foo]",
				"[foo]: /url",
			},
			expected: []string{
				"[foo](/url)",
			},
		},
		{
			name: "First definition wins",
			lines: []string{
				"[foo]",
				"[foo]: /first",
				"[foo]: /second",
			},
			expected: []string{
				"[foo](/first)",
			},
		},
		{
			name: "Link names are not case sensitive",
			lines: []string{
				"[foo]",
				"[FOO]: /url",
			},
			expected: []string{
				"[foo](/url)",
			},
		},
	}

	for _, c := range cases {
		t.Run(
			c.name,
			func(t *testing.T) {
				t.Parallel()
				g := NewGomegaWithT(t)

				f := &Functions{}
				actual := f.inlineLinks(c.lines)
				g.Expect(actual).To(Equal(c.expected))
			})
	}
}
