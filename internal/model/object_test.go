package model

import (
	"testing"

	"github.com/go-logr/logr"
	"github.com/theunrepentantgeek/crddoc/internal/config"

	. "github.com/onsi/gomega"
)

func TestObject_Property_ReturnsExpectedContent(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	cfg := &config.Config{}
	pkg := NewPackage(cfg, logr.Discard())
	err := pkg.LoadFile(testdataPath(t, "person_types.go"))
	if err != nil {
		t.Fatalf("Failed to load file: %v", err)
	}

	dec, ok := pkg.Declaration("PersonResourceSpec")
	g.Expect(ok).To(BeTrue())

	obj, ok := dec.(*Object)
	g.Expect(ok).To(BeTrue())

	fullName, ok := obj.Property("FullName")
	g.Expect(ok).To(BeTrue())
	g.Expect(fullName).NotTo(BeNil())

	children, ok := obj.Property("Children")
	g.Expect(ok).To(BeTrue())
	g.Expect(children).NotTo(BeNil())
}
