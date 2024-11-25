package model

import "github.com/pkg/errors"

// PackageMarkers captures specific package markers read from the source code.
type PackageMarkers struct {
	Name     string       // Name of this package.
	Module   string       // Reference to use when importing this package.
	Group    MarkerValue  // Controller-Gen Group of this package.
	Version  MarkerValue  // Controller-Gen Version of this package.
	optional MarkerSwitch // Whether properties are optional by default
	required MarkerSwitch // Whether properties are required by default
}

func NewPackageMarkers() PackageMarkers {
	return PackageMarkers{
		Group:    MakeMarkerValue("groupName"),
		Version:  MakeMarkerValue("versionName"),
		optional: MakeMarkerSwitch("kubebuilder", "validation", "Optional"),
		required: MakeMarkerSwitch("kubebuilder", "validation", "Required"),
	}
}

func (m *PackageMarkers) Update(markers *Markers) error {
	if err := m.Group.Update(markers); err != nil {
		return errors.Wrap(err, "failed to update package markers group")
	}

	if err := m.Version.Update(markers); err != nil {
		return errors.Wrap(err, "failed to update package markers version")
	}

	m.optional.Update(markers)
	m.required.Update(markers)

	return nil
}

func (m *PackageMarkers) Merge(other PackageMarkers) error {
	if err := m.Group.Merge(other.Group); err != nil {
		return errors.Wrap(err, "failed to merge package markers group")
	}

	if err := m.Version.Merge(other.Version); err != nil {
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
