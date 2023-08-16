package model

import "strings"

// parseComments iterates over the comments and returns the description and any comments that are
// not part of the description
func parseComments(comments []string) ([]string, *Markers) {
	description := make([]string, 0, len(comments))
	commands := NewMarkers()

	for _, c := range comments {
		c = strings.TrimPrefix(c, "//")
		c = strings.TrimPrefix(c, " ")
		if strings.HasPrefix(c, "+") {
			commands.Add(c)
		} else {
			description = append(description, c)
		}
	}

	return description, commands
}

// "+kubebuilder:validation:Required"
//func parseMarker(commands []string)
