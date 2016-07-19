package cmd

import (
	"CodeFly/global"
	"CodeFly/reader"

	"github.com/urfave/cli"
)

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
			Destination: &global.GenCmdInfo.Lang,
		},
		cli.StringFlag{
			Name:        "input, i",
			Usage:       "被生成thrift文件路径",
			Destination: &global.GenCmdInfo.Input,
		},
		cli.StringFlag{
			Name:        "output, o",
			Usage:       "生成文件的输出路径",
			Destination: &global.GenCmdInfo.Output,
		},
	},
	Action: func(c *cli.Context) error {

		info := global.GenCmdInfo

		if err := info.CheckGenerateCommandInfoValidity(); err != nil {
			return err
		}

		if err := reader.Reader.ReadThrift(info); err != nil {
			return err
		}

		if err := reader.Reader.CheckNameSpace(info); err != nil {
			return err
		}

		// TODO: 测试入口，需要删除
		reader.SwiftReader.InitSwiftThrift(reader.Reader)

		return nil
	},
}
