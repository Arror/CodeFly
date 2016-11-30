package cmd

import (
	"os"

	"github.com/urfave/cli"
)

// Root Root命令
var Root = cli.NewApp()

// Execute 执行Root命令
func Execute() {
	if err := Root.Run(os.Args); err != nil {
		panic(err.Error())
	}
}

func init() {
	Root.Commands = []cli.Command{
		Gen,
	}
}
