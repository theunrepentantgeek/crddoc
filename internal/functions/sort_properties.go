package functions

import (
	"slices"
	"strings"

	"github.com/theunrepentantgeek/crddoc/internal/model"
)

// sortProperties sorts a list of properties according to the configuration.
// If UseGoFieldNames is true, sorts by Go field names.
// Otherwise, sorts by serialized names (from JSON/YAML tags).
func (f *Functions) sortProperties(properties []*model.Property) []*model.Property {
	if properties == nil {
		return nil
	}

	// Create a copy to avoid modifying the original slice
	result := slices.Clone(properties)

	// Sort using the appropriate comparison function
	if f.cfg != nil && f.cfg.UseGoFieldNames {
		slices.SortFunc(result, f.alphabeticalPropertyComparisonByField)
	} else {
		slices.SortFunc(result, f.alphabeticalPropertyComparisonByName)
	}

	return result
}

// alphabeticalPropertyComparisonByField does a case insensitive comparison of the Go field names
// of the two properties, allowing them to be sorted.
func (f *Functions) alphabeticalPropertyComparisonByField(left *model.Property, right *model.Property) int {
	leftName := strings.ToLower(left.Field)
	rightName := strings.ToLower(right.Field)

	return strings.Compare(leftName, rightName)
}

// alphabeticalPropertyComparisonByName does a case insensitive comparison of the serialized names
// of the two properties, allowing them to be sorted.
func (f *Functions) alphabeticalPropertyComparisonByName(left *model.Property, right *model.Property) int {
	leftName := strings.ToLower(left.Name)
	rightName := strings.ToLower(right.Name)

	return strings.Compare(leftName, rightName)
}