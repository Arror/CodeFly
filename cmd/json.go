package cmd

import (
	"log"

	"CodeFly/command"
	"CodeFly/printer"
	"CodeFly/reader"

	"github.com/urfave/cli"
)

var genInfo = &command.Command{}

// JSONGenerate json代码生成命令
var JSONGenerate = cli.Command{
	Name:      "json",
	ShortName: "json",
	Usage:     "Command of generate the target language code.",
	Flags: []cli.Flag{
		cli.StringFlag{
			Name:        "lang, l",
			Usage:       "Target language.",
			Destination: &genInfo.Lang,
		},
		cli.StringFlag{
			Name:        "input, i",
			Usage:       "Input thrift file path.",
			Destination: &genInfo.Input,
		},
		cli.StringFlag{
			Name:        "output, o",
			Usage:       "File output path.",
			Destination: &genInfo.Output,
		},
	},
	Action: func(c *cli.Context) error {

		if err := genInfo.CheckGenerateCommandInfo(); err != nil {
			log.Fatalln(err.Error())
		}

		ts, err := reader.ReadThrift(genInfo)
		if err != nil {
			log.Fatalln(err.Error())
		}

		if err := reader.CheckLanguageNameSpace(ts, genInfo); err != nil {
			log.Fatalln(err.Error())
		}

		printer.Generate(ts, genInfo)

		return nil
	},
}
