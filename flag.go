package kubo

import "strings"

// Flag represents a flag for a command.
type Flag struct {
	Name        string
	Aliases     []string
	Description string

	// Bool is whether this flag has a value or not.
	//
	// If this is set, this flag will be used as a boolean flag (e.g.
	// 'command --flag'), which means it does not need a value after it.
	Bool bool
}

// nameAndAliases returns the name and aliases as a comma seperated string
func (flag *Flag) nameAndAliases() string {
	names := []string{flag.Name}
	for _, alias := range flag.Aliases {
		names = append(names, alias)
	}
	return strings.Join(names, ", ")
}
