package util

import "strings"

// TrimPrefix removes the specified prefix from all strings in the slice.
func TrimPrefix(prefix string, strs []string) []string {
	result := make([]string, len(strs))
	for i, str := range strs {
		result[i] = strings.TrimPrefix(str, prefix)
	}

	return result
}
