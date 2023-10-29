package functions

import (
	"fmt"
	"regexp"
	"strings"
)

// referenceDefinitionRegex is a regular expression that matches the definition of a reference link.
// It captures the name and the URL.
var referenceDefinitionRegex = regexp.MustCompile(`^\s*\[(?P<name>[^\]]+)\]:\s*(?P<url>.*)$`)

// linkReferenceRegex is a regular expression that matches use of a reference link.
// It captures the name.
// No match is returned for inline markdown links
var linkReferenceRegex = regexp.MustCompile(`\[(?P<name>[^\]]+)\]`)

// inline links converts reference style links into inline ones.
// Reference style links don't work well within markdown tables, and there's a risk of
// naming conflicts given we generate all content into a single file. Inlining links
// avoids problems.
func (f *Functions) inlineLinks(lines []string) []string {
	defs := make(map[string]string)
	result := make([]string, 0, len(lines))

	// Scan through lines, identifying and removing reference definitions while
	// building up result
	for _, line := range lines {
		matches := referenceDefinitionRegex.FindStringSubmatch(line)
		if matches == nil {
			// This is not a reference definition, so add it to the result
			result = append(result, line)
			continue
		}

		// This is a reference definition, so add it to the map if we don't already
		// have a definition with this name
		linkText := matches[1]
		destination := matches[2]
		key := strings.ToLower(linkText)

		if _, ok := defs[key]; !ok {
			defs[key] = destination
		}
	}

	// Iterate through lines, replacing any reference links with inline links
	for i, line := range result {
		result[i] = linkReferenceRegex.ReplaceAllStringFunc(line, func(match string) string {
			// Extract the name of the link
			name := linkReferenceRegex.FindStringSubmatch(match)[1]
			key := strings.ToLower(name)
			if destination, ok := defs[key]; ok {
				// This is a reference link, so replace it with an inline link
				return fmt.Sprintf("[%s](%s)", name, destination)
			}

			// This is not a reference link, so leave it alone
			return match
		})
	}

	return result
}
