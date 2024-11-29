package model

import "github.com/pkg/errors"

// PackageMarkers captures specific package markers read from the source code.
type PackageMarkers struct {
	Name           string       // Name of this package.
	Module         string       // Reference to use when importing this package.
	group          MarkerValue  // Controller-Gen Group of this package.
	version        MarkerValue  // Controller-Gen Version of this package.
	optional       MarkerSwitch // Whether properties are optional by default
	required       MarkerSwitch // Whether properties are required by default
	DefaultGroup   string       // Default group, based on directory name
	DefaultVersion string       // Default version, based on directory name
}

func NewPackageMarkers() *PackageMarkers {
	return &PackageMarkers{
		group:    MakeMarkerValue("groupName"),
		version:  MakeMarkerValue("versionName"),
		optional: MakeMarkerSwitch("kubebuilder", "validation", "Optional"),
		required: MakeMarkerSwitch("kubebuilder", "validation", "Required"),
	}
}

func (m *PackageMarkers) Update(markers *Markers) error {
	if err := m.group.Update(markers); err != nil {
		return errors.Wrap(err, "failed to update package markers group")
	}

	if err := m.version.Update(markers); err != nil {
		return errors.Wrap(err, "failed to update package markers version")
	}

	m.optional.Update(markers)
	m.required.Update(markers)

	return nil
}

func (m *PackageMarkers) Merge(other *PackageMarkers) error {
	if err := m.group.Merge(other.group); err != nil {
		return errors.Wrap(err, "failed to merge package markers group")
	}

	if err := m.version.Merge(other.version); err != nil {
		return errors.Wrap(err, "failed to merge package markers version")
	}

	return nil
}

func (m *PackageMarkers) PropertiesRequiredByDefault() string {
	if m.required.Seen() {
		return "Required"
	}

	if m.optional.Seen() {
		return "Optional"
	}

	return ""
}

// Group returns the group of the package, using the configured controller-runtime marker if set,
// or the directory name if not.
func (m *PackageMarkers) Group() string {
	if grp, ok := m.group.Value(); ok {
		return grp
	}

	return m.DefaultGroup
}

// Version returns the version of the package, using the configured controller-runtime marker if
// set, or the directory name if not.
func (m *PackageMarkers) Version() string {
	if ver, ok := m.version.Value(); ok {
		return ver
	}

	return m.DefaultVersion
}
