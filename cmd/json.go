package cmd

import (
	"log"

	"CodeFly/generator"
	"CodeFly/parameter"
	"CodeFly/reader"

	"github.com/urfave/cli"
)

var param = &parameter.Parameter{}

// JSONGenerate json代码生成命令
var JSONGenerate = cli.Command{
	Name:      "json",
	ShortName: "json",
	Usage:     "Command of generate the target language code.",
	Flags: []cli.Flag{
		cli.StringFlag{
			Name:        "lang, l",
			Usage:       "Target language.",
			Destination: &param.Lang,
		},
		cli.StringFlag{
			Name:        "input, i",
			Usage:       "Input thrift file path.",
			Destination: &param.Input,
		},
		cli.StringFlag{
			Name:        "output, o",
			Usage:       "File output path.",
			Destination: &param.Output,
		},
	},
	Action: func(c *cli.Context) error {

		if err := param.CheckParameter(); err != nil {
			log.Fatalln(err.Error())
		}

		ts, err := reader.ReadThrift(param)
		if err != nil {
			log.Fatalln(err.Error())
		}

		gen, err := generator.MakeGenerator(param)
		if err != nil {
			log.Fatalln(err.Error())
		}

		gen.Generate(ts, param)

		return nil
	},
}
