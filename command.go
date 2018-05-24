package kubo

import "fmt"

// Command represents a command or a subcommand.
type Command struct {
	Name        string
	Aliases     []string
	Description string

	Arguments []Argument // should be in order
	Flags     []Flag

	// Run runs the command.
	//
	// Any error returned is propogated and returned to the main Run function
	// of the app.
	//
	// When reading from and printing to the console, Stdin and Stdout from
	// the context is used.
	Run func(*Context) error

	// Used for generating help command.
	parent   *Command
	children []*Command
}

// AddCommand adds a child command.
func (cmd *Command) AddCommand(child *Command) {
	child.parent = cmd
	cmd.children = append(cmd.children, child)
}

// command returns the child command with the given name or alias.
func (cmd *Command) command(nameOrAlias string) (*Command, error) {
	for _, child := range cmd.children {
		if child.Name == nameOrAlias {
			return child, nil
		}
		for _, alias := range child.Aliases {
			if alias == nameOrAlias {
				return child, nil
			}
		}
	}
	return nil, fmt.Errorf("command not found: %s", nameOrAlias)
}

// flag returns the flag with the given name or alias.
func (cmd *Command) flag(nameOrAlias string) (Flag, error) {
	for _, flag := range cmd.Flags {
		if flag.Name == nameOrAlias {
			return flag, nil
		}
		for _, alias := range flag.Aliases {
			if alias == nameOrAlias {
				return flag, nil
			}
		}
	}
	return Flag{}, fmt.Errorf("flag not found: %s", nameOrAlias)
}
