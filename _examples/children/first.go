package main

import (
	"fmt"

	"github.com/ravernkoh/kubo"
	"github.com/ravernkoh/kubo/kuboutil"
)

var first = &kubo.Command{
	Name:        "first",
	Aliases:     []string{"f"},
	Description: "prints out the given flags",
	Flags: []kubo.Flag{
		{
			Name:        "one",
			Description: "flag number one (optional)",
		},
		{
			Name:        "two",
			Description: "flag number two",
		},
		{
			Name:        "three",
			Description: "flag number three",
			Bool:        true,
		},
	},
	Run: func(ctx *kubo.Context) error {
		one, oneErr := kuboutil.Int(ctx.Flag("one"))

		two, err := kuboutil.Float32(ctx.Flag("two"))
		if err != nil {
			return err
		}

		three, err := kuboutil.Bool(ctx.Flag("three"))
		if err != nil {
			return err
		}

		if oneErr != nil {
			fmt.Fprintln(ctx.Stdout(), "flag one: not provided")
		} else {
			fmt.Fprintf(ctx.Stdout(), "flag one: %d\n", one)
		}
		fmt.Fprintf(ctx.Stdout(), "flag two: %f\n", two)
		fmt.Fprintf(ctx.Stdout(), "flag three: %v\n", three)

		return nil
	},
}
