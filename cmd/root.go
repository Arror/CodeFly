package cmd

import (
	"os"

	"github.com/urfave/cli"
)

// Root Root command
var Root = cli.NewApp()

// Execute Excuate root command
func Execute() {
	if err := Root.Run(os.Args); err != nil {
		panic(err.Error())
	}
}

func init() {
	Root.Commands = []cli.Command{
		JSONCommand,
	}
}
