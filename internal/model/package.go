package model

import (
	"strings"

	"github.com/go-logr/logr"
	"golang.org/x/exp/maps"
	"golang.org/x/exp/slices"

	"github.com/theunrepentantgeek/crddoc/internal/config"
)

// Package is a struct containing all of the declarations found in a package directory.
type Package struct {
	cfg          *config.Config
	declarations map[string]Declaration // Dictionary of all objects in package, keyed by name.
	ranks        map[string]int         // Dictionary of ranks (depth from root), keyed by name.
	metadata     PackageMetadata
	log          logr.Logger
}

type Order string

const (
	OrderAlphabetical = "alphabetical"
	OrderRanked       = "ranked"
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
		ranks:        make(map[string]int, len(decl)),
		metadata:     metadata,
		log:          log,
	}

	for _, d := range decl {
		result.declarations[d.Name()] = d
	}

	result.catalogCrossReferences()
	result.calculateRanks()
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
		slices.SortFunc(result, p.alphabeticalObjectComparison)
	case OrderRanked:
		// Sort the objects by rank, then alphabetical
		slices.SortFunc(result, p.rankedObjectComparison)
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

// Object returns the object with the given name, if there is one.
func (p *Package) Object(name string) (*Object, bool) {
	dec, ok := p.Declaration(name)
	if !ok {
		return nil, false
	}

	obj, ok := dec.(*Object)
	return obj, ok
}

// Group returns the group of the package.
func (p *Package) Group() string {
	return p.metadata.Group
}

// Version returns the version of the package.
func (p *Package) Version() string {
	return p.metadata.Version
}

// Module returns the module of the package.
func (p *Package) Module() string {
	return p.metadata.Module
}

// Rank returns the usage rank (depth from the root resource) of the given declaration.
func (p *Package) Rank(name string) int {
	return p.ranks[name]
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

func (p *Package) calculateRanks() {
	for name, decl := range p.declarations {
		if decl.Kind() != ResourceDeclaration {
			continue
		}

		p.calculateRanksFromRoot(name, 0)
	}
}

func (p *Package) calculateRanksFromRoot(
	name string,
	rank int,
) {
	if r, ok := p.ranks[name]; ok && r <= rank {
		// We've already walked this declaration and it has a lower rank than we'd give it.
		return
	}

	p.ranks[name] = rank
	decl, ok := p.declarations[name]
	if !ok {
		// Shouldn't happen, but just in case.
		return
	}

	if host, ok := decl.(PropertyContainer); ok {
		for _, prop := range host.Properties() {
			p.calculateRanksFromRoot(prop.Type.Id(), rank+1)
		}
	}
}

func (*Package) alphabeticalObjectComparison(left Declaration, right Declaration) int {
	leftName := strings.ToLower(left.Name())
	rightName := strings.ToLower(right.Name())
	return strings.Compare(leftName, rightName)
}

func (p *Package) rankedObjectComparison(left Declaration, right Declaration) int {
	leftRank := p.ranks[left.Id()]
	rightRank := p.ranks[right.Id()]

	if leftRank < rightRank {
		return -1
	}

	if leftRank > rightRank {
		return 1
	}

	leftName := strings.ToLower(left.Name())
	rightName := strings.ToLower(right.Name())
	return strings.Compare(leftName, rightName)
}
