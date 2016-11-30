package cmd

import (
	"log"

	"CodeFly/distributor"
	"CodeFly/model"
	"CodeFly/reader"

	"github.com/urfave/cli"
)

var genInfo = &model.GenerateCommandInfo{}

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

		ts, err := reader.ReadThrift(genInfo.Input)
		if err != nil {
			log.Fatalln(err.Error())
		}

		if err := reader.CheckLanguageNameSpace(genInfo.Lang, ts); err != nil {
			log.Fatalln(err.Error())
		}

		distributor.Distribute(ts, genInfo)

		return nil
	},
}
