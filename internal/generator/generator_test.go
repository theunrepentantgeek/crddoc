package generator

import (
	"os"
	"path/filepath"
	"testing"

	. "github.com/onsi/gomega"

	"github.com/go-logr/logr"

	"github.com/theunrepentantgeek/crddoc/internal/config"
	"github.com/theunrepentantgeek/crddoc/internal/packageloader"
)

func TestGenerator_GenerateToMultipleFiles_CreatesCorrectFiles(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	// Arrange
	cfg := config.Standard()
	cfg.Mode = "multiple-file"
	log := logr.Discard()

	loader := packageloader.New(cfg, log)
	pkg, err := loader.LoadDirectory("../packageloader/testdata")
	g.Expect(err).ToNot(HaveOccurred())

	gen := New(cfg, log)
	err = gen.LoadTemplates()
	g.Expect(err).ToNot(HaveOccurred())

	// Create a temporary directory
	tempDir := t.TempDir()

	// Act
	err = gen.GenerateToMultipleFiles(pkg, tempDir, log)

	// Assert
	g.Expect(err).ToNot(HaveOccurred())

	// Check that files were created
	files, err := os.ReadDir(tempDir)
	g.Expect(err).ToNot(HaveOccurred())
	g.Expect(files).To(HaveLen(4)) // PartyReference, PartyResource, PartyResourceSpec, PartyResourceStatus

	// Check that specific expected files exist
	expectedFiles := []string{
		"PartyReference.md",
		"PartyResource.md", 
		"PartyResourceSpec.md",
		"PartyResourceStatus.md",
	}

	for _, expectedFile := range expectedFiles {
		filePath := filepath.Join(tempDir, expectedFile)
		_, err := os.Stat(filePath)
		g.Expect(err).ToNot(HaveOccurred(), "Expected file %s to exist", expectedFile)

		// Check that the file has content
		content, err := os.ReadFile(filePath)
		g.Expect(err).ToNot(HaveOccurred())
		g.Expect(len(content)).To(BeNumerically(">", 0), "Expected file %s to have content", expectedFile)
	}
}

func TestGenerator_GenerateToFile_SingleFileMode_StillWorks(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	// Arrange
	cfg := config.Standard()
	cfg.Mode = "single-file"
	log := logr.Discard()

	loader := packageloader.New(cfg, log)
	pkg, err := loader.LoadDirectory("../packageloader/testdata")
	g.Expect(err).ToNot(HaveOccurred())

	gen := New(cfg, log)
	err = gen.LoadTemplates()
	g.Expect(err).ToNot(HaveOccurred())

	// Create a temporary file
	tempDir := t.TempDir()
	outputFile := filepath.Join(tempDir, "output.md")

	// Act
	err = gen.GenerateToFile(pkg, outputFile, log)

	// Assert
	g.Expect(err).ToNot(HaveOccurred())

	// Check that the file was created and has content
	content, err := os.ReadFile(outputFile)
	g.Expect(err).ToNot(HaveOccurred())
	g.Expect(len(content)).To(BeNumerically(">", 0))
	g.Expect(string(content)).To(ContainSubstring("testdata"))
	g.Expect(string(content)).To(ContainSubstring("PartyResource"))
}