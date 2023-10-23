package packageloader

import (
	"os"
	"path/filepath"
	"sync"

	"github.com/go-logr/logr"
	"github.com/pkg/errors"
	"github.com/theunrepentantgeek/crddoc/internal/config"
	"github.com/theunrepentantgeek/crddoc/internal/model"
	"golang.org/x/sync/errgroup"
)

type PackageLoader struct {
	cfg     *config.Config
	log     logr.Logger
	name    string // Name of this package
	ref     string // Reference to use when importing this package
	group   string // Group of this package
	version string // Version of this package

}

func New(
	cfg *config.Config,
	log logr.Logger,
) *PackageLoader {
	return &PackageLoader{
		cfg: cfg,
		log: log,
	}
}

// LoadDirectory scans a directory for Go files and loads them into a Package.
// folder is the full path to the directory to load.
func (loader *PackageLoader) LoadDirectory(folder string) (*model.Package, error) {
	return loader.load(folder, "*.go")
}

// LoadFile loads a single Go file into a  Package.
// file is the full path to the file to load.
func (loader *PackageLoader) LoadFile(file string) (*model.Package, error) {
	dir, name := filepath.Split(file)
	return loader.load(dir, name)
}

func (loader *PackageLoader) load(folder string, glob string) (*model.Package, error) {
	fdr, name := filepath.Split(folder)
	group := filepath.Base(fdr)

	loader.log.Info(
		"Loading package",
		"name", name,
		"folder", fdr)

	var lock sync.Mutex

	loader.name = name
	loader.group = group
	loader.version = name

	files, err := os.ReadDir(folder)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to read directory %s", folder)
	}

	var declarations []model.Declaration
	var metadata model.PackageMetadata
	var eg errgroup.Group
	for _, f := range files {
		f := f

		if f.IsDir() {
			continue
		}

		match, err := filepath.Match(glob, f.Name())
		if err != nil {
			return nil, errors.Wrapf(err, "failed to match file %s with pattern %s", f.Name(), glob)
		}
		if !match {
			continue
		}

		eg.Go(func() error {
			var path = filepath.Join(folder, f.Name())
			fl := NewFileLoader(path, loader.log)
			err := fl.Load()
			if err != nil {
				return err
			}

			decls := fl.Declarations()
			lock.Lock()
			defer lock.Unlock()

			declarations = append(declarations, decls...)
			if grp, ok := fl.Group(); ok {
				if !metadata.TrySetGroup(grp) {
					loader.log.Info(
						"Multiple values for 'group' found in package",
						"file", path,
						"existing", metadata.Group,
						"new", grp)
				}
			}

			if ver, ok := fl.Version(); ok {
				if !metadata.TrySetVersion(ver) {
					loader.log.Info(
						"Multiple values for 'version' found in package",
						"file", path,
						"existing", metadata.Version,
						"new", ver)
				}
			}

			return nil
		})
	}

	if err := eg.Wait(); err != nil {
		return nil, errors.Wrapf(err, "failed to load directory %s", folder)
	}

	pkg := model.NewPackage(declarations, metadata, loader.cfg, loader.log)

	return pkg, nil
}
