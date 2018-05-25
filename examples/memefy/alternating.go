package main

import (
	"fmt"
	"strings"

	"github.com/ravernkoh/kubo"
)

var alternating = &kubo.Command{
	Name:        "alternating",
	Description: "alternates between upper and lower case",
	Aliases:     []string{"alt"},
	Arguments: []kubo.Argument{
		{Name: "text"},
	},
	Run: func(ctx *kubo.Context) error {
		text, err := ctx.Argument("text")
		if err != nil {
			return err
		}

		var textRunes []rune
		for i, r := range text {
			if i%2 == 0 {
				textRunes = append(textRunes, []rune(strings.ToUpper(string(r)))...)
			} else {
				textRunes = append(textRunes, []rune(strings.ToLower(string(r)))...)
			}
		}

		fmt.Fprintln(ctx.Stdout(), string(textRunes))

		return nil
	},
}
