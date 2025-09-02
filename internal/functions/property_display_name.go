package functions

import "github.com/theunrepentantgeek/crddoc/internal/model"

// propertyDisplayName returns the appropriate name for a property based on the configuration.
// If UseGoFieldNames is true, returns the Go field name.
// Otherwise, returns the serialized name (from JSON/YAML tags).
func (f *Functions) propertyDisplayName(property *model.Property) string {
	if f.cfg != nil && f.cfg.UseGoFieldNames {
		return property.Field
	}
	return property.Name
}