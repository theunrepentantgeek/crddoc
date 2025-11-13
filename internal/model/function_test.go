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
	g.Expect(functions).To(HaveLen(6))

	// Check function names
	functionNames := make([]string, len(functions))
	for i, fn := range functions {
		functionNames[i] = fn.Name
	}

	g.Expect(functionNames).
		To(ConsistOf("GetName", "SetName", "IsAdult", "UpdateAge", "Compare", "Lookup"))
}

//nolint:revive,funlen,maintidx // comprehensive test coverage requires detailed test cases
func TestObject_Function_ReturnsExpectedFunction(t *testing.T) {
	t.Parallel()

	cases := map[string]struct {
		functionName   string
		expectedExists bool
		checkDetails   func(*testing.T, *GomegaWithT, *model.Function)
	}{
		"GetName exists with value receiver": {
			functionName:   "GetName",
			expectedExists: true,
			checkDetails: func(t *testing.T, g *GomegaWithT, fn *model.Function) {
				t.Helper()
				g.Expect(fn.Name).To(Equal("GetName"))
				g.Expect(fn.Receiver.Name()).To(Equal("PersonWithMethods"))
				g.Expect(fn.IsPointerReceiver).To(BeFalse())
				g.Expect(fn.Parameters).To(BeEmpty())
				g.Expect(fn.Results).To(HaveLen(1))
				assertHasResult(t, fn.Results, "", "string")
			},
		},
		"SetName exists with pointer receiver": {
			functionName:   "SetName",
			expectedExists: true,
			checkDetails: func(t *testing.T, g *GomegaWithT, fn *model.Function) {
				t.Helper()
				g.Expect(fn.Name).To(Equal("SetName"))
				g.Expect(fn.Receiver.Name()).To(Equal("PersonWithMethods"))
				g.Expect(fn.IsPointerReceiver).To(BeTrue())
				g.Expect(fn.Parameters).To(HaveLen(1))
				assertHasParameter(t, fn.Parameters, "name", "string")
				g.Expect(fn.Results).To(BeEmpty())
			},
		},
		"IsAdult exists": {
			functionName:   "IsAdult",
			expectedExists: true,
			checkDetails: func(t *testing.T, g *GomegaWithT, fn *model.Function) {
				t.Helper()
				g.Expect(fn.Name).To(Equal("IsAdult"))
				g.Expect(fn.Results).To(HaveLen(1))
				assertHasResult(t, fn.Results, "", "bool")
			},
		},
		"UpdateAge exists with pointer receiver and multiple operations": {
			functionName:   "UpdateAge",
			expectedExists: true,
			checkDetails: func(t *testing.T, g *GomegaWithT, fn *model.Function) {
				t.Helper()
				g.Expect(fn.Name).To(Equal("UpdateAge"))
				g.Expect(fn.Receiver.Name()).To(Equal("PersonWithMethods"))
				g.Expect(fn.IsPointerReceiver).To(BeTrue())
				g.Expect(fn.Parameters).To(HaveLen(1))
				assertHasParameter(t, fn.Parameters, "newAge", "int")
				g.Expect(fn.Results).To(HaveLen(1))
				assertHasResult(t, fn.Results, "", "int")
			},
		},
		"Compare exists with named return values": {
			functionName:   "Compare",
			expectedExists: true,
			checkDetails: func(t *testing.T, g *GomegaWithT, fn *model.Function) {
				t.Helper()
				g.Expect(fn.Name).To(Equal("Compare"))
				g.Expect(fn.Parameters).To(HaveLen(1))
				assertHasParameter(t, fn.Parameters, "other", "PersonWithMethods")
				g.Expect(fn.Results).To(HaveLen(2))
				assertHasResult(t, fn.Results, "equal", "bool")
				assertHasResult(t, fn.Results, "ageDiff", "int")
			},
		},
		"Lookup exists with variadic parameters": {
			functionName:   "Lookup",
			expectedExists: true,
			checkDetails: func(t *testing.T, g *GomegaWithT, fn *model.Function) {
				t.Helper()
				g.Expect(fn.Name).To(Equal("Lookup"))
				g.Expect(fn.Parameters).To(HaveLen(1))
				assertHasParameter(t, fn.Parameters, "path", "string")
				// Verify the parameter is marked as variadic
				g.Expect(fn.Parameters[0].IsVariadic).To(BeTrue())
				g.Expect(fn.Results).To(HaveLen(2))
				assertHasResult(t, fn.Results, "", "string")
				assertHasResult(t, fn.Results, "", "bool")
			},
		},
		"NonExistentMethod does not exist": {
			functionName:   "NonExistentMethod",
			expectedExists: false,
		},
	}

	cfg := &config.Config{}
	loader := packageloader.New(cfg, logr.Discard())

	g := NewGomegaWithT(t)
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
					c.checkDetails(t, g, fn)
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
	g.Expect(fn).NotTo(BeNil())

	if fn != nil {
		g.Expect(fn.Description()).NotTo(BeEmpty())
		g.Expect(fn.Description()[0]).To(ContainSubstring("returns the name"))
	}
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
	// 6 from person_with_methods.go + 2 from person_with_methods_extra.go
	g.Expect(functions).To(HaveLen(8))

	// Check that we have the extra methods
	functionNames := make([]string, len(functions))
	for i, fn := range functions {
		functionNames[i] = fn.Name
	}

	g.Expect(functionNames).To(ContainElement("GetFullInfo"))
	g.Expect(functionNames).To(ContainElement("IncrementAge"))
}

// assertHasParameter asserts that a parameter list contains a parameter
// with the given name and type.
func assertHasParameter(t *testing.T, params []model.Parameter, name, typeName string) {
	t.Helper()
	g := NewGomegaWithT(t)
	g.Expect(params).NotTo(BeEmpty(), "parameter list should not be empty")

	found := false

	for _, p := range params {
		if p.Name == name && p.Type.Name() == typeName {
			found = true

			break
		}
	}

	g.Expect(found).To(BeTrue(), "expected to find parameter %s of type %s", name, typeName)
}

// assertHasResult asserts that a result list contains a result with the given name and type.
func assertHasResult(t *testing.T, results []model.Parameter, resultName, typeName string) {
	t.Helper()
	g := NewGomegaWithT(t)
	g.Expect(results).NotTo(BeEmpty(), "result list should not be empty")

	found := false

	for _, r := range results {
		nameMatches := resultName == "" || r.Name == resultName
		if nameMatches && r.Type.Name() == typeName {
			found = true

			break
		}
	}

	if resultName == "" {
		g.Expect(found).To(BeTrue(), "expected to find result of type %s", typeName)
	} else {
		g.Expect(found).To(BeTrue(), "expected to find result %s of type %s", resultName, typeName)
	}
}
