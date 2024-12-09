package packageloader

import (
	"bufio"
	"os"
	"path"
	"path/filepath"
	"strings"
	"sync"

	kerrors "k8s.io/apimachinery/pkg/util/errors"

	"github.com/go-logr/logr"
	"github.com/pkg/errors"

	"github.com/theunrepentantgeek/crddoc/internal/config"
	"github.com/theunrepentantgeek/crddoc/internal/model"
	"github.com/theunrepentantgeek/crddoc/internal/typefilter"
)

type PackageLoader struct {
	cfg         *config.Config
	log         logr.Logger
	typeFilters *typefilter.List
}

func New(
	cfg *config.Config,
	log logr.Logger,
) *PackageLoader {
	return &PackageLoader{
		cfg:         cfg,
		log:         log,
		typeFilters: typefilter.New(cfg),
	}
}

// LoadDirectory scans a directory for Go files and loads them into a Package.
// folder is the full path to the directory to load.
func (loader *PackageLoader) LoadDirectory(folder string) (*model.Package, error) {
	fqfolder, err := filepath.Abs(folder)
	if err != nil {
		return nil, errors.Wrapf(err, "converting %s to fully qualified path", folder)
	}

	return loader.load(fqfolder, "*.go")
}

// LoadFile loads a single Go file into a  Package.
// file is the full path to the file to load.
func (loader *PackageLoader) LoadFile(file string) (*model.Package, error) {
	dir, name := filepath.Split(file)

	return loader.load(dir, name)
}

func (loader *PackageLoader) load(
	folder string,
	glob string,
) (*model.Package, error) {
	// Start our scan for files to load
	filesToLoad := make(chan string) // fully qualified file paths to load
	errs := make(chan error)         // all errors encountered during loading

	go loader.findFiles(folder, glob, filesToLoad, errs)

	// Start workers to parse the files
	loadedFiles := make(chan *FileLoader) // files after parsing

	var wg sync.WaitGroup

	const numWorkers = 4
	for range numWorkers {
		wg.Add(1)

		go loader.parseFiles(filesToLoad, loadedFiles, errs, &wg)
	}

	// Collect all the parse results together
	packages := make(chan *model.Package) // the final package
	go loader.collectDeclarations(folder, loadedFiles, packages)

	// Accumulate errors
	finalerror := make(chan error) // the final error (if any)
	go loader.collectErrors(errs, finalerror)

	// Wait for all workers to finish
	wg.Wait()

	// Shut down the pipeline
	close(loadedFiles)
	close(errs)

	// Check for any errors
	if err := <-finalerror; err != nil {
		return nil, err
	}

	// Wait for the package to be collected
	pkg := <-packages

	return pkg, nil
}

// findFiles scans for files to load and sends them to the channel.
// Once all files have been found, the channel is closed to signal completion.
//
//nolint:revive // Inherent complexity gives this a score of 8
func (loader *PackageLoader) findFiles(
	folder string,
	glob string,
	filesToLoad chan<- string,
	errs chan<- error,
) {
	defer close(filesToLoad)

	fdr, name := filepath.Split(folder)

	loader.log.Info(
		"Loading package",
		"name", name,
		"folder", fdr)

	files, err := os.ReadDir(folder)
	if err != nil {
		errs <- errors.Wrapf(err, "failed to read directory %s", folder)

		return
	}

	for _, f := range files {
		if f.IsDir() {
			continue
		}

		match, err := filepath.Match(glob, f.Name())
		if err != nil {
			errs <- errors.Wrapf(err, "failed to match file %s with pattern %s", f.Name(), glob)

			continue
		}

		if match {
			fqfn := filepath.Join(folder, f.Name())
			loader.log.V(3).Info("Found file", "file", fqfn)
			filesToLoad <- fqfn
		}
	}
}

func (loader *PackageLoader) parseFiles(
	filesToLoad <-chan string,
	fileLoaders chan<- *FileLoader,
	errs chan<- error,
	wg *sync.WaitGroup,
) {
	defer wg.Done()

	for file := range filesToLoad {
		loader.log.V(2).Info("Parsing file", "file", file)
		fl := NewFileLoader(file, loader.log, loader.typeFilters)

		err := fl.Load()
		if err != nil {
			errs <- errors.Wrapf(err, "failed to load file %s", file)

			continue
		}

		fileLoaders <- fl
	}
}

func (loader *PackageLoader) collectDeclarations(
	folder string,
	loadedFiles <-chan *FileLoader,
	packages chan<- *model.Package,
) {
	metadata := loader.readMetadata(folder)

	var declarations []model.Declaration

	for fl := range loadedFiles {
		loader.log.V(3).Info("Collecting declarations", "file", fl.name)

		decls := fl.Declarations()
		declarations = append(declarations, decls...)

		if err := metadata.Merge(fl.PackageMarkers()); err != nil {
			loader.log.Error(err, "Failed to merge package markers")
		}
	}

	pkg := model.NewPackage(declarations, metadata, loader.cfg, loader.log)
	packages <- pkg
}

func (*PackageLoader) collectErrors(
	errs <-chan error,
	finalerror chan<- error,
) {
	//nolint:prealloc // usual case will be zero errors
	var allErrors []error
	for err := range errs {
		allErrors = append(allErrors, err)
	}

	finalerror <- kerrors.NewAggregate(allErrors)
}

func (loader *PackageLoader) readMetadata(folder string) *model.PackageMarkers {
	// Split keeps the trailing `/` in parent, but Base doesn't care.
	parent, ver := filepath.Split(folder)
	grp := filepath.Base(parent)

	result := model.NewPackageMarkers()
	result.Name = ver
	result.DefaultGroup = grp
	result.DefaultVersion = ver

	if mod, ok := loader.tryReadMetadata(folder); ok {
		result.Module = mod
	}

	return result
}

func (loader *PackageLoader) tryReadMetadata(folder string) (string, bool) {
	// If go.mod exists, read the package path from there
	f, err := filepath.Abs(folder)
	if err != nil {
		return "", false
	}

	if pkg, ok := tryReadMetadataFromFolder(f); ok {
		return pkg, true
	}

	// Didn't find it, look in a parent folder instead
	parent, name := filepath.Split(f)
	if parent != folder {
		if pkg, ok := loader.tryReadMetadata(parent); ok {
			return path.Join(pkg, name), true
		}
	}

	return "", false
}

func tryReadMetadataFromFolder(
	folder string,
) (string, bool) {
	modFilename := filepath.Join(folder, "go.mod")
	if mod, err := os.OpenFile(modFilename, os.O_RDONLY, 0); err == nil {
		defer mod.Close()

		scanner := bufio.NewScanner(mod)
		for scanner.Scan() {
			line := scanner.Text()

			// If line starts with 'module ', return the mod of the line
			if mod, ok := strings.CutPrefix(line, "module "); ok {
				return strings.TrimSpace(mod), true
			}
		}
	}

	return "", false
}
