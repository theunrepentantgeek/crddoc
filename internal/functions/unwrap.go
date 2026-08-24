package functions

import "strings"

func (*Functions) unwrap(content []string) string {
	if !containsList(content) {
		return unwrapText(content)
	}

	blocks := splitCommentBlocks(content)
	result := make([]string, 0, len(blocks))

	for _, block := range blocks {
		if block.list {
			result = append(result, strings.Join(block.lines, "\n"))
		} else if text := unwrapText(trimBlankLines(block.lines)); text != "" {
			result = append(result, text)
		}
	}

	return strings.Join(result, "\n\n")
}

func (*Functions) unwrapTable(content []string) string {
	if !containsList(content) {
		return unwrapText(content)
	}

	blocks := splitCommentBlocks(content)
	var result strings.Builder

	for _, block := range blocks {
		if block.list {
			result.WriteString(renderHTMLList(block.lines))
		} else {
			result.WriteString(unwrapText(trimBlankLines(block.lines)))
		}
	}

	return result.String()
}

func unwrapText(content []string) string {
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

type commentBlock struct {
	lines []string
	list  bool
}

func containsList(content []string) bool {
	for _, line := range content {
		if isListItem(line) {
			return true
		}
	}

	return false
}

func splitCommentBlocks(content []string) []commentBlock {
	var result []commentBlock

	for len(content) > 0 {
		list := isListItem(content[0])
		end := 1

		for end < len(content) {
			if isListItem(content[end]) {
				if list {
					end++

					continue
				}

				break
			}

			if list && isListContinuation(content[end]) {
				end++

				continue
			}

			if !list {
				end++

				continue
			}

			break
		}

		result = append(result, commentBlock{
			lines: content[:end],
			list:  list,
		})
		content = content[end:]
	}

	return result
}

func isListItem(line string) bool {
	return len(line) > 1 &&
		(line[0] == '-' || line[0] == '*') &&
		(line[1] == ' ' || line[1] == '\t')
}

func isListContinuation(line string) bool {
	return len(line) > 0 && (line[0] == ' ' || line[0] == '\t')
}

func trimBlankLines(lines []string) []string {
	for len(lines) > 0 && strings.TrimSpace(lines[0]) == "" {
		lines = lines[1:]
	}

	for len(lines) > 0 && strings.TrimSpace(lines[len(lines)-1]) == "" {
		lines = lines[:len(lines)-1]
	}

	return lines
}

func renderHTMLList(lines []string) string {
	var result strings.Builder
	var item []string

	result.WriteString("<ul>")

	for _, line := range lines {
		if isListItem(line) {
			writeHTMLListItem(&result, item)
			item = []string{strings.TrimSpace(line[1:])}
		} else {
			item = append(item, strings.TrimSpace(line))
		}
	}

	writeHTMLListItem(&result, item)
	result.WriteString("</ul>")

	return result.String()
}

func writeHTMLListItem(result *strings.Builder, lines []string) {
	if len(lines) == 0 {
		return
	}

	result.WriteString("<li>")
	result.WriteString(strings.Join(lines, " "))
	result.WriteString("</li>")
}
