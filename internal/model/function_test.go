package model_test

import (
	"testing"

	. "github.com/onsi/gomega"

	"github.com/go-logr/logr"

	"github.com/theunrepentantgeek/crddoc/internal/config"
	"github.com/theunrepentantgeek/crddoc/internal/model"
	"github.com/theunrepentantgeek/crddoc/internal/packageloader"
)

func TestObject_Functions_ReturnsExpectedContent(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	cfg := &config.Config{}
	loader := packageloader.New(cfg, logr.Discard())

	pkg, err := loader.LoadFile(testdataPath(t, "person_with_methods.go"))
	g.Expect(err).To(Succeed())
	g.Expect(pkg).NotTo(BeNil())

	dec, ok := pkg.Declaration("PersonWithMethods")
	g.Expect(ok).To(BeTrue())

	obj, ok := dec.(*model.Object)
	g.Expect(ok).To(BeTrue())

	// Verify we have the expected functions
	functions := obj.Functions()
	g.Expect(functions).To(HaveLen(5))

	// Check function names
	functionNames := make([]string, len(functions))
	for i, fn := range functions {
		functionNames[i] = fn.Name
	}
	g.Expect(functionNames).To(ConsistOf("GetName", "SetName", "IsAdult", "UpdateAge", "Compare"))
}

func TestObject_Function_ReturnsExpectedFunction(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	cases := map[string]struct {
		functionName   string
		expectedExists bool
		checkDetails   func(*testing.T, *model.Function)
	}{
		"GetName exists with value receiver": {
			functionName:   "GetName",
			expectedExists: true,
			checkDetails: func(t *testing.T, fn *model.Function) {
				t.Helper()
				g := NewGomegaWithT(t)
				g.Expect(fn.Name).To(Equal("GetName"))
				g.Expect(fn.Receiver.Name()).To(Equal("PersonWithMethods"))
				g.Expect(fn.IsPointerReceiver).To(BeFalse())
				g.Expect(fn.Parameters).To(BeEmpty())
				g.Expect(fn.Results).To(HaveLen(1))
				g.Expect(fn.Results[0].Type.Name()).To(Equal("string"))
			},
		},
		"SetName exists with pointer receiver": {
			functionName:   "SetName",
			expectedExists: true,
			checkDetails: func(t *testing.T, fn *model.Function) {
				t.Helper()
				g := NewGomegaWithT(t)
				g.Expect(fn.Name).To(Equal("SetName"))
				g.Expect(fn.Receiver.Name()).To(Equal("PersonWithMethods"))
				g.Expect(fn.IsPointerReceiver).To(BeTrue())
				g.Expect(fn.Parameters).To(HaveLen(1))
				g.Expect(fn.Parameters[0].Name).To(Equal("name"))
				g.Expect(fn.Parameters[0].Type.Name()).To(Equal("string"))
				g.Expect(fn.Results).To(BeEmpty())
			},
		},
		"IsAdult exists": {
			functionName:   "IsAdult",
			expectedExists: true,
			checkDetails: func(t *testing.T, fn *model.Function) {
				t.Helper()
				g := NewGomegaWithT(t)
				g.Expect(fn.Name).To(Equal("IsAdult"))
				g.Expect(fn.Results).To(HaveLen(1))
				g.Expect(fn.Results[0].Type.Name()).To(Equal("bool"))
			},
		},
		"UpdateAge exists with pointer receiver and multiple operations": {
			functionName:   "UpdateAge",
			expectedExists: true,
			checkDetails: func(t *testing.T, fn *model.Function) {
				t.Helper()
				g := NewGomegaWithT(t)
				g.Expect(fn.Name).To(Equal("UpdateAge"))
				g.Expect(fn.Receiver.Name()).To(Equal("PersonWithMethods"))
				g.Expect(fn.IsPointerReceiver).To(BeTrue())
				g.Expect(fn.Parameters).To(HaveLen(1))
				g.Expect(fn.Parameters[0].Name).To(Equal("newAge"))
				g.Expect(fn.Parameters[0].Type.Name()).To(Equal("int"))
				g.Expect(fn.Results).To(HaveLen(1))
				g.Expect(fn.Results[0].Type.Name()).To(Equal("int"))
			},
		},
		"Compare exists with named return values": {
			functionName:   "Compare",
			expectedExists: true,
			checkDetails: func(t *testing.T, fn *model.Function) {
				t.Helper()
				g := NewGomegaWithT(t)
				g.Expect(fn.Name).To(Equal("Compare"))
				g.Expect(fn.Parameters).To(HaveLen(1))
				g.Expect(fn.Parameters[0].Name).To(Equal("other"))
				g.Expect(fn.Parameters[0].Type.Name()).To(Equal("PersonWithMethods"))
				g.Expect(fn.Results).To(HaveLen(2))
				g.Expect(fn.Results[0].Name).To(Equal("equal"))
				g.Expect(fn.Results[0].Type.Name()).To(Equal("bool"))
				g.Expect(fn.Results[1].Name).To(Equal("ageDiff"))
				g.Expect(fn.Results[1].Type.Name()).To(Equal("int"))
			},
		},
		"NonExistentMethod does not exist": {
			functionName:   "NonExistentMethod",
			expectedExists: false,
		},
	}

	cfg := &config.Config{}
	loader := packageloader.New(cfg, logr.Discard())

	pkg, err := loader.LoadFile(testdataPath(t, "person_with_methods.go"))
	g.Expect(err).To(Succeed())
	g.Expect(pkg).NotTo(BeNil())

	dec, ok := pkg.Declaration("PersonWithMethods")
	g.Expect(ok).To(BeTrue())

	obj, ok := dec.(*model.Object)
	g.Expect(ok).To(BeTrue())

	for n, c := range cases {
		t.Run(n, func(t *testing.T) {
			t.Parallel()

			g := NewGomegaWithT(t)

			fn, ok := obj.Function(c.functionName)
			g.Expect(ok).To(Equal(c.expectedExists))

			if c.expectedExists {
				g.Expect(fn).NotTo(BeNil())
				g.Expect(fn.Name).To(Equal(c.functionName))
				g.Expect(fn.DeclaredOn()).To(Equal(obj))

				if c.checkDetails != nil {
					c.checkDetails(t, fn)
				}
			}
		})
	}
}

func TestFunction_Description_ParsesComments(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	cfg := &config.Config{}
	loader := packageloader.New(cfg, logr.Discard())

	pkg, err := loader.LoadFile(testdataPath(t, "person_with_methods.go"))
	g.Expect(err).To(Succeed())
	g.Expect(pkg).NotTo(BeNil())

	dec, ok := pkg.Declaration("PersonWithMethods")
	g.Expect(ok).To(BeTrue())

	obj, ok := dec.(*model.Object)
	g.Expect(ok).To(BeTrue())

	// Check that GetName has a description
	fn, ok := obj.Function("GetName")
	g.Expect(ok).To(BeTrue())
	g.Expect(fn.Description()).NotTo(BeEmpty())
	g.Expect(fn.Description()[0]).To(ContainSubstring("returns the name"))
}

func TestObject_Functions_FromMultipleFiles(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	cfg := &config.Config{}
	loader := packageloader.New(cfg, logr.Discard())

	// Load the entire testdata directory to get methods from both files
	pkg, err := loader.LoadDirectory("testdata")
	g.Expect(err).To(Succeed())
	g.Expect(pkg).NotTo(BeNil())

	dec, ok := pkg.Declaration("PersonWithMethods")
	g.Expect(ok).To(BeTrue())

	obj, ok := dec.(*model.Object)
	g.Expect(ok).To(BeTrue())

	// Verify we have functions from both files
	functions := obj.Functions()
	// 5 from person_with_methods.go + 2 from person_with_methods_extra.go
	g.Expect(functions).To(HaveLen(7))

	// Check that we have the extra methods
	functionNames := make([]string, len(functions))
	for i, fn := range functions {
		functionNames[i] = fn.Name
	}
	g.Expect(functionNames).To(ContainElement("GetFullInfo"))
	g.Expect(functionNames).To(ContainElement("IncrementAge"))
}
