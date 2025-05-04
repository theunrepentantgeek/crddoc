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
		cfg:       b.Config,
		resources: make(map[string]*Resource, len(b.Resources)),
		objects:   make(map[string]*Object, len(b.Objects)),
		enums:     make(map[string]*Enum, len(b.Enums)),
		ranks:     make(map[string]int, len(b.Resources)+len(b.Objects)+len(b.Enums)),
		metadata:  b.Metadata,
		log:       b.Log,
	}

	// Add all resources
	for _, d := range b.Resources {
		d.SetPackage(result)
		result.resources[d.Name()] = d
	}

	// Add all objects
	for _, d := range b.Objects {
		d.SetPackage(result)
		result.objects[d.Name()] = d
	}

	// Add all enums
	for _, d := range b.Enums {
		d.SetPackage(result)
		result.enums[d.Name()] = d
	}

	result.catalogCrossReferences()
	result.calculateRanks()

	return result
}
