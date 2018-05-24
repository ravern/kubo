package kubo

import "regexp"

// TabSize is the number of spaces a tab occupies.
var TabSize = 8

// Regexps for parsing flags
var (
	longFlagRegexp  = regexp.MustCompile("^--([a-zA-Z0-9][a-zA-Z0-9\\-_]*[a-zA-Z0-9])$")
	shortFlagRegexp = regexp.MustCompile("^-([a-zA-Z])$")
)
