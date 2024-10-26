package model

import "strings"

// ParseComments iterates over the comments and returns the description and any commands that are
// not part of the description
// TODO: Add optional varidic "cleanup" functions to allow for more complex comment parsing/cleanup.
func ParseComments(comments []string) ([]string, *Markers) {
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
