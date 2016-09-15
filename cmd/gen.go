package cmd

import (
	"fmt"

	"CodeFly/distributor"
	"CodeFly/model"
	"CodeFly/reader"

	"github.com/urfave/cli"
)

var genInfo = &model.GenerateCommandInfo{}

// Gen 代码生成命令
var Gen = cli.Command{
	Name:      "gen",
	ShortName: "g",
	Usage:     "代码生成命令",
	UsageText: "通过thrift文件生成目标语言代码",
	Flags: []cli.Flag{
		cli.StringFlag{
			Name:        "lang, l",
			Usage:       "目标语言名称",
			Destination: &genInfo.Lang,
		},
		cli.StringFlag{
			Name:        "input, i",
			Usage:       "被生成thrift文件路径",
			Destination: &genInfo.Input,
		},
		cli.StringFlag{
			Name:        "output, o",
			Usage:       "生成文件的输出路径",
			Destination: &genInfo.Output,
		},
	},
	Action: func(c *cli.Context) error {

		if err := genInfo.CheckGenerateCommandInfo(); err != nil {
			return err
		}

		switch genInfo.Lang {
		case model.Swift:
			break
		default:
			return fmt.Errorf("未被支持的语言")
		}

		ts, err := reader.ReadThrift(genInfo.Input)
		if err != nil {
			return err
		}

		if err := reader.CheckLanguageNameSpace(genInfo.Lang, ts); err != nil {
			return err
		}

		if err := distributor.Distribute(ts, genInfo); err != nil {
			return err
		}

		return nil
	},
}
