package cmd

import (
	"os"
	"github.com/mitchellh/cli"
	"log"
)


func InitCmd() {
	c := cli.NewCLI("app", "1.0.0")
	c.Args = os.Args[1:]
	c.Commands = map[string]cli.CommandFactory{
		"foo": func() (cli.Command, error) {
			return &Command{}, nil
		},
	}

	_, err := c.Run()
	if err != nil {
		log.Println(err)
	}

	//os.Exit(exitStatus)
}
