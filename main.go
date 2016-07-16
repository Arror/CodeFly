package main

import (
	"os"

	"CodeFly/cmd"

	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "CodeFly"
	app.Usage = "Tools for iOS developer"
	app.Commands = []cli.Command{
		cmd.Gen,
	}
	app.Run(os.Args)
}
