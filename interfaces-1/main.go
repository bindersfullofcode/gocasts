package main

import (
	"github.com/bfoc/gocasts/interfaces-1/printer"
	"github.com/bfoc/gocasts/interfaces-1/user"
	"log"
)

func main() {
	client := user.New()
	formatter := &printer.Printer{client}

	if err := formatter.Print(); err != nil {
		log.Fatal(err)
	}
}