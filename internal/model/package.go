package model

import (
	"github.com/theunrepentantgeek/crddoc/internal/config"

	"github.com/go-logr/logr"
	"golang.org/x/exp/maps"
	"golang.org/x/exp/slices"
)

// Package is a struct containing all of the declarations found in a package directory
type Package struct {
	cfg          *config.Config
	declarations map[string]Declaration // Dictionary of all objects in package, keyed by name
	metadata     PackageMetadata
	log          logr.Logger
}

type Order string

const (
	OrderAlphabetical = "alphabetical"
)

func NewPackage(
	decl []Declaration,
	metadata PackageMetadata,
	cfg *config.Config,
	log logr.Logger,
) *Package {
	result := &Package{
		cfg:          cfg,
		declarations: make(map[string]Declaration, len(decl)),
		metadata:     metadata,
		log:          log,
	}

	for _, d := range decl {
		result.declarations[d.Name()] = d
	}

	result.catalogCrossReferences()
	return result
}

func (p *Package) Name() string {
	return p.metadata.Name
}

func (p *Package) Declarations(order Order) []Declaration {
	result := maps.Values(p.declarations)

	// Sort the declarations as specified
	switch order {
	case OrderAlphabetical:
		// Sort the objects alphabetically
		slices.SortFunc(result, alphabeticalObjectComparison)
	}

	return result
}

// Declaration returns the declaration with the given name, if found.
func (p *Package) Declaration(name string) (Declaration, bool) {
	dec, ok := p.declarations[name]
	if !ok {
		return nil, false
	}

	return dec, ok
}

// Object returns the object with the given name, if there is one
func (p *Package) Object(name string) (*Object, bool) {
	dec, ok := p.Declaration(name)
	if !ok {
		return nil, false
	}

	obj, ok := dec.(*Object)
	return obj, ok
}

// Group returns the group of the package
func (p *Package) Group() string {
	return p.metadata.Group
}

// Version returns the version of the package
func (p *Package) Version() string {
	return p.metadata.Version
}

// Module returns the module of the package
func (p *Package) Module() string {
	return p.metadata.Module
}

func (p *Package) catalogCrossReferences() {
	usages := p.indexUsage()
	for name, usage := range usages {
		if obj, ok := p.declarations[name]; ok {
			slices.SortFunc(usage, ComparePropertyReferences)
			obj.SetUsage(usage)
		}
	}
}

func (p *Package) indexUsage() map[string][]PropertyReference {
	result := make(map[string][]PropertyReference)
	for _, dec := range p.declarations {
		// Index references from an object
		if host, ok := dec.(PropertyContainer); ok {
			for _, prop := range host.Properties() {
				id := prop.Type.Id()
				if _, ok := p.declarations[id]; ok {
					ref := NewPropertyReference(dec.Name(), dec.Id(), prop.Name)
					result[id] = append(result[id], ref)
				}
			}
		}
	}

	return result
}

func alphabeticalObjectComparison(left Declaration, right Declaration) int {
	if left.Name() < right.Name() {
		return -1
	}

	if left.Name() > right.Name() {
		return 1
	}

	return 0
}
