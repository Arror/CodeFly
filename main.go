package main

import (
	"os"

	"CodeFly/cmd"
	"CodeFly/info"

	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "CodeFly"
	app.Usage = "iOS开发者工具集"
	app.Version = info.Version
	app.Author = info.Author
	app.Email = info.Email
	app.Commands = []cli.Command{
		cmd.Gen,
	}
	app.Run(os.Args)
}
