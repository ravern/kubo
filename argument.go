package kubo

// Argument represents an argument for a command.
type Argument struct {
	Name string

	// Multiple is whether this argument collects multiple arguments.
	//
	// This is typically used to support a dynamic number of arguments (e.g.
	// 'command <argument1> <argument2> <arguments...>'). It should only
	// be used at the end of the argument list.
	Multiple bool
}
