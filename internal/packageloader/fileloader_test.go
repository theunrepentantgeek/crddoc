package packageloader

import (
	"os"
	"path/filepath"
	"testing"

	. "github.com/onsi/gomega"

	"github.com/go-logr/logr"
	"github.com/onsi/gomega/types"

	"github.com/theunrepentantgeek/crddoc/internal/config"
	"github.com/theunrepentantgeek/crddoc/internal/typefilter"
)

func TestFileLoader_Load_GivenPartyFile_ReturnsExpectedImports(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	cfg := &config.Config{}
	filter := typefilter.New(cfg)

	fl := NewFileLoader(testdataPath(t, "party_types.go"), logr.Discard(), filter)
	g.Expect(fl.Load()).To(Succeed())

	g.Expect(fl.importReferences).To(HaveLen(6))
	g.Expect(fl.importReferences).Should(
		HaveKeyWithValue(
			"errors",
			HaveImportPath("errors")))
	g.Expect(fl.importReferences).Should(
		HaveKeyWithValue(
			"fmt",
			HaveImportPath("fmt")))
	g.Expect(fl.importReferences).Should(
		HaveKeyWithValue(
			"metav1",
			HaveImportPath("k8s.io/apimachinery/pkg/apis/meta/v1")))
	g.Expect(fl.importReferences).Should(
		HaveKeyWithValue(
			"kerrors",
			HaveImportPath("k8s.io/apimachinery/pkg/util/errors")))
	g.Expect(fl.importReferences).Should(
		HaveKeyWithValue(
			"conditions",
			HaveImportPath("github.com/Azure/azure-service-operator/v2/pkg/genruntime/conditions")))
}

func HaveImportPath(path string) types.GomegaMatcher {
	return HaveField("ImportPath", Equal(path))
}

func TestFileLoader_Load_GivenPartyFile_ReturnsExpectedResourceCount(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	cfg := &config.Config{}
	filter := typefilter.New(cfg)

	fl := NewFileLoader(testdataPath(t, "party_types.go"), logr.Discard(), filter)
	g.Expect(fl.Load()).To(Succeed())

	g.Expect(fl.resources).To(HaveLen(1))
}

func TestFileLoader_Load_GivenPartyFile_ReturnsExpectedObjectCount(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	cfg := &config.Config{}
	filter := typefilter.New(cfg)

	fl := NewFileLoader(testdataPath(t, "party_types.go"), logr.Discard(), filter)
	g.Expect(fl.Load()).To(Succeed())

	g.Expect(fl.objects).To(HaveLen(3))
}

func TestFileLoader_Load_GivenPartyFile_ReturnsExpectedEnumCount(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	cfg := &config.Config{}
	filter := typefilter.New(cfg)

	fl := NewFileLoader(testdataPath(t, "party_types.go"), logr.Discard(), filter)
	g.Expect(fl.Load()).To(Succeed())

	g.Expect(fl.enums).To(HaveKey("PartyKind"))
	g.Expect(fl.enums).To(HaveLen(1))

	e := fl.enums["PartyKind"]
	g.Expect(e.Values()).To(HaveLen(3))
}

func TestFileLoader_Load_GivenPartyFile_ReturnsExpectedGroup(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	cfg := &config.Config{}
	filter := typefilter.New(cfg)

	testfile := testdataPath(t, "party_types.go")
	fl := NewFileLoader(testfile, logr.Discard(), filter)
	g.Expect(fl.Load()).To(Succeed())

	grp := fl.PackageMarkers().Group()
	g.Expect(grp).To(Equal("colourmodel"))
}

func TestFileLoader_Load_GivenPartyFile_ReturnsExpectedVersion(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	cfg := &config.Config{}
	filter := typefilter.New(cfg)

	fl := NewFileLoader(testdataPath(t, "party_types.go"), logr.Discard(), filter)
	g.Expect(fl.Load()).To(Succeed())

	ver := fl.PackageMarkers().Version()
	g.Expect(ver).To(Equal("v1beta3"))
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
