package cmd

import (
	"log"

	"CodeFly/maker"
	"CodeFly/parameter"
	"CodeFly/reader"

	"github.com/urfave/cli"
)

var param = &parameter.Parameter{}

// JSONGenerate Json generate command
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

		generator, err := maker.MakeGenerator(param)
		if err != nil {
			log.Fatalln(err.Error())
		}

		generator.Generate(ts, param)

		return nil
	},
}
