package cmd

import (
	"CodeFly/global"
	"CodeFly/resizer"

	"github.com/urfave/cli"
)

// Resize Icon 生成命令
var Resize = cli.Command{
	Name:      "resize",
	ShortName: "r",
	Usage:     "iOS Icon 生成命令",
	UsageText: "生成 iOS Icon 尺寸的图片文件",
	Flags: []cli.Flag{
		cli.StringFlag{
			Name:        "input, i",
			Usage:       "icon 文件路径",
			Destination: &global.ImageResize.Input,
		},
	},
	Action: func(c *cli.Context) error {

		info := global.ImageResize

		if err := info.CheckImageResizeInfoInfoValidity(); err != nil {
			return err
		}

		iconResizer := resizer.IconResizer
		if err := iconResizer.ResizeIcon(info); err != nil {
			return err
		}
		return nil
	},
}
