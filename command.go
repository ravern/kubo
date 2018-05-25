package kubo

import (
	"fmt"
	"strings"
)

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

// Add adds a child command.
func (cmd *Command) Add(child *Command) {
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
	return nil, fmt.Errorf("command not defined: %s", nameOrAlias)
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
	return Flag{}, fmt.Errorf("flag not defined: %s", nameOrAlias)
}

// Help returns a generated help command which prints usage details on run.
func (cmd *Command) Help() *Command {
	return &Command{
		Name:        "help",
		Aliases:     []string{"h"},
		Description: "prints description and usage details",
		Run: func(ctx *Context) error {
			fmt.Fprintln(ctx.Stdout(), cmd.Usage())
			return nil
		},
	}
}

// Usage returns the usage details.
func (cmd *Command) Usage() string {
	// Find the maximum number of tabs
	var maxLen int
	for _, flag := range cmd.Flags {
		nameAndAliases := flag.nameAndAliases()
		if len(nameAndAliases) > maxLen {
			maxLen = len(nameAndAliases)
		}
	}
	for _, child := range cmd.children {
		nameAndAliases := child.nameAndAliases()
		if len(nameAndAliases) > maxLen {
			maxLen = len(nameAndAliases)
		}
	}
	maxTabs := maxLen/TabSize + 1

	// usage is where the usage string will be built up
	var usage strings.Builder

	// Name and description
	usage.WriteString(fmt.Sprintln("name"))
	usage.WriteString(fmt.Sprintf("\t%s - %s", cmd.fullName(), cmd.Description))

	// Command usage
	usage.WriteString("\n\n")
	usage.WriteString(fmt.Sprintln("usage"))
	commandUsages := cmd.commandUsages()
	for i, commandUsage := range commandUsages {
		usage.WriteString(fmt.Sprintf("\t%s %s", cmd.fullName(), commandUsage))
		if i != len(commandUsages)-1 {
			usage.WriteString("\n")
		}
	}

	// Flags
	if len(cmd.Flags) > 0 {
		usage.WriteString("\n\n")
		usage.WriteString(fmt.Sprintln("flags"))
		for i, flag := range cmd.Flags {
			nameAndAliases := flag.nameAndAliases()
			usage.WriteString(fmt.Sprintf(
				"\t%s%s%s",
				nameAndAliases,
				tabs(maxTabs-len(nameAndAliases)/TabSize),
				flag.Description,
			))
			if i != len(cmd.Flags)-1 {
				usage.WriteString("\n")
			}
		}
	}

	// Commands
	if len(cmd.children) > 0 {
		usage.WriteString("\n\n")
		usage.WriteString(fmt.Sprintln("commands"))
		for i, child := range cmd.children {
			nameAndAliases := child.nameAndAliases()
			usage.WriteString(fmt.Sprintf(
				"\t%s%s%s",
				nameAndAliases,
				tabs(maxTabs-len(nameAndAliases)/TabSize),
				child.Description,
			))
			if i != len(cmd.children)-1 {
				usage.WriteString("\n")
			}
		}
	}

	return usage.String()
}

// fullName returns the full name of the command (including parent names).
func (cmd *Command) fullName() string {
	name := cmd.Name
	cmd = cmd.parent
	for cmd != nil {
		name = fmt.Sprintf("%s %s", cmd.Name, name)
		cmd = cmd.parent
	}
	return name
}

// commandUsages returns the specific command usage possibilities.
func (cmd *Command) commandUsages() []string {
	var usages []string
	if len(cmd.Arguments) > 0 {
		usage := fmt.Sprintf("<%s>", cmd.Arguments[0].Name)
		for _, arg := range cmd.Arguments[1:] {
			if arg.Multiple {
				usage = fmt.Sprintf("%s <%s>...", usage, arg.Name)
			} else {
				usage = fmt.Sprintf("%s <%s>", usage, arg.Name)
			}
		}
		usages = append(usages, usage)
	}
	if len(cmd.children) > 0 {
		usages = append(usages, "<command>")
	}

	return usages
}

func (cmd *Command) nameAndAliases() string {
	names := []string{cmd.Name}
	for _, alias := range cmd.Aliases {
		names = append(names, alias)
	}
	return strings.Join(names, ", ")
}
