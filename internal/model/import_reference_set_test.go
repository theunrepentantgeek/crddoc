// filepath: /workspaces/crddoc/internal/model/import_reference_set_test.go
package model

import (
	"testing"

	. "github.com/onsi/gomega"
)

func Test_ImportReferenceSet_Add_GivenImportReference_AddsToSet(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	// Arrange
	set := make(ImportReferenceSet)
	ref1 := ImportReference{
		ImportPath: "github.com/example/pkg1",
		Alias:      "pkg1",
	}
	ref2 := ImportReference{
		ImportPath: "github.com/example/pkg2",
		Alias:      "pkg2",
	}

	// Act
	set.Add(ref1)
	set.Add(ref2)

	// Assert
	g.Expect(set).To(HaveLen(2))
	g.Expect(set[ref1.Alias]).To(Equal(ref1))
	g.Expect(set[ref2.Alias]).To(Equal(ref2))
}

func Test_ImportReferenceSet_LookupImportPath(t *testing.T) {
	t.Parallel()

	pkg := ImportReference{
		ImportPath: "github.com/example/pkg",
		Alias:      "pkg",
	}

	// Setup test cases
	testCases := map[string]struct {
		refSet       ImportReferenceSet
		ref          TypeReference
		expectedPath string
	}{
		"Empty package in TypeReference returns false": {
			refSet: NewImportReferenceSet(pkg),
			ref: TypeReference{
				pkg:  "",
				name: "SomeType",
			},
			expectedPath: "",
		},
		"Unknown package in TypeReference returns false": {
			refSet: NewImportReferenceSet(pkg),
			ref: TypeReference{
				pkg:  "otherpkg",
				name: "SomeType",
			},
			expectedPath: "",
		},
		"Known package in TypeReference returns import path": {
			refSet: NewImportReferenceSet(pkg),
			ref: TypeReference{
				pkg:  "pkg",
				name: "SomeType",
			},
			expectedPath: "github.com/example/pkg",
		},
	}

	// Run the tests
	for name, c := range testCases {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			g := NewGomegaWithT(t)

			// Act
			importPath, found := c.refSet.LookupImportPath(c.ref)

			// Assert
			if c.expectedPath == "" {
				g.Expect(found).To(BeFalse())
			} else {
				g.Expect(found).To(BeTrue())
				g.Expect(importPath).To(Equal(c.expectedPath))
			}
		})
	}
}
