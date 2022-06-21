package command

import (
	"fmt"
)

type PingCommand struct {
	All bool `short:"a" long:"all" description:"Add all files"`
}

var pingCommand PingCommand

func (c *PingCommand) Execute(args []string) error {
	fmt.Printf("Adding (all=%v): %#v - %#v\n", c.All, args, c)
	return nil
}

func init() {
	//parser.AddCommand("ping",
	//	"Add a file",
	//	"The add command adds a file to the repository. Use -a to add all files.",
	//	&pingCommand)
}
