package config

// ClassDiagram captures all the configuration available for class diagrams.
type ClassDiagram struct {
	// Enabled turns class diagram generation on.
	// Defaults to true if any class diagram options are set, false otherwise.
	Enabled *bool `yaml:"enabled"`

	// DependencySplitThreshold is the threshold above which dependencies on the class diagram
	// are split into two groups, either above/below the root object, or left/right.
	// Defaults to 6.
	// Splitting into two groups helps to reduce the aspect ration of the diagram, increasing
	// readability.
	DependencySplitThreshold *int `yaml:"dependencySplitThreshold"`
}

// Empty returns true if no class diagram options are set.
func (c *ClassDiagram) Empty() bool {
	return c.Enabled == nil &&
		c.DependencySplitThreshold == nil
}
