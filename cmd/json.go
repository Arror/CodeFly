package cmd

import (
	"log"

	"CodeFly/command"
	"CodeFly/lang/swift/generator"
	"CodeFly/reader"

	"github.com/urfave/cli"
)

var info = &command.Command{}

// JSONGenerate json代码生成命令
var JSONGenerate = cli.Command{
	Name:      "json",
	ShortName: "json",
	Usage:     "Command of generate the target language code.",
	Flags: []cli.Flag{
		cli.StringFlag{
			Name:        "lang, l",
			Usage:       "Target language.",
			Destination: &info.Lang,
		},
		cli.StringFlag{
			Name:        "input, i",
			Usage:       "Input thrift file path.",
			Destination: &info.Input,
		},
		cli.StringFlag{
			Name:        "output, o",
			Usage:       "File output path.",
			Destination: &info.Output,
		},
	},
	Action: func(c *cli.Context) error {

		if err := info.CheckCommand(); err != nil {
			log.Fatalln(err.Error())
		}

		ts, err := reader.ReadThrift(info)
		if err != nil {
			log.Fatalln(err.Error())
		}

		generator.Generate(ts, info)

		return nil
	},
}
