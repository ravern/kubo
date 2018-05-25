package main

import (
	"fmt"

	"github.com/ravernkoh/kubo"
)

var second = &kubo.Command{
	Name:        "second",
	Aliases:     []string{"s"},
	Description: "throws an error",
	Run: func(ctx *kubo.Context) error {
		return fmt.Errorf("something went wrong")
	},
}
