package main

import (
	"fmt"
	"os"

	"github.com/ravernkoh/kubo"
)

var root = &kubo.Command{
	Name:        "children",
	Description: "prints out the given arguments",
	Arguments: []kubo.Argument{
		{
			Name:     "arguments",
			Multiple: true,
		},
	},
	Run: func(ctx *kubo.Context) error {
		arguments, err := ctx.ArgumentMultiple("arguments")
		if err != nil {
			return err
		}

		fmt.Fprint(ctx.Stdout(), "arguments:")
		for _, argument := range arguments {
			fmt.Fprintf(ctx.Stdout(), " %s", argument)
		}
		fmt.Fprintln(ctx.Stdout())

		return nil
	},
}

func main() {
	first.Add(first.Help())

	second.Add(second.Help())

	root.Add(root.Help())
	root.Add(first)
	root.Add(second)

	if err := kubo.NewApp(root).Run(os.Args); err != nil {
		fmt.Printf("error: %s\n", err)
	}
}
