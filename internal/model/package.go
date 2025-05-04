package model

import (
	"slices"
	"strings"

	"github.com/go-logr/logr"

	"github.com/theunrepentantgeek/crddoc/internal/config"
)

// Package is a struct containing all of the declarations found in a package directory.
type Package struct {
	cfg       *config.Config
	resources map[string]*Resource // Dictionary of resources in package, keyed by name
	objects   map[string]*Object   // Dictionary of objects in package, keyed by name
	enums     map[string]*Enum     // Dictionary of enums in package, keyed by name
	ranks     map[string]int       // Dictionary of ranks (depth from root), keyed by name
	metadata  *PackageMarkers
	log       logr.Logger
}

type Order string

const (
	OrderAlphabetical = "alphabetical"
	OrderRanked       = "ranked"
)

func (p *Package) Name() string {
	return p.metadata.Name
}

func (p *Package) Declarations(order Order) []Declaration {
	if p == nil || (len(p.resources) == 0 && len(p.objects) == 0 && len(p.enums) == 0) {
		return nil
	}

	// Calculate total size for all declarations
	totalDeclarations := len(p.resources) + len(p.objects) + len(p.enums)

	// Collect all declarations into a single slice
	allDeclarations := make([]Declaration, 0, totalDeclarations)

	// Add resources
	for _, res := range p.resources {
		allDeclarations = append(allDeclarations, res)
	}

	// Add objects
	for _, obj := range p.objects {
		allDeclarations = append(allDeclarations, obj)
	}

	// Add enums
	for _, enum := range p.enums {
		allDeclarations = append(allDeclarations, enum)
	}

	// Sort the declarations as specified
	switch order {
	case OrderAlphabetical:
		// Sort the objects alphabetically
		slices.SortFunc(allDeclarations, p.alphabeticalObjectComparison)
	case OrderRanked:
		// Sort the objects by rank, then alphabetical
		slices.SortFunc(allDeclarations, p.rankedObjectComparison)
	}

	return allDeclarations
}

// Declaration returns the declaration with the given name, if found.
func (p *Package) Declaration(name string) (Declaration, bool) {
	if p == nil {
		return nil, false
	}

	if res, ok := p.resources[name]; ok {
		return res, true
	}

	if obj, ok := p.objects[name]; ok {
		return obj, true
	}

	if enum, ok := p.enums[name]; ok {
		return enum, true
	}

	return nil, false
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

// Group returns the group of the package, if known.
func (p *Package) Group() string {
	if grp, ok := p.metadata.group.Value(); ok {
		return grp
	}

	return ""
}

// Version returns the version of the package, if known.
func (p *Package) Version() string {
	if ver, ok := p.metadata.version.Value(); ok {
		return ver
	}

	return ""
}

// Module returns the module of the package.
func (p *Package) Module() string {
	return p.metadata.Module
}

func (p *Package) PropertiesRequiredByDefault() string {
	return p.metadata.PropertiesRequiredByDefault()
}

// Rank returns the usage rank (depth from the root resource) of the given declaration.
func (p *Package) Rank(name string) int {
	return p.ranks[name]
}

func (p *Package) catalogCrossReferences() {
	// canSetUsage is a marker interface for declarations that can keep track of their usage.
	type canSetUsage interface {
		SetUsage(refs []PropertyReference)
	}

	usages := p.indexUsage()
	for name, usage := range usages {
		if decl, ok := p.Declaration(name); !ok {
			if can, ok := decl.(canSetUsage); ok {
				slices.SortFunc(usage, ComparePropertyReferences)
				can.SetUsage(usage)
			}
		}
	}
}

func (p *Package) indexUsage() map[string][]PropertyReference {
	result := make(map[string][]PropertyReference)

	// Handle resources
	for _, dec := range p.resources {
		p.addPropertyReferences(dec.ID(), dec.Name(), dec, result)
	}

	// Handle objects
	for _, dec := range p.objects {
		p.addPropertyReferences(dec.ID(), dec.Name(), dec, result)
	}

	// Enums don't have properties to reference, so we skip them

	return result
}

func (p *Package) addPropertyReferences(
	hostID string,
	name string,
	container PropertyContainer,
	refs map[string][]PropertyReference,
) {
	for _, prop := range container.Properties() {
		id := prop.Type.ID()
		if _, ok := p.Declaration(id); ok {
			ref := NewPropertyReference(name, hostID, prop.Name)
			refs[id] = append(refs[id], ref)
		}
	}
}

// calculateRanks calculates the ranks of all declarations in the package.
// The rank is the depth from the root resource, with resources having a rank of 0.
func (p *Package) calculateRanks() {
	// Only resources should have rank 0 (root)
	for name, decl := range p.resources {
		if decl.Kind() == ResourceDeclaration {
			p.calculateRanksFromRoot(name, 0)
		}
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

	decl, ok := p.Declaration(name)
	if !ok {
		return
	}

	ctr, ok := decl.(PropertyContainer)
	if !ok {
		return
	}

	// Walk through the properties of this declaration
	for _, prop := range ctr.Properties() {
		p.calculateRanksFromRoot(prop.Type.ID(), rank+1)
	}
}

func (*Package) alphabeticalObjectComparison(left Declaration, right Declaration) int {
	leftName := strings.ToLower(left.Name())
	rightName := strings.ToLower(right.Name())

	return strings.Compare(leftName, rightName)
}

func (p *Package) rankedObjectComparison(left Declaration, right Declaration) int {
	leftRank := p.ranks[left.ID()]
	rightRank := p.ranks[right.ID()]

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
