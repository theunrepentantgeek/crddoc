package model

import "strings"

type Markers struct {
	markers [][]string
}

func NewMarkers() *Markers {
	return &Markers{}
}

// Add a marker to the list
func (m *Markers) Add(marker string) {
	// Split the marker into a path
	path := strings.Split(
		strings.TrimPrefix(marker, "+"),
		":")

	// Trim the path
	for i, p := range path {
		path[i] = strings.Trim(p, " ")
	}

	// Add to our list
	m.markers = append(m.markers, path)
}

func (m *Markers) Count() int {
	return len(m.markers)
}

// Lookup a marker value by path, returning the final value
func (m *Markers) Lookup(path ...string) (string, bool) {
	// Iterate over the markers
	for _, marker := range m.markers {
		// If the marker path matches the lookup path, return the value
		if len(marker) == len(path)+1 {
			returnValue, shouldReturn := checkMarker(marker, path)
			if shouldReturn {
				return returnValue, true
			}
		}
	}

	// No match
	return "", false
}

// Exists returns true if the marker exists
func (m *Markers) Exists(path ...string) bool {
	// Iterate over the markers
	for _, marker := range m.markers {
		// If the marker path matches the lookup path, return the value
		if len(marker) == len(path) {
			_, shouldReturn := checkMarker(marker, path)
			if shouldReturn {
				return true
			}
		}
	}

	// No match
	return false
}

func checkMarker(marker []string, path []string) (string, bool) {
	for i, p := range path {
		if marker[i] != p {
			return "", false
		}
	}

	return marker[len(marker)-1], true
}
