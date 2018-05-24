package main

import (
	"fmt"

	"github.com/ravernkoh/kubo"
)

var alternating = &kubo.Command{
	Name:    "alternating",
	Aliases: []string{"alt"},
	Arguments: []kubo.Argument{
		{Name: "text"},
	},
	Run: func(ctx *kubo.Context) error {
		fmt.Fprintln(ctx.Stdout(), "Alternating")
		return nil
	},
}
