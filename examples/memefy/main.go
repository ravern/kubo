package main

import (
	"fmt"
	"os"

	"github.com/ravernkoh/kubo"
)

var root = &kubo.Command{
	Name: "memefy",
	Run: func(ctx *kubo.Context) error {
		fmt.Fprintln(ctx.Stdout(), "Root")
		return nil
	},
}

func main() {
	root.Add(alternating)
	root.Add(spaceout)

	if err := kubo.NewApp(root).Run(os.Args); err != nil {
		fmt.Printf("error: %v\n", err)
	}
}
