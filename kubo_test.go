package kubo_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/ravernkoh/kubo"
)

var root = &kubo.Command{
	Name:        "root",
	Description: "prints root",
	Run: func(ctx *kubo.Context) error {
		fmt.Fprintln(ctx.Stdout(), "root")
		return nil
	},
}

var flags = &kubo.Command{
	Name:        "flags",
	Description: "prints flags",
	Flags: []kubo.Flag{
		{
			Name:        "normal",
			Description: "normal flag",
		},
		{
			Name:        "alias",
			Aliases:     []string{"a"},
			Description: "alias flag",
		},
		{
			Name:        "bool",
			Description: "bool flag",
			Bool:        true,
		},
	},
	Run: func(ctx *kubo.Context) error {
		n, err := ctx.Flag("normal")
		if err != nil {
			return err
		}

		a, err := ctx.Flag("alias")
		if err != nil {
			return err
		}

		b, err := ctx.Flag("bool")
		if err != nil {
			return err
		}

		fmt.Fprintf(ctx.Stdout(), "flag normal: %s\n", n)
		fmt.Fprintf(ctx.Stdout(), "flag alias: %s\n", a)
		fmt.Fprintf(ctx.Stdout(), "flag bool: %v\n", b)

		return nil
	},
}

var arguments = &kubo.Command{
	Name:        "arguments",
	Description: "prints arguments",
	Arguments: []kubo.Argument{
		{
			Name: "normal",
		},
		{
			Name:     "multiple",
			Multiple: true,
		},
	},
	Run: func(ctx *kubo.Context) error {
		n, err := ctx.Argument("normal")
		if err != nil {
			return err
		}

		m, err := ctx.ArgumentMultiple("multiple")
		if err != nil {
			return err
		}

		fmt.Fprintf(ctx.Stdout(), "argument normal: %s\n", n)
		fmt.Fprintf(ctx.Stdout(), "argument multiple: %s\n", strings.Join(m, ", "))

		return nil
	},
}

func init() {
	flags.Add(flags.Help())

	arguments.Add(arguments.Help())

	root.Add(root.Help())
	root.Add(flags)
	root.Add(arguments)
}

func TestApp(t *testing.T) {
	tests := []struct {
		cmd  string
		want string
	}{
		{
			cmd:  "root",
			want: "root\n",
		},
		{
			cmd:  "root flags --normal hello -a world",
			want: "flag normal: hello\nflag alias: world\nflag bool: false\n",
		},
		{
			cmd:  "root arguments hello world again",
			want: "argument normal: hello\nargument multiple: world, again\n",
		},
	}

	for i, test := range tests {
		var b strings.Builder

		app := kubo.NewApp(root)
		app.Stdout = &b
		if err := app.Run(strings.Fields(test.cmd)); err != nil {
			fmt.Fprintf(&b, "error: %v", err)
		}

		got := b.String()
		if test.want != got {
			t.Errorf("test %d: want %v, got %v", i, test.want, got)
		}
	}
}
