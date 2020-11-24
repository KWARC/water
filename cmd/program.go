// Package cmd contains command infrastructure code for the 'water' command
package cmd

import (
	"fmt"
	"log"

	"github.com/pkg/errors"
	"github.com/spf13/pflag"
)

// Program represents a main program
type Program struct {
	Name     string
	Commands []Command
}

// Register registers a command in this program
func (p *Program) Register(command Command) {
	p.Commands = append(p.Commands, command)
}

// findCommand finds a command by name or returns nil
func (p Program) findCommand(name string) Command {
	for _, c := range p.Commands {
		if c.Name() == name {
			return c
		}
	}
	return nil
}

// Main executes this program and returns with an appropriate error code.
func (p Program) Main(l *log.Logger, argv []string) {

	// there must be at least one argument
	if len(argv) < 1 {
		l.Fatal(errors.New("Need at least one argument")) // TODO: Proper help messages
	}

	// find the command
	cmdName := argv[0]
	cmd := p.findCommand(cmdName)
	if cmd == nil {
		l.Fatal(errors.Errorf("Unknown subcommand %q", cmdName))
	}

	// create a new flagset and parse the arguments
	flagset := pflag.NewFlagSet(fmt.Sprintf("%s %s", p.Name, cmdName), pflag.ContinueOnError)
	cmd.Init(flagset)
	if err := flagset.Parse(argv[1:]); err != nil {
		l.Fatal(errors.Wrap(err, "Unable to parse arguments"))
	}

	// run the command
	if err := cmd.Main(l, flagset.Args()); err != nil {
		l.Fatal(err)
	}
}

// Command represents a single subcommand of the 'water' program
type Command interface {
	// Name returns the name of the command
	Name() string

	// Init initializes this command on the provided flagset
	Init(flagset *pflag.FlagSet)

	// Main executes this command after flagset has been parsed.
	// argv are the remaining arguments
	Main(l *log.Logger, argv []string) error
}
