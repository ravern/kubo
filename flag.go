package kubo

import (
	"fmt"
	"strings"
)

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
	names := []string{fmt.Sprintf("--%s", flag.Name)}
	for _, alias := range flag.Aliases {
		if len(alias) > 1 {
			names = append(names, fmt.Sprintf("--%s", alias))
		} else {
			names = append(names, fmt.Sprintf("-%s", alias))
		}
	}
	return strings.Join(names, ", ")
}
