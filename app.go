package kubo

import (
	"fmt"
	"io"
	"os"
)

// App represents a command line app.
type App struct {
	Root *Command // root command

	Stdin  io.Reader // default is os.Stdin
	Stdout io.Writer // default is os.Stdout
}

// NewApp creates a new app with the given root command.
func NewApp(root *Command) *App {
	return &App{
		Root:   root,
		Stdin:  os.Stdin,
		Stdout: os.Stdout,
	}
}

// Run runs the app with the given arguments.
func (a *App) Run(args []string) error {
	args = args[1:]

	cmd := a.Root
	for {
		// Verify command arguments
		for i, arg := range cmd.Arguments {
			if arg.Multiple && i != len(cmd.Arguments)-1 {
				panic(fmt.Errorf("command %s: multiple can only be used in last argument", cmd.Name))
			}
		}

		// Context to be passed to the command
		ctx := Context{
			arguments: make(map[string]string),
			flags:     make(map[string]string),
			stdin:     a.Stdin,
			stdout:    a.Stdout,
		}

		// Parse all the flags in the raw arguments
		for i := 0; i < len(args); i++ {
			arg := args[i]

			name, ok := parseFlagName(arg)
			if ok {
				flag, err := cmd.flag(name)
				if err != nil {
					return err
				}

				var value string
				if flag.Bool {
					value = fmt.Sprint(true)
					args = append(args[:i], args[i+1:]...)
				} else if i+1 < len(args) {
					value = args[i+1]
					args = append(args[:i], args[i+2:]...)
				} else {
					return fmt.Errorf("no value found for flag: %s", name)
				}
				i-- // decrement to keep index correct

				ctx.flags[name] = value
			}
		}

		// Parse as child command
		if len(args) > 0 {
			child, err := cmd.command(args[0])
			if err == nil {
				cmd = child
				continue
			}
		}

		// Parse as arguments
		for _, arg := range cmd.Arguments {
			if len(args) == 0 {
				return fmt.Errorf("argument not found: %s", arg.Name)
			}

			if arg.Multiple {
				ctx.argumentMultipleName = arg.Name
				ctx.argumentMultipleValue = args
				args = nil
				break
			}

			ctx.arguments[arg.Name] = args[0]
			args = args[1:]
		}
		if len(args) > 0 {
			return fmt.Errorf("extra arguments supplied")
		}

		// Run the command
		return cmd.Run(&ctx)
	}
}

// parseFlagName parses the given argument for a flag name, returning the name
// and a flag whether it was found.
func parseFlagName(arg string) (string, bool) {
	matches := longFlagRegexp.FindStringSubmatch(arg)
	if len(matches) > 1 {
		return matches[1], true
	}

	matches = shortFlagRegexp.FindStringSubmatch(arg)
	if len(matches) > 1 {
		return matches[1], true
	}

	return "", false
}
