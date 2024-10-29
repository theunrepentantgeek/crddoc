package model

import "strings"

type Markers struct {
	name     string
	value    string
	children map[string]*Markers
}

func NewMarkers() *Markers {
	return &Markers{
		children: make(map[string]*Markers),
	}
}

// Add a marker to the list.
func (m *Markers) Add(marker string) {
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

// Lookup a marker value by path, returning the final value.
func (m *Markers) Lookup(path ...string) (string, bool) {
	if len(path) == 0 {
		return m.value, true
	}

	child, ok := m.children[path[0]]
	if !ok {
		return "", false
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
