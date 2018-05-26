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
	cmd := a.Root
	for {
		tmpArgs := args[1:]

		// Verify that multiple is only used once in the arguments
		for i, arg := range cmd.Arguments {
			if arg.Multiple && i != len(cmd.Arguments)-1 {
				panic(fmt.Errorf("command %s: multiple can only be used in last argument", cmd.Name))
			}
		}

		// Create the context to pass to the command
		ctx := Context{
			arguments: make(map[string]string),
			flags:     make(map[string]string),
			stdin:     a.Stdin,
			stdout:    a.Stdout,
		}

		// flagErr is used to hold the flag not found error, which can
		// only be returned if no subcommand is found
		var flagErr error

		// Parse all the flags in the arguments
		for i := 0; i < len(tmpArgs); i++ {
			arg := tmpArgs[i]

			name, ok := parseFlagName(arg)
			if ok {
				// Try to find the flag definition
				flag, err := cmd.flag(name)
				if err != nil {
					// Since it is not found, hold the flag
					// not found error for later and simply
					// let it parse as per normal
					flagErr = err
					flag.Bool = true
				}

				var value string
				if flag.Bool {
					value = fmt.Sprint(true)
					tmpArgs = append(append([]string{}, tmpArgs[:i]...), tmpArgs[i+1:]...)
				} else if i+1 < len(tmpArgs) {
					value = tmpArgs[i+1]
					tmpArgs = append(append([]string{}, tmpArgs[:i]...), tmpArgs[i+2:]...)
				} else {
					return fmt.Errorf("no value found for flag: %s", name)
				}
				i-- // decrement to keep index correct

				// Don't set the flag in the context since it
				// was not defined in the command
				if flagErr == nil {
					ctx.flags[flag.Name] = value
				}
			}
		}

		// Set all flags with Bool to false if not set to true
		for _, flag := range cmd.Flags {
			if !flag.Bool {
				continue
			}

			// Since flag is a bool flag and it is not set to true,
			// set it to false
			if _, err := ctx.Flag(flag.Name); err != nil {
				ctx.flags[flag.Name] = "false"
			}
		}

		// Parse raw arguments as child command
		if len(tmpArgs) > 0 {
			// Try to find child command
			child, err := cmd.command(tmpArgs[0])
			if err != nil {
				// If no child command is found and it not possibly
				// an argument, then return the command not found
				// error
				if len(cmd.Arguments) == 0 {
					return err
				}
			} else {
				// Set the command to the child and pop the first
				// argument that matches the child command name
				cmd = child
				for i, arg := range args {
					if arg == tmpArgs[0] {
						args = append(append([]string{}, args[:i]...), args[i+1:]...)
						break
					}
				}
				continue
			}
		}

		// Since no subcommand is found, the flag not found error should
		// be returned if it is not nil
		if flagErr != nil {
			return flagErr
		}

		// Parse raw arguments as arguments
		for _, arg := range cmd.Arguments {
			if len(tmpArgs) == 0 {
				return fmt.Errorf("argument not found: %s", arg.Name)
			}

			if arg.Multiple {
				ctx.argumentMultipleName = arg.Name
				ctx.argumentMultipleValue = tmpArgs
				tmpArgs = nil
				break
			}

			ctx.arguments[arg.Name] = tmpArgs[0]
			tmpArgs = tmpArgs[1:]
		}
		if len(tmpArgs) > 0 {
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
