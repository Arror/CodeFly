package cmd

import (
	"CodeFly/resizer"
	"CodeFly/validity"

	"github.com/urfave/cli"
)

// ImageResize 命令信息对象
var resizeInfo = &validity.ImageResizeInfo{}

// IconResizer Resizer对象
var iconResizer = &resizer.Resizer{}

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
			Destination: &resizeInfo.Input,
		},
	},
	Action: func(c *cli.Context) error {

		if err := resizeInfo.CheckImageResizeInfoInfoValidity(); err != nil {
			return err
		}

		if err := iconResizer.ResizeIcon(resizeInfo); err != nil {
			return err
		}
		return nil
	},
}
