package model

import (
	"strings"

	"golang.org/x/exp/slices"
)

func formatComments(
	comments []string,
	name string,
) []string {
	result := slices.Clone(comments)
	if len(comments) == 0 {
		return result
	}

	if s, ok := strings.CutPrefix(result[0], name+": "); ok {
		result[0] = strings.TrimLeft(s, " ")
	}

	return result
}
