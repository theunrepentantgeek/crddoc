package packageloader

import (
	"os"
	"testing"

	. "github.com/onsi/gomega"
)

func Test_TryReadMetadata_WhenModFileMissing_ReturnsFalse(t *testing.T) {
	g := NewGomegaWithT(t)

	loader := PackageLoader{}
	_, ok := loader.tryReadMetadata("/testdata/missing")
	g.Expect(ok).To(BeFalse())
}

func Test_TryReadMetadata_WhenModFilePresent_ReturnsExpectedResults(t *testing.T) {
	g := NewGomegaWithT(t)

	loader := PackageLoader{}
	cwd, err := os.Getwd()
	g.Expect(err).NotTo(HaveOccurred())

	pkg, ok := loader.tryReadMetadata(cwd)
	g.Expect(ok).To(BeTrue())
	g.Expect(pkg).To(Equal("github.com/theunrepentantgeek/crddoc/internal/packageloader"))
}
