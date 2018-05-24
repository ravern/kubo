package kubo

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
