package model_test

import (
	"os"
	"path/filepath"
	"testing"

	. "github.com/onsi/gomega"

	"github.com/go-logr/logr"

	"github.com/theunrepentantgeek/crddoc/internal/config"
	"github.com/theunrepentantgeek/crddoc/internal/model"
	"github.com/theunrepentantgeek/crddoc/internal/packageloader"
)

func TestPackage_LoadFile_LoadsExpectedContent(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	cfg := &config.Config{}

	loader := packageloader.New(cfg, logr.Discard())

	pkg, err := loader.LoadFile(testdataPath(t, "person_types.go"))
	if err != nil {
		t.Fatalf("Failed to load file: %v", err)
	}

	g.Expect(len(pkg.Declarations(model.OrderAlphabetical))).To(Equal(4))
}

func TestPackage_Objects_ReturnsExpectedSequence(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	cfg := &config.Config{}
	loader := packageloader.New(cfg, logr.Discard())

	pkg, err := loader.LoadFile(testdataPath(t, "person_types.go"))
	if err != nil {
		t.Fatalf("Failed to load file: %v", err)
	}

	declarations := pkg.Declarations(model.OrderAlphabetical)
	g.Expect(len(declarations)).To(Equal(4))

	g.Expect(declarations[0].Name()).To(Equal("PersonReference"))
	g.Expect(declarations[1].Name()).To(Equal("PersonResource"))
	g.Expect(declarations[2].Name()).To(Equal("PersonResourceSpec"))
	g.Expect(declarations[3].Name()).To(Equal("PersonResourceStatus"))
}

func TestPackage_Declaration_ReturnsExpectedObjects(t *testing.T) {
	t.Parallel()

	cases := map[string]struct {
		expected bool
	}{
		"PersonResource": {
			expected: true,
		},
		"PersonResourceSpec": {
			expected: true,
		},
		"PersonResourceStatus": {
			expected: true,
		},
		"PersonReference": {
			expected: true,
		},
	}

	cfg := &config.Config{}
	loader := packageloader.New(cfg, logr.Discard())

	pkg, err := loader.LoadFile(testdataPath(t, "person_types.go"))
	if err != nil {
		t.Fatalf("Failed to load file: %v", err)
	}

	for name, c := range cases {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			g := NewGomegaWithT(t)

			decl, ok := pkg.Declaration(name)
			if c.expected {
				g.Expect(ok).To(BeTrue())
				g.Expect(decl).NotTo(BeNil())
			} else {
				g.Expect(ok).To(BeFalse())
				g.Expect(decl).To(BeNil())
			}
		})
	}
}

// loadTestData is a helper used to load a testdata source file.
func testdataPath(
	t *testing.T,
	filename string, //nolint:unparam // other file names will be used as testing improves.
) string {
	t.Helper()

	wd, err := os.Getwd()
	if err != nil {
		t.Fatalf("Failed to get working directory: %v", err)
	}

	result := filepath.Join(wd, "testdata", filename)

	return result
}

func TestPackage_Rank_ReturnesExpectedResults(t *testing.T) {
	t.Parallel()

	cases := map[string]struct {
		expectedRank int
	}{
		"PersonResource": {
			expectedRank: 0,
		},
		"PersonResourceSpec": {
			expectedRank: 1,
		},
		"PersonResourceStatus": {
			expectedRank: 1,
		},
		"PersonReference": {
			expectedRank: 2,
		},
	}

	cfg := &config.Config{}
	loader := packageloader.New(cfg, logr.Discard())

	pkg, err := loader.LoadFile(testdataPath(t, "person_types.go"))
	if err != nil {
		t.Fatalf("Failed to load file: %v", err)
	}

	for name, c := range cases {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			g := NewGomegaWithT(t)

			rank := pkg.Rank(name)
			g.Expect(rank).To(Equal(c.expectedRank))
		})
	}
}
