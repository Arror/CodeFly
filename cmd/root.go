package cmd

import (
	"os"

	"github.com/urfave/cli"
)

var root = cli.NewApp()

// Execute Excuate root command
func Execute() {
	if err := root.Run(os.Args); err != nil {
		panic(err.Error())
	}
}

func init() {
	root.Commands = []cli.Command{
		jsonCommand,
	}
}
