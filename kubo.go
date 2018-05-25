package kubo

import (
	"regexp"
	"strings"
)

// TabSize is the number of spaces a tab occupies.
var TabSize = 8

// tabs returns the given number of tabs as a string.
func tabs(n int) string {
	var tabs strings.Builder
	for i := 0; i < n; i++ {
		tabs.WriteRune('\t')
	}
	return tabs.String()
}

// Regexps for parsing flags
var (
	longFlagRegexp  = regexp.MustCompile("^--([a-zA-Z0-9][a-zA-Z0-9\\-_]*[a-zA-Z0-9])$")
	shortFlagRegexp = regexp.MustCompile("^-([a-zA-Z])$")
)
