package typefilter

import (
	"regexp"
	"strings"
)

func createGlobber(glob string) *regexp.Regexp {
	// Convert from wildcards into case insensitive regex
	regex := "(?i)" + regexp.QuoteMeta(glob)
	regex = strings.Replace(regex, "\\*", ".*", -1)

	return regexp.MustCompile(regex)
}
