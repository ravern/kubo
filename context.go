package kubo

import (
	"fmt"
	"io"
)

// Context represents the runtime context of a command.
type Context struct {
	arguments map[string]string
	flags     map[string]string

	argumentMultipleName  string
	argumentMultipleValue []string

	stdin  io.Reader
	stdout io.Writer
}

// Argument returns the argument with the given name and an error if it doesn't
// exist.
func (ctx *Context) Argument(name string) (string, error) {
	arg, ok := ctx.arguments[name]
	if !ok {
		return "", fmt.Errorf("argument not found: %s", name)
	}
	return arg, nil
}

// ArgumentMultiple returns the argument with the given name as a collected
// argument and an error if it doesn't exist.
func (ctx *Context) ArgumentMultiple(name string) ([]string, error) {
	if ctx.argumentMultipleName != name || ctx.argumentMultipleValue == nil {
		return nil, fmt.Errorf("multiple argument not found: %s", name)
	}
	return ctx.argumentMultipleValue, nil
}

// Flag returns the argument with the given name and an error if it doesn't
// exist.
func (ctx *Context) Flag(name string) (string, error) {
	arg, ok := ctx.flags[name]
	if !ok {
		return "", fmt.Errorf("flag not found: %s", name)
	}
	return arg, nil
}

// Stdin returns the stdin defined in the app.
func (ctx *Context) Stdin() io.Reader {
	return ctx.stdin
}

// Stdout returns the stdout defined in the app.
func (ctx *Context) Stdout() io.Writer {
	return ctx.stdout
}
