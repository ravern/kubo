[![GoDoc](https://img.shields.io/badge/godoc-reference-blue.svg)](https://godoc.org/github.com/ravernkoh/kubo)

# Kubo
Lightweight package to write command line apps in Go.

## Aims
To be (and remain) as **lightweight** as possible, while still providing
sufficient features to build rich command line applications. Also to be easy to
**understand** and **use**.

## Installation
For now, use `go get` (will add support for `vgo` in a later update).
```bash
$ go get github.com/ravernkoh/kubo
```

## Usage

### Basic
The most basic app has just one command, with no arguments and no flags.

```go
app := kubo.NewApp(&kubo.Command{
    Name: "basic",
    Description: "a basic hello world app",
    Run: func(ctx *kubo.Context) error {
        // Prints 'hello, world!'
        fmt.Fprintln(ctx.Stdout(), "hello, world!")
    },
})
```

`Run` then runs the app and returns an error which should be handled (usually
by simply printing it out).

```go
// Blocks until the command is completed
if err := app.Run(); err != nil {
    fmt.Printf("error: %v\n", err)
}
```

In this case, the app simply prints `"hello, world!"` (more will be explained on
the context later).

```bash
$ basic
hello, world!
```

### Commands
The building block of a command line app is a command. Flags, arguments and child
commands can all be defined on a command.

```go
kubo.Command{
    Name: "commands",
    Description: "a random command",
    Run: func(ctx *kubo.Context) error {
        // Prints 'random'
        fmt.Fprintln(ctx.Stdout(), "random")
    },
}
```

The `Run` function is the function that will be called if the raw arguments are
successfully parsed by the app. Usually, code will be written in this function.

### Flags
Defining flags on a command is easy.

```go
kubo.Command{
    Name: "flags",
    Description: "a command with flags",
    // Defines the flags for this command
    Flags: []kubo.Flag{
        {
            Name: "one",
            Description: "the first flag",
        },
        {
            Name: "two",
            Description: "the second flag",
        },
    },
}
```

The code above defines two flags, called `one` and `two`, which will be available
for use with the command.

```bash
$ flags --one value1 --two value2
```

Flags can have aliases, which defines alternate names for them.

```go
kubo.Flag{
    Name: "one",
    Description: "the first flag",
    Aliases: []string{"o"},
}
```

For single letter flags, only a single dash needs to be used.

```bash
$ flags -o value1 --two value2
```

Flags also have a field called `Bool`. If this is set to true, then no value
needs to be passed to them.

```go
kubo.Flag{
    Name: "two",
    Description: "the second flag",
    Bool: true,
}
```

The resulting value would be `"true"` if the flag is set and `"false"` if the
flag is not set.

*Note that once `Bool` is set, no value* should *be passed to the flag, as the
parser will not try to parse for the flag value.*

```bash
$ flags -o value1 --two
```

### Arguments
Defining arguments on a command is also easy.

```go
kubo.Command{
    Name: "arguments",
    Description: "a command with arguments",
    // Defines the arguments for this command
    Arguments: []kubo.Argument{
        {
            Name: "one",
        },
        {
            Name: "two",
        },
    },
}
```

The code above defines two arguments, `one` and `two`. The order in which you
define the arguments matters. This is because arguments are parsed by their
positions and not by their names.

```bash
$ arguments value1 value2
```

This will result in `one` having the value `"value1"` and `two` having the value
`"value2"`.

Arguments can also have a field of `Multiple`, which causes the argument to
collect multiple values.

```go
kubo.Argument{
    Name: "two",
    Multiple: true,
}
```

*Note that only the last argument can have `Multiple` set.*

```bash
$ arguments value1 value2 value3 value4
```

This will result in `two` having the value `["value2", "value3", "value4"]`.

### Contexts
The context passed in the run function is used to get the arguments and flags
that were parsed from the raw arguments.

```go
kubo.Command{
    Name: "contexts",
    Description: "a command with arguments and flags",
    Arguments: []kubo.Argument{
        {Name: "argument"},
    },
    Flags: []kubo.Flag{
        {Name: "flag"},
    },
    Run: func(ctx *kubo.Context) error {
        // Gets the argument called 'argument'
        argument, err := ctx.Argument("argument")
        if err != nil {
            return err
        }

        // Gets the flag called 'flag'
        flag, err := ctx.Flag("flag")
        if err != nil {
            return err
        }

        fmt.Fprintf(ctx.Stdout(), "argument: %s, flag: %s\n", argument, flag)
    },
}
```

The context also contains the methods for `Stdin` and `Stdout`, which *should* be
used to read from and write to the console. They can be configured in the app
itself (which will pass these values to the context).

```go
// Default values
app.Stdin = os.Stdin
app.Stdout = os.Stdout
```

When getting arguments and flags from the context, sometimes their values need
to be converted to other types. For that purpose, the `kuboutil` package can be
used.

```go
// Gets the argument called 'argument' and converts it to an int
argument, err := kuboutil.Int(ctx.Argument("argument"))
if err != nil {
    return err
}
```

These conversion utilities automatically propogate the error from the `Argument`
method.

### Child commands
Commands can have child commands.

```go
parent := &kubo.Command{
    Name: "parent",
}

child := &kubo.Command{
    Name: "child",
}

// Makes 'child' a child of the 'parent' command
parent.Add(child)
```

These child commands can be called by passing in their name.

```bash
$ parent child
```

Child commands can have flags, arguments, and even child commands of their own!

```go
parent := &kubo.Command{
    Name: "parent",
}

child := &kubo.Command{
    Name: "child",
}

grandchild := &kubo.Command{
    Name: "grandchild",
}

// Makes 'grandchild' a child of the 'child' command
child.Add(grandchild)

// Makes 'child' a child of the 'parent' command
parent.Add(child)
```

They can then be called by passing in their names.

```bash
$ parent child grandchild
```

### Help command
A help command can be generated for each command.

```go
complex := &kubo.Command{
    Name: "complex",
    Description: "some complex command",
}

// Makes 'help' a child command of the 'complex' command
complex.Add(complex.Help())
```

The help command can be called using `help`.

```bash
$ complex help
```

## Examples
More examples can be found in the `_examples` folder.

## Alternatives
Of course, there are other more notable packages for building command line apps
(Forgive me if I missed yours out).

+ [github.com/spf13/cobra](https://github.com/spf13/cobra)
+ [github.com/urfave/cli](https://github.com/urfave/cli)

Kubo was built as some practice for me and also to contribute to open source
software!

## License
This project is licensed under the GNU Public License 3.0.

## Author
+ Ravern Koh ([ravernkoh@gmail.com](mailto://ravernkoh@gmail.com))