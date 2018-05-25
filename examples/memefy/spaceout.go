package main

import (
	"fmt"

	"github.com/ravernkoh/kubo"
	"github.com/ravernkoh/kubo/kuboutil"
)

var spaceout = &kubo.Command{
	Name:        "spaceout",
	Aliases:     []string{"spc"},
	Description: "inserts spaces between characters",
	Arguments: []kubo.Argument{
		{Name: "text"},
	},
	Flags: []kubo.Flag{
		{
			Name:        "space",
			Aliases:     []string{"s"},
			Description: "nunber of spaces between each character",
		},
	},
	Run: func(ctx *kubo.Context) error {
		text, err := ctx.Argument("text")
		if err != nil {
			return err
		}

		space, err := kuboutil.Int(ctx.Flag("space"))
		if err != nil {
			space = 1
		}

		var textRunes []rune
		for _, r := range text {
			textRunes = append(textRunes, r)
			for i := 0; i < space; i++ {
				textRunes = append(textRunes, ' ')
			}
		}

		fmt.Fprintln(ctx.Stdout(), string(textRunes))

		return nil
	},
}
