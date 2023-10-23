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

	cfg := &config.Config{}
	loader := packageloader.New(cfg, logr.Discard())
	pkg, err := loader.LoadFile(testdataPath(t, "person_types.go"))
	if err != nil {
		t.Fatalf("Failed to load file: %v", err)
	}

	dec, ok := pkg.Declaration("PersonResourceSpec")
	g.Expect(ok).To(BeTrue())

	obj, ok := dec.(*model.Object)
	g.Expect(ok).To(BeTrue())

	fullName, ok := obj.Property("fullName")
	g.Expect(ok).To(BeTrue())
	g.Expect(fullName).NotTo(BeNil())

	children, ok := obj.Property("children")
	g.Expect(ok).To(BeTrue())
	g.Expect(children).NotTo(BeNil())
}
