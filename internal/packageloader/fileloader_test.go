package packageloader

import (
	"os"
	"path/filepath"
	"testing"

	. "github.com/onsi/gomega"

	"github.com/go-logr/logr"
)

func TestFileLoader_Load_GivenKnownFile_ReturnsExpectedResults(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	fl := NewFileLoader(testdataPath(t, "party_types.go"), logr.Discard())
	g.Expect(fl.Load()).To(Succeed())

	g.Expect(fl.resources).To(HaveLen(1))
	g.Expect(fl.objects).To(HaveLen(3))
	g.Expect(fl.enums).To(HaveLen(1))
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
