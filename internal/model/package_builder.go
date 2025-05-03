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
	// Calculate total size for all declarations
	totalDeclarations := len(b.Resources) + len(b.Objects) + len(b.Enums)

	result := &Package{
		cfg:          b.Config,
		declarations: make(map[string]Declaration, totalDeclarations),
		ranks:        make(map[string]int, totalDeclarations),
		metadata:     b.Metadata,
		log:          b.Log,
	}

	// Add all resources
	for _, d := range b.Resources {
		d.SetPackage(result)
		result.declarations[d.Name()] = d
	}

	// Add all objects
	for _, d := range b.Objects {
		d.SetPackage(result)
		result.declarations[d.Name()] = d
	}

	// Add all enums
	for _, d := range b.Enums {
		d.SetPackage(result)
		result.declarations[d.Name()] = d
	}

	result.catalogCrossReferences()
	result.calculateRanks()

	return result
}
