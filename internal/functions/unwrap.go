package functions

import "strings"

func (*Functions) unwrap(content []string) string {
	var result strings.Builder

	// Initial conditions chosen so we don't add a leading space every time
	var leadingSpace bool

	trailingSpace := true

	for _, line := range content {
		leadingSpace = strings.HasPrefix(line, " ")
		if !leadingSpace && !trailingSpace {
			result.WriteString(" ")
		}

		if len(strings.TrimSpace(line)) == 0 {
			result.WriteString("<br/>")

			trailingSpace = true
		} else {
			result.WriteString(line)

			trailingSpace = strings.HasSuffix(line, " ")
		}
	}

	return result.String()
}
