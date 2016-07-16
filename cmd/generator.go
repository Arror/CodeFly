package cmd

import (
	"CodeFly/global"

	"github.com/urfave/cli"
)

// Gen 代码生成命令
var Gen = cli.Command{
	Name:      "gen",
	ShortName: "g",
	Usage:     "Code generator.",
	UsageText: "Generate target language file using thrift file",
	Flags: []cli.Flag{
		cli.StringFlag{
			Name:        "lang, l",
			Usage:       "Target language name",
			Destination: &global.GenCmdInfo.Lang,
		},
		cli.StringFlag{
			Name:        "input, i",
			Usage:       "Input thrift file path",
			Destination: &global.GenCmdInfo.Input,
		},
		cli.StringFlag{
			Name:        "output, o",
			Usage:       "Target language file output path",
			Destination: &global.GenCmdInfo.Output,
		},
	},
	Action: func(c *cli.Context) error {
		if err := global.GenCmdInfo.CheckValidity(); err != nil {
			return err
		}
		return nil
	},
}
