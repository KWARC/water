package main

import (
	"log"
	"os"

	"github.com/kwarc/water/cmd"
	"github.com/kwarc/water/subs"
)

func main() {
	// initialize the program
	program := cmd.Program{
		Name: "water",
	}
	program.Register(subs.Init)

	// create a new logger and run
	logger := log.New(os.Stderr, "", log.LstdFlags)
	program.Main(logger, os.Args[1:])
}
