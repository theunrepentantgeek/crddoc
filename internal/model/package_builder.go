package model

import (
	"github.com/go-logr/logr"

	"github.com/theunrepentantgeek/crddoc/internal/config"
)

// PackageBuilder is a builder for Package instances.
type PackageBuilder struct {
	Resources []*Resource
	Objects   []*Object
	Enums     []*Enum
	Metadata  *PackageMarkers
	Config    *config.Config
	Log       logr.Logger
}

// Build creates a new Package from the builder.
func (b *PackageBuilder) Build() *Package {
	result := &Package{
		cfg:      b.Config,
		ranks:    make(map[string]int, len(b.Resources)+len(b.Objects)+len(b.Enums)),
		metadata: b.Metadata,
		log:      b.Log,
	}

	result.resources = indexByName(b.Resources, result)
	result.objects = indexByName(b.Objects, result)
	result.enums = indexByName(b.Enums, result)

	result.catalogCrossReferences()
	result.calculateRanks()

	return result
}

type indexable interface {
	Name() string
	SetPackage(pkg *Package)
}

// indexByName creates a map of items indexed by their name.
func indexByName[N indexable](items []N, pkg *Package) map[string]N {
	result := make(map[string]N, len(items))

	for _, item := range items {
		item.SetPackage(pkg)
		result[item.Name()] = item
	}

	return result
}
