package packageloader

import (
	"os"
	"path/filepath"
	"testing"

	. "github.com/onsi/gomega"
	"github.com/theunrepentantgeek/crddoc/internal/config"
	"github.com/theunrepentantgeek/crddoc/internal/typefilter"

	"github.com/go-logr/logr"
)

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

	fl := NewFileLoader(testdataPath(t, "party_types.go"), logr.Discard(), filter)
	g.Expect(fl.Load()).To(Succeed())

	grp, ok := fl.Group()
	g.Expect(ok).To(BeTrue())
	g.Expect(grp).To(Equal("colourmodel"))
}

func TestFileLoader_Load_GivenPartyFile_ReturnsExpectedVersion(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	cfg := &config.Config{}
	filter := typefilter.New(cfg)

	fl := NewFileLoader(testdataPath(t, "party_types.go"), logr.Discard(), filter)
	g.Expect(fl.Load()).To(Succeed())

	ver, ok := fl.Version()
	g.Expect(ok).To(BeTrue())
	g.Expect(ver).To(Equal("v1beta3"))
}

// loadTestData is a helper used to load a testdata source file
func testdataPath(
	t *testing.T,
	filename string, // nolint:unparam // other file names will be used as testing improves
) string {
	t.Helper()

	wd, err := os.Getwd()
	if err != nil {
		t.Fatalf("Failed to get working directory: %v", err)
	}

	result := filepath.Join(wd, "testdata", filename)
	return result
}
