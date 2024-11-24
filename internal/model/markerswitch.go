package model

// MarkerSwitch captures the presence of a specific marker read from the source code.
type MarkerSwitch struct {
	path []string
	seen bool
}

// MakeMarkerSwitch creates a new MArkerSwitch with the given marker path.
// path is the marker path to look for.
func MakeMarkerSwitch(path ...string) MarkerSwitch {
	return MarkerSwitch{
		path: path,
	}
}

// Update checks for the desired marker in passed set of markers, updating the value if found.
func (m *MarkerSwitch) Update(markers *Markers) {
	// Look for the marker we need
	if _, ok := markers.Lookup(m.path...); ok {
		m.seen = true
	}
}

// Merge combines the supplied MetadataSwitch with the current one,
// returning an error if the values differ.
func (m *MarkerSwitch) Merge(other MarkerSwitch) {
	if other.Seen() {
		m.seen = true
	}
}

// Seen returns true if this marker has been seen, false otherwise.
func (m *MarkerSwitch) Seen() bool {
	return m.seen
}
