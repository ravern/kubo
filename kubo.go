package kubo

import (
	"fmt"
	"regexp"
	"strconv"
)

// TabSize is the number of spaces a tab occupies.
var TabSize = 8

// Regexps for parsing flags
var (
	longFlagRegexp  = regexp.MustCompile("^--([a-zA-Z0-9][a-zA-Z0-9\\-_]*[a-zA-Z0-9])$")
	shortFlagRegexp = regexp.MustCompile("^-([a-zA-Z])$")
)

// Int returns the given error if it is not nil, or the string as an int and an
// error if it can't be converted.
func Int(v string, err error) (int, error) {
	if err != nil {
		return 0, err
	}

	i, err := strconv.Atoi(v)
	if err != nil {
		return 0, fmt.Errorf("could not convert %s to an int: %v", v, err)
	}

	return i, nil
}

// Bool returns the given error if it is not nil, or the string as a bool and an
// error if it can't be converted.
func Bool(v string, err error) (bool, error) {
	if err != nil {
		return false, err
	}

	if v == "true" {
		return true, nil
	} else if v == "false" {
		return false, nil
	}

	return false, fmt.Errorf("could not convert %s to a bool", v)
}
