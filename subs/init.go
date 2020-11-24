package subs

import (
	"log"

	"github.com/kwarc/water/cmd"
	"github.com/spf13/pflag"
)

// Init is the 'water init' command
var Init cmd.Command = initC{}

type initC struct{}

func (initC) Name() string { return "init" }

func (initC) Init(flagset *pflag.FlagSet) {}

func (initC) Main(l *log.Logger, argv []string) error {
	// TODO: Initialize command stucture
	l.Println("Nothing went wrong")
	return nil
}
