package cmd

import (
	//"github.com/mitchellh/cli"
)

type Command struct {
	
}

func (c *Command) Run(args []string) int{

	return 0
}

func (c *Command) Help() string {
	return "this is help"
}

func (c *Command) Synopsis() string {
	return "ppppp"
}



