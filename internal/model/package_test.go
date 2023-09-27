package model

import (
	"github.com/go-logr/logr"
	"github.com/onsi/gomega"
	"github.com/theunrepentantgeek/crddoc/internal/config"

	"os"
	"path/filepath"
	"testing"
)

func TestPackage_LoadFile_LoadsExpectedContent(t *testing.T) {
	t.Parallel()
	g := gomega.NewWithT(t)

	cfg := &config.Config{}
	pkg := NewPackage(cfg, logr.Discard())

	err := pkg.LoadFile(testdataPath(t, "person_types.go"))
	if err != nil {
		t.Fatalf("Failed to load file: %v", err)
	}

	g.Expect(len(pkg.declarations)).To(gomega.Equal(4))
}

func TestPackage_Objects_ReturnsExpectedSequence(t *testing.T) {
	t.Parallel()
	g := gomega.NewWithT(t)

	cfg := &config.Config{}
	pkg := NewPackage(cfg, logr.Discard())

	err := pkg.LoadFile(testdataPath(t, "person_types.go"))
	if err != nil {
		t.Fatalf("Failed to load file: %v", err)
	}

	declarations := pkg.Declarations(OrderAlphabetical)
	g.Expect(len(declarations)).To(gomega.Equal(4))

	g.Expect(declarations[0].Name()).To(gomega.Equal("PersonReference"))
	g.Expect(declarations[1].Name()).To(gomega.Equal("PersonResource"))
	g.Expect(declarations[2].Name()).To(gomega.Equal("PersonResourceSpec"))
	g.Expect(declarations[3].Name()).To(gomega.Equal("PersonResourceStatus"))
}

func TestPackage_Object_ReturnsExpectedObjects(t *testing.T) {
	t.Parallel()
	g := gomega.NewWithT(t)

	cfg := &config.Config{}
	pkg := NewPackage(cfg, logr.Discard())

	err := pkg.LoadFile(testdataPath(t, "person_types.go"))
	if err != nil {
		t.Fatalf("Failed to load file: %v", err)
	}

	resource, ok := pkg.Declaration("PersonResource")
	g.Expect(ok).To(gomega.BeTrue())
	g.Expect(resource).NotTo(gomega.BeNil())

	spec, ok := pkg.Declaration("PersonResourceSpec")
	g.Expect(ok).To(gomega.BeTrue())
	g.Expect(spec).NotTo(gomega.BeNil())

	status, ok := pkg.Declaration("PersonResourceStatus")
	g.Expect(ok).To(gomega.BeTrue())
	g.Expect(status).NotTo(gomega.BeNil())

	ref, ok := pkg.Declaration("PersonReference")
	g.Expect(ok).To(gomega.BeTrue())
	g.Expect(ref).NotTo(gomega.BeNil())
}

// loadTestData is a helper used to load a testdata source file
func testdataPath(t *testing.T, filename string) string {
	t.Helper()

	wd, err := os.Getwd()
	if err != nil {
		t.Fatalf("Failed to get working directory: %v", err)
	}

	filepath := filepath.Join(wd, "testdata", filename)
	return filepath
}
