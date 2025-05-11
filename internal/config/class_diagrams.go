package config

// ClassDiagrams captures all the configuration available for class diagrams.
type ClassDiagrams struct {
	// Enabled turns class diagram generation on.
	// Defaults to true if any class diagram options are set, false otherwise.
	Enabled *bool `yaml:"enabled"`
}

// Empty returns true if no class diagram options are set.
func (c *ClassDiagrams) Empty() bool {
	return c.Enabled == nil
}
