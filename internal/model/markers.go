package model

import "strings"

type Markers struct {
	name     string
	value    string
	children map[string]*Markers
}

func NewMarkers(markers ...string) *Markers {
	result := &Markers{
		children: make(map[string]*Markers),
	}

	for _, m := range markers {
		result.Add(m)
	}

	return result
}

func (m *Markers) Value() string {
	return m.value
}

// Add a marker to the list.
func (m *Markers) Add(marker string) {
	// Trim any leading whitespace or comment markers
	marker = strings.TrimPrefix(marker, "//")
	marker = strings.TrimPrefix(marker, " ")
	marker = strings.TrimPrefix(marker, "+")

	if n, rest, ok := strings.Cut(marker, ":"); ok {
		child := m.requireChild(n)
		if rest != "" {
			child.Add(rest)
		}

		return
	}

	if n, v, ok := strings.Cut(marker, "="); ok {
		child := m.requireChild(n)
		child.value = v
	} else {
		m.requireChild(marker)
	}
}

func (m *Markers) Any() bool {
	return len(m.children) > 0
}

// Lookup a marker value by path, returning the final marker.
func (m *Markers) Lookup(path ...string) (Markers, bool) {
	if len(path) == 0 {
		return *m, true
	}

	child, ok := m.children[path[0]]
	if !ok {
		return Markers{}, false
	}

	return child.Lookup(path[1:]...)
}

// Exists returns true if the marker exists.
func (m *Markers) Exists(path ...string) bool {
	_, ok := m.Lookup(path...)

	return ok
}

func (m *Markers) requireChild(name string) *Markers {
	result, ok := m.children[name]
	if !ok {
		result = NewMarkers()
		result.name = name
		m.children[name] = result
	}

	return result
}
