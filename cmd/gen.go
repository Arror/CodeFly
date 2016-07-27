package cmd

import (
	"fmt"
	"path/filepath"

	"CodeFly/generator"
	"CodeFly/parser"

	"github.com/urfave/cli"
)

// GenerateCommandInfo 命令信息结构
type GenerateCommandInfo struct {
	Lang   string
	Input  string
	Output string
}

var genInfo = &GenerateCommandInfo{}

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

		if err := genInfo.checkGenerateCommandInfo(); err != nil {
			return err
		}

		ts, err := generator.ReadThrift(genInfo.Input)
		if err != nil {
			return err
		}

		if err := generator.CheckNameSpace(genInfo.Lang, ts); err != nil {
			return err
		}

		if err := generator.Distributor(ts, genInfo.Lang, genInfo.Input, genInfo.Output); err != nil {
			return err
		}

		return nil
	},
}

func (gci *GenerateCommandInfo) checkGenerateCommandInfo() error {

	if gci.Lang == "" {
		return fmt.Errorf("语言名称为空")
	}
	switch gci.Lang {
	case parser.Swift:
		break
	default:
		return fmt.Errorf("未被支持的语言")
	}

	if gci.Input == "" {
		return fmt.Errorf("thrift文件路径为空")
	}
	p, err := filepath.Abs(gci.Input)
	if err != nil {
		return fmt.Errorf("thrift文件路径错误")
	}
	gci.Input = p

	if gci.Output == "" {
		return fmt.Errorf("输出文件夹路径为空")
	}
	p, err = filepath.Abs(gci.Output)
	if err != nil {
		return fmt.Errorf("输出文件路径错误")
	}
	gci.Output = p

	return nil
}
