package main

import (
	"fmt"
	"os"

	"github.com/ravernkoh/kubo"
)

var root = &kubo.Command{
	Name:        "memefy",
	Description: "automating memeing of text",
	Run: func(ctx *kubo.Context) error {
		fmt.Fprintln(ctx.Stdout(), "Root")
		return nil
	},
}

func main() {
	alternating.Add(alternating.Help())

	spaceout.Add(spaceout.Help())

	root.Add(alternating)
	root.Add(spaceout)
	root.Add(root.Help())

	if err := kubo.NewApp(root).Run(os.Args); err != nil {
		fmt.Printf("error: %v\n", err)
	}
}
