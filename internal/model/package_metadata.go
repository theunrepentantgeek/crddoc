package model

type PackageMetadata struct {
	Name    string // Name of this package
	Module  string // Reference to use when importing this package
	Group   string // Controller-Gen Group of this package
	Version string // Controller-Gen Version of this package
}

// TrySetGroup will set the group if it is not already set,
// returning true if successful and false if unchanged
func (m *PackageMetadata) TrySetGroup(group string) bool {
	if m.Group == "" || m.Group == group {
		m.Group = group
		return true
	}

	return false
}

// TrySetVersion will set the version if it is not already set,
// returning true if successful and false if unchanged
func (m *PackageMetadata) TrySetVersion(version string) bool {
	if m.Version == "" || m.Version == version {
		m.Version = version
		return true
	}

	return false
}
