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
	fileMode := config.FileModeMultipleFile
	cfg := config.Standard()
	cfg.SetFileMode(&fileMode)

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
	// Expected: PartyReference, PartyResource, PartyResourceSpec, PartyResourceStatus
	g.Expect(files).To(HaveLen(4))

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
		g.Expect(content).ToNot(BeEmpty(), "Expected file %s to have content", expectedFile)
	}
}

func TestGenerator_GenerateToFile_SingleFileMode_StillWorks(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	// Arrange
	cfg := config.Standard()
	fileMode := config.FileModeSingleFile
	cfg.SetFileMode(&fileMode)

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
	g.Expect(content).ToNot(BeEmpty())
	g.Expect(string(content)).To(ContainSubstring("testdata"))
	g.Expect(string(content)).To(ContainSubstring("PartyResource"))
}

func TestGenerator_GenerateToFile_WithIncludeFunctions_IncludesFunctions(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	// Arrange
	cfg := config.Standard()
	cfg.IncludeFunctions = true

	log := logr.Discard()

	loader := packageloader.New(cfg, log)
	pkg, err := loader.LoadDirectory("../model/testdata")
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
	g.Expect(content).ToNot(BeEmpty())

	// Check that functions section is present
	g.Expect(string(content)).To(ContainSubstring("### Functions"))
	g.Expect(string(content)).To(ContainSubstring("GetName"))
	g.Expect(string(content)).To(ContainSubstring("SetName"))
	g.Expect(string(content)).To(ContainSubstring("IsAdult"))
}

func TestGenerator_GenerateToFile_WithoutIncludeFunctions_DoesNotIncludeFunctions(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	// Arrange
	cfg := config.Standard()
	cfg.IncludeFunctions = false

	log := logr.Discard()

	loader := packageloader.New(cfg, log)
	pkg, err := loader.LoadDirectory("../model/testdata")
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
	g.Expect(content).ToNot(BeEmpty())

	// Check that functions section is NOT present
	g.Expect(string(content)).ToNot(ContainSubstring("### Functions"))
	g.Expect(string(content)).ToNot(ContainSubstring("GetName"))
}
