package model

import (
	"github.com/go-logr/logr"

	"github.com/theunrepentantgeek/crddoc/internal/config"
)

// PackageBuilder is a builder for Package instances.
type PackageBuilder struct {
	Resources  []*Resource
	Objects    map[string]*Object
	Enums      []*Enum
	Interfaces map[string]*Interface
	Metadata   *PackageMarkers
	Config     *config.Config
	Log        logr.Logger
}

// AddResources adds resources to the builder.
func (b *PackageBuilder) AddResources(resources ...*Resource) {
	b.Resources = append(b.Resources, resources...)
}

// AddObjects adds objects to the builder.
func (b *PackageBuilder) AddObjects(objects ...*Object) {
	for _, obj := range objects {
		if _, exists := b.Objects[obj.ID()]; exists {
			b.Log.V(1).Info(
				"Duplicate object ID encountered; overwriting previous object",
				"objectID", obj.ID())
		}

		b.Objects[obj.ID()] = obj
	}
}

// AddEnums adds enums to the builder.
func (b *PackageBuilder) AddEnums(enums ...*Enum) {
	b.Enums = append(b.Enums, enums...)
}

// AddInterfaces adds interfaces to the builder.
func (b *PackageBuilder) AddInterfaces(interfaces ...*Interface) {
	for _, iface := range interfaces {
		if _, exists := b.Interfaces[iface.ID()]; exists {
			b.Log.V(1).Info(
				"Duplicate interface ID encountered; overwriting previous interface",
				"interfaceID", iface.ID())
		}

		b.Interfaces[iface.ID()] = iface
	}
}

// Build creates a new Package from the builder.
func (b *PackageBuilder) Build() *Package {
	l := len(b.Resources) + len(b.Objects) + len(b.Enums) + len(b.Interfaces)
	result := &Package{
		cfg:      b.Config,
		ranks:    make(map[string]int, l),
		metadata: b.Metadata,
		log:      b.Log,
	}

	for _, o := range b.Objects {
		o.SetPackage(result)
	}

	result.objects = b.Objects

	for _, i := range b.Interfaces {
		i.SetPackage(result)
	}

	result.interfaces = b.Interfaces

	result.resources = indexByName(b.Resources, result)
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
