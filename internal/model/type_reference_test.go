package model

import (
	"testing"

	. "github.com/onsi/gomega"

	"github.com/dave/dst"
)

var (
	clusterID = dst.NewIdent("Cluster")
	apiID     = dst.NewIdent("api")
	fullID    = dst.SelectorExpr{
		X:   apiID,
		Sel: clusterID,
	}
	optionalClusterID = &dst.StarExpr{
		X: clusterID,
	}
	optionalFullID = dst.StarExpr{
		X: &fullID,
	}
)

func Test_TypeReference_nameOf_GivenExpression_returnsExpectedResult(t *testing.T) {
	t.Parallel()

	cases := map[string]struct {
		expr     dst.Expr
		expected string
	}{
		"Simple Identifier": {
			expr:     clusterID,
			expected: "Cluster",
		},
		"Optional identifier": {
			expr:     optionalClusterID,
			expected: "Cluster",
		},
		"Full identifier": {
			expr:     &fullID,
			expected: "api.Cluster",
		},
		"Optional full identifier": {
			expr:     &optionalFullID,
			expected: "api.Cluster",
		},
	}

	for n, c := range cases {
		t.Run(n, func(t *testing.T) {
			t.Parallel()
			g := NewGomegaWithT(t)

			result := displayOf(c.expr)
			g.Expect(result).To(Equal(c.expected))
		})
	}
}

func Test_TypeReference_pkgOf_GivenExpression_returnsExpectedResult(t *testing.T) {
	t.Parallel()

	cases := map[string]struct {
		expr     dst.Expr
		expected string
	}{
		"Simple Identifier": {
			expr:     clusterID,
			expected: "",
		},
		"Optional identifier": {
			expr:     optionalClusterID,
			expected: "",
		},
		"Full identifier": {
			expr:     &fullID,
			expected: "api",
		},
		"Optional full identifier": {
			expr:     &optionalFullID,
			expected: "api",
		},
	}

	for n, c := range cases {
		t.Run(n, func(t *testing.T) {
			t.Parallel()
			g := NewGomegaWithT(t)

			result := pkgOf(c.expr)
			g.Expect(result).To(Equal(c.expected))
		})
	}
}
