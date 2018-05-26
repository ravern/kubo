[![GoDoc](https://img.shields.io/badge/godoc-reference-blue.svg)](https://godoc.org/github.com/ravernkoh/kubo)

# Kubo
Lightweight package to write command line apps in Go.

## Aims
To be (and remain) as **lightweight** as possible, while still providing
sufficient features to build rich command line applications. Also to be easy to
**understand** and **use**.

## Installation
For now, use `go get` (will move to `vgo` in a later commit).
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
})
```

`Run` then runs the app and returns an error which should be handled 
(usually by simply printing it out).

```go
if err := app.Run(); err != nil {
    fmt.Printf("error: %v\n", err)
}
```

### Flags
Defining flags on a command is easy.

```go
kubo.Command{
    Name: "flags",
    Description: "a command with flags",
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

### Child commands
Commands can have child commands.

```go
parent := &kubo.Command{
    Name: "parent",
}

child := &kubo.Command{
    Name: "child",
}

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

child.Add(grandchild)
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

complex.Add(complex.Help())
```

The help command can be called using `help`.

```bash
$ complex help
```

## Examples
More examples can be found in the `_examples` folder.
