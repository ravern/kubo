package kubo

import (
	"io"
)

// Context represents the runtime context of a command.
type Context struct {
	arguments map[string]string
	multiple  []string // multiple arguments value
	flags     map[string]string

	stdin  io.Reader
	stdout io.Writer
}

// Stdin returns the stdin defined in the app.
func (ctx *Context) Stdin() io.Reader {
	return ctx.stdin
}

// Stdout returns the stdout defined in the app.
func (ctx *Context) Stdout() io.Writer {
	return ctx.stdout
}
