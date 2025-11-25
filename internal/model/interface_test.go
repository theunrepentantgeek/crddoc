package model_test

import (
	"testing"

	. "github.com/onsi/gomega"

	"github.com/go-logr/logr"

	"github.com/theunrepentantgeek/crddoc/internal/config"
	"github.com/theunrepentantgeek/crddoc/internal/model"
	"github.com/theunrepentantgeek/crddoc/internal/packageloader"
)

func TestInterface_TryNewInterface_ParsesInterface(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	cfg := &config.Config{}
	loader := packageloader.New(cfg, logr.Discard())

	pkg, err := loader.LoadFile(testdataPath(t, "interfaces_types.go"))
	g.Expect(err).To(Succeed())
	g.Expect(pkg).NotTo(BeNil())

	// Check that Greeter interface exists
	dec, ok := pkg.Declaration("Greeter")
	g.Expect(ok).To(BeTrue())
	g.Expect(dec).NotTo(BeNil())
	g.Expect(dec.Kind()).To(Equal(model.InterfaceDeclaration))

	iface, ok := dec.(*model.Interface)
	g.Expect(ok).To(BeTrue())
	g.Expect(iface.Name()).To(Equal("Greeter"))
}

func TestInterface_Methods_ReturnsExpectedMethods(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	cfg := &config.Config{}
	loader := packageloader.New(cfg, logr.Discard())

	pkg, err := loader.LoadFile(testdataPath(t, "interfaces_types.go"))
	g.Expect(err).To(Succeed())
	g.Expect(pkg).NotTo(BeNil())

	// Check Speaker interface methods
	iface, ok := pkg.Interface("Speaker")
	g.Expect(ok).To(BeTrue())
	g.Expect(iface).NotTo(BeNil())

	methods := iface.Methods()
	g.Expect(methods).To(HaveLen(2))

	// Check method names (sorted alphabetically)
	methodNames := make([]string, len(methods))
	for i, m := range methods {
		methodNames[i] = m.Name
	}
	g.Expect(methodNames).To(ConsistOf("Speak", "Volume"))
}

func TestInterface_Method_ReturnsExpectedMethod(t *testing.T) {
	t.Parallel()

	cases := map[string]struct {
		interfaceName  string
		methodName     string
		expectedExists bool
		checkDetails   func(*testing.T, *GomegaWithT, *model.Function)
	}{
		"Greeter.Greet exists": {
			interfaceName:  "Greeter",
			methodName:     "Greet",
			expectedExists: true,
			checkDetails: func(t *testing.T, g *GomegaWithT, fn *model.Function) {
				t.Helper()
				g.Expect(fn.Name).To(Equal("Greet"))
				g.Expect(fn.Parameters).To(BeEmpty())
				g.Expect(fn.Results).To(HaveLen(1))
			},
		},
		"Speaker.Speak exists with parameter": {
			interfaceName:  "Speaker",
			methodName:     "Speak",
			expectedExists: true,
			checkDetails: func(t *testing.T, g *GomegaWithT, fn *model.Function) {
				t.Helper()
				g.Expect(fn.Name).To(Equal("Speak"))
				g.Expect(fn.Parameters).To(HaveLen(1))
				g.Expect(fn.Parameters[0].Name).To(Equal("message"))
				g.Expect(fn.Results).To(HaveLen(1))
			},
		},
		"Speaker.Volume exists with no parameters": {
			interfaceName:  "Speaker",
			methodName:     "Volume",
			expectedExists: true,
			checkDetails: func(t *testing.T, g *GomegaWithT, fn *model.Function) {
				t.Helper()
				g.Expect(fn.Name).To(Equal("Volume"))
				g.Expect(fn.Parameters).To(BeEmpty())
				g.Expect(fn.Results).To(HaveLen(1))
			},
		},
		"Speaker.NonExistent does not exist": {
			interfaceName:  "Speaker",
			methodName:     "NonExistent",
			expectedExists: false,
		},
	}

	cfg := &config.Config{}
	loader := packageloader.New(cfg, logr.Discard())

	g := NewGomegaWithT(t)
	pkg, err := loader.LoadFile(testdataPath(t, "interfaces_types.go"))
	g.Expect(err).To(Succeed())
	g.Expect(pkg).NotTo(BeNil())

	for n, c := range cases {
		t.Run(n, func(t *testing.T) {
			t.Parallel()

			g := NewGomegaWithT(t)

			iface, ok := pkg.Interface(c.interfaceName)
			g.Expect(ok).To(BeTrue())

			method, ok := iface.Method(c.methodName)
			g.Expect(ok).To(Equal(c.expectedExists))

			if c.expectedExists && c.checkDetails != nil {
				g.Expect(method).NotTo(BeNil())
				c.checkDetails(t, g, method)
			}
		})
	}
}

func TestInterface_Implementations_ReturnsImplementingObjects(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	cfg := &config.Config{}
	loader := packageloader.New(cfg, logr.Discard())

	pkg, err := loader.LoadFile(testdataPath(t, "interfaces_types.go"))
	g.Expect(err).To(Succeed())
	g.Expect(pkg).NotTo(BeNil())

	// Check Greeter implementations
	greeter, ok := pkg.Interface("Greeter")
	g.Expect(ok).To(BeTrue())
	g.Expect(greeter).NotTo(BeNil())

	impls := greeter.Implementations()
	g.Expect(impls).To(HaveLen(2))

	// Check implementation names (sorted alphabetically)
	implNames := make([]string, len(impls))
	for i, impl := range impls {
		implNames[i] = impl.Name()
	}
	g.Expect(implNames).To(ConsistOf("Human", "Robot"))
}

func TestObject_ImplementsInterfaces_ReturnsImplementedInterfaces(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	cfg := &config.Config{}
	loader := packageloader.New(cfg, logr.Discard())

	pkg, err := loader.LoadFile(testdataPath(t, "interfaces_types.go"))
	g.Expect(err).To(Succeed())
	g.Expect(pkg).NotTo(BeNil())

	// Check Human implements both Greeter and Speaker
	human, ok := pkg.Object("Human")
	g.Expect(ok).To(BeTrue())
	g.Expect(human).NotTo(BeNil())

	interfaces := human.ImplementsInterfaces()
	g.Expect(interfaces).To(HaveLen(2))

	// Check interface names (sorted alphabetically)
	ifaceNames := make([]string, len(interfaces))
	for i, iface := range interfaces {
		ifaceNames[i] = iface.Name()
	}
	g.Expect(ifaceNames).To(ConsistOf("Greeter", "Speaker"))

	// Check Robot implements only Greeter
	robot, ok := pkg.Object("Robot")
	g.Expect(ok).To(BeTrue())
	g.Expect(robot).NotTo(BeNil())

	robotInterfaces := robot.ImplementsInterfaces()
	g.Expect(robotInterfaces).To(HaveLen(1))
	g.Expect(robotInterfaces[0].Name()).To(Equal("Greeter"))
}

func TestInterface_Description_ParsesComments(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	cfg := &config.Config{}
	loader := packageloader.New(cfg, logr.Discard())

	pkg, err := loader.LoadFile(testdataPath(t, "interfaces_types.go"))
	g.Expect(err).To(Succeed())
	g.Expect(pkg).NotTo(BeNil())

	// Check that Greeter has a description
	iface, ok := pkg.Interface("Greeter")
	g.Expect(ok).To(BeTrue())
	g.Expect(iface).NotTo(BeNil())
	g.Expect(iface.Description()).NotTo(BeEmpty())
	g.Expect(iface.Description()[0]).To(ContainSubstring("interface for things that can greet"))
}

func TestPackage_Declarations_IncludesInterfaces(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	cfg := &config.Config{}
	loader := packageloader.New(cfg, logr.Discard())

	pkg, err := loader.LoadFile(testdataPath(t, "interfaces_types.go"))
	g.Expect(err).To(Succeed())
	g.Expect(pkg).NotTo(BeNil())

	declarations := pkg.Declarations(model.OrderAlphabetical)
	g.Expect(declarations).NotTo(BeNil())

	// Find interfaces in declarations
	interfaceCount := 0
	for _, decl := range declarations {
		if decl.Kind() == model.InterfaceDeclaration {
			interfaceCount++
		}
	}
	g.Expect(interfaceCount).To(Equal(3)) // Greeter, Speaker, MultiTalent
}
