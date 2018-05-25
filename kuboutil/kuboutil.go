package kuboutil

import (
	"fmt"
	"strconv"
)

// Int returns the given error if it is not nil, or the given string as an int
// and an error if it can't be converted.
func Int(v string, err error) (int, error) {
	if err != nil {
		return 0, err
	}

	i, err := strconv.Atoi(v)
	if err != nil {
		return 0, fmt.Errorf("could not convert %s to an int", v)
	}

	return i, nil
}

// Float32 returns the given error if it is not nil, or the given string as a
// float32 and an error if it can't be converted.
func Float32(v string, err error) (float32, error) {
	if err != nil {
		return 0, err
	}

	f, err := strconv.ParseFloat(v, 32)
	if err != nil {
		return 0, fmt.Errorf("could not convert %s to an float32", v)
	}

	return float32(f), nil
}

// Bool returns the given error if it is not nil, or the given string as a bool
// and an error if it can't be converted.
func Bool(v string, err error) (bool, error) {
	if err != nil {
		return false, err
	}

	b, err := strconv.ParseBool(v)
	if err != nil {
		return false, fmt.Errorf("could not convert %s to an bool", v)
	}

	return b, nil
}
