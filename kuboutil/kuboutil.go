package kuboutil

import (
	"fmt"
	"strconv"
)

// Int returns the given string as an int and an error if it can't be converted
// or if an error was given.
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

// Int8 returns the given string as an int8 and an error if it can't be converted
// or if an error was given.
func Int8(v string, err error) (int8, error) {
	if err != nil {
		return 0, err
	}

	i, err := strconv.ParseInt(v, 10, 8)
	if err != nil {
		return 0, fmt.Errorf("could not convert %s to an int", v)
	}

	return int8(i), nil
}

// Int16 returns the given string as an int16 and an error if it can't be converted
// or if an error was given.
func Int16(v string, err error) (int16, error) {
	if err != nil {
		return 0, err
	}

	i, err := strconv.ParseInt(v, 10, 16)
	if err != nil {
		return 0, fmt.Errorf("could not convert %s to an int", v)
	}

	return int16(i), nil
}

// Int32 returns the given string as an int32 and an error if it can't be converted
// or if an error was given.
func Int32(v string, err error) (int32, error) {
	if err != nil {
		return 0, err
	}

	i, err := strconv.ParseInt(v, 10, 32)
	if err != nil {
		return 0, fmt.Errorf("could not convert %s to an int", v)
	}

	return int32(i), nil
}

// Int64 returns the given string as an int64 and an error if it can't be converted
// or if an error was given.
func Int64(v string, err error) (int64, error) {
	if err != nil {
		return 0, err
	}

	i, err := strconv.ParseInt(v, 10, 64)
	if err != nil {
		return 0, fmt.Errorf("could not convert %s to an int", v)
	}

	return int64(i), nil
}

// Uint8 returns the given string as an uint8 and an error if it can't be converted
// or if an error was given.
func Uint8(v string, err error) (uint8, error) {
	if err != nil {
		return 0, err
	}

	i, err := strconv.ParseUint(v, 10, 8)
	if err != nil {
		return 0, fmt.Errorf("could not convert %s to an int", v)
	}

	return uint8(i), nil
}

// Uint16 returns the given string as an uint16 and an error if it can't be converted
// or if an error was given.
func Uint16(v string, err error) (uint16, error) {
	if err != nil {
		return 0, err
	}

	i, err := strconv.ParseUint(v, 10, 16)
	if err != nil {
		return 0, fmt.Errorf("could not convert %s to an int", v)
	}

	return uint16(i), nil
}

// Uint32 returns the given string as an uint32 and an error if it can't be converted
// or if an error was given.
func Uint32(v string, err error) (uint32, error) {
	if err != nil {
		return 0, err
	}

	i, err := strconv.ParseUint(v, 10, 32)
	if err != nil {
		return 0, fmt.Errorf("could not convert %s to an int", v)
	}

	return uint32(i), nil
}

// Uint64 returns the given string as an uint64 and an error if it can't be converted
// or if an error was given.
func Uint64(v string, err error) (uint64, error) {
	if err != nil {
		return 0, err
	}

	i, err := strconv.ParseUint(v, 10, 64)
	if err != nil {
		return 0, fmt.Errorf("could not convert %s to an int", v)
	}

	return i, nil
}

// Float32 returns the given string as a float32 and an error if it can't be
// converted or if an error was given.
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

// Float64 returns the given string as a float64 and an error if it can't be
// converted or if an error was given.
func Float64(v string, err error) (float64, error) {
	if err != nil {
		return 0, err
	}

	f, err := strconv.ParseFloat(v, 64)
	if err != nil {
		return 0, fmt.Errorf("could not convert %s to an float64", v)
	}

	return float64(f), nil
}

// Bool returns the given string as an bool and an error if it can't be converted
// or if an error was given.
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
