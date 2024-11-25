package model

import (
	"github.com/dave/dst"
	"github.com/pkg/errors"
)

type PropertyMarkers struct {
	optional MarkerSwitch // Whether properties are optional by default
	required MarkerSwitch // Whether properties are required by default
}

func NewPropertyMarkers() *PropertyMarkers {
	return &PropertyMarkers{
		optional: MakeMarkerSwitch("kubebuilder", "validation", "Optional"),
		required: MakeMarkerSwitch("kubebuilder", "validation", "Required"),
	}
}

func (m *PropertyMarkers) Parse(
	markers *Markers,
) error {
	m.optional.Update(markers)
	m.required.Update(markers)

	return nil
}

func (m *PropertyMarkers) ParseDecorations(decs dst.NodeDecs) error {
	leadingMarkers := NewMarkers(decs.Start...)
	if err := m.Parse(leadingMarkers); err != nil {
		return errors.Wrap(err, "parsing leading decorations")
	}

	trailingMarkers := NewMarkers(decs.End...)
	if err := m.Parse(trailingMarkers); err != nil {
		return errors.Wrap(err, "parsing trailing decorations")
	}

	return nil
}

func (m *PropertyMarkers) Optional() bool {
	return m.optional.Seen()
}

func (m *PropertyMarkers) Required() bool {
	return m.required.Seen()
}

func (m *PropertyMarkers) Merge(other *PropertyMarkers) error {
	m.optional.Merge(other.optional)
	m.required.Merge(other.required)

	return nil
}
