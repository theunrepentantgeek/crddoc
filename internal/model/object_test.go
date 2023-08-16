package model

import (
	"testing"

	"github.com/go-logr/logr"
	. "github.com/onsi/gomega"
)

func TestObject_Property_ReturnsExpectedContent(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	pkg := NewPackage(logr.Discard())
	err := pkg.LoadFile(testdataPath(t, "person_types.go"))
	if err != nil {
		t.Fatalf("Failed to load file: %v", err)
	}

	obj, ok := pkg.Object("PersonResource")
	g.Expect(ok).To(BeTrue())

	spec, ok := obj.Property("Spec")
	g.Expect(ok).To(BeTrue())
	g.Expect(spec).NotTo(BeNil())

	status, ok := obj.Property("Status")
	g.Expect(ok).To(BeTrue())
	g.Expect(status).NotTo(BeNil())
}
