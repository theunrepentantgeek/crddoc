package model

import "github.com/pkg/errors"

// MarkerValue captures the value of a specific marker read from the source code.
type MarkerValue struct {
	path  []string
	value *string
}

// MakeMarkerValue creates a new MetadataValue that looks for the given marker path.
// path is the path to the desired value.
func MakeMarkerValue(path ...string) MarkerValue {
	return MarkerValue{
		path: path,
	}
}

// Update reads the value from the passed set of markers, updating the value if found.
// If a new value is found that's different from the current value, we return an error.
func (m *MarkerValue) Update(markers *Markers) error {
	// Look for the marker we need
	marker, ok := markers.Lookup(m.path...)
	if !ok {
		// No value found, nothing to do
		return nil
	}

	return m.SetValue(marker.Value())
}

// Merge combines the supplied MetadataValue with the current one,
// returning an error if the values differ.
func (m *MarkerValue) Merge(other MarkerValue) error {
	if v, ok := other.Value(); ok {
		return m.SetValue(v)
	}

	return nil
}

func (m *MarkerValue) SetValue(value string) error {
	// If we don't have a value, just store the one we're given
	if m.value == nil {
		m.value = &value

		return nil
	}

	// If the value is different, return an error
	if *m.value != value {
		return errors.Errorf(
			"metadata value %q does not match existing value %q",
			value,
			*m.value)
	}

	return nil
}

// Value returns the current value of the metadata, if known.
func (m MarkerValue) Value() (string, bool) {
	if m.value == nil {
		return "", false
	}

	return *m.value, true
}
