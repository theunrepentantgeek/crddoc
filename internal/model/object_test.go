package model_test

import (
	"testing"

	. "github.com/onsi/gomega"

	"github.com/go-logr/logr"

	"github.com/theunrepentantgeek/crddoc/internal/config"
	"github.com/theunrepentantgeek/crddoc/internal/model"
	"github.com/theunrepentantgeek/crddoc/internal/packageloader"
)

func TestObject_Property_ReturnsExpectedContent(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	cases := map[string]struct {
		propertyName   string
		expectedExists bool
	}{
		"fullName exists": {
			propertyName:   "fullName",
			expectedExists: true,
		},
		"children exists": {
			propertyName:   "children",
			expectedExists: true,
		},
		"profession does not exist": {
			propertyName:   "profession",
			expectedExists: false,
		},
	}

	cfg := &config.Config{}
	loader := packageloader.New(cfg, logr.Discard())

	pkg, err := loader.LoadFile(testdataPath(t, "person_types.go"))
	g.Expect(err).To(Succeed())
	g.Expect(pkg).NotTo(BeNil())

	dec, ok := pkg.Declaration("PersonResourceSpec")
	g.Expect(ok).To(BeTrue())

	obj, ok := dec.(*model.Object)
	g.Expect(ok).To(BeTrue())

	for n, c := range cases {
		t.Run(n, func(t *testing.T) {
			t.Parallel()

			g := NewGomegaWithT(t)

			prop, ok := obj.Property(c.propertyName)
			g.Expect(ok).To(Equal(c.expectedExists))

			if c.expectedExists {
				g.Expect(prop).NotTo(BeNil())
				g.Expect(prop.Name).To(Equal(c.propertyName))
				g.Expect(prop.DeclaredOn).To(Equal(obj))
			}
		})
	}
}
