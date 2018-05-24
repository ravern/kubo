package main

import (
	"fmt"

	"github.com/ravernkoh/kubo"
)

var spaceout = &kubo.Command{
	Name:    "spaceout",
	Aliases: []string{"spc"},
	Arguments: []kubo.Argument{
		{Name: "text"},
	},
	Flags: []kubo.Flag{
		{Name: "space"},
	},
	Run: func(ctx *kubo.Context) error {
		fmt.Fprintln(ctx.Stdout(), "Spaceout")
		return nil
	},
}
