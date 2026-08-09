package functions

import "github.com/theunrepentantgeek/crddoc/internal/model"

// asObjectReferenceList wraps property references with template list metadata,
// retaining one reference per referencing object.
func asObjectReferenceList(references []model.PropertyReference) []ListItem[model.PropertyReference] {
	seen := make(map[string]struct{}, len(references))
	result := make([]model.PropertyReference, 0, len(references))

	for _, reference := range references {
		if _, ok := seen[reference.HostID]; ok {
			continue
		}

		seen[reference.HostID] = struct{}{}
		result = append(result, reference)
	}

	return asList(result)
}
