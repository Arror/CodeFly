package cmd

import (
	"CodeFly/generator/operator"
	"CodeFly/generator/reader"
	"CodeFly/validity"

	"github.com/urfave/cli"
)

// GenCmdInfo 命令信息对象
var genCmdInfo = &validity.GenerateCommandInfo{}

// Reader ThriftReader对象
var commonReader = &reader.ThriftReader{}

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
			Destination: &genCmdInfo.Lang,
		},
		cli.StringFlag{
			Name:        "input, i",
			Usage:       "被生成thrift文件路径",
			Destination: &genCmdInfo.Input,
		},
		cli.StringFlag{
			Name:        "output, o",
			Usage:       "生成文件的输出路径",
			Destination: &genCmdInfo.Output,
		},
	},
	Action: func(c *cli.Context) error {

		if err := genCmdInfo.CheckGenerateCommandInfoValidity(); err != nil {
			return err
		}

		if err := commonReader.ReadThrift(genCmdInfo); err != nil {
			return err
		}

		if err := commonReader.CheckNameSpace(genCmdInfo); err != nil {
			return err
		}

		if err := operator.Print(commonReader, genCmdInfo); err != nil {
			return err
		}

		return nil
	},
}
