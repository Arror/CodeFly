package resizer

import (
	"errors"
	"fmt"
	"image/png"
	"os"

	"CodeFly/global"

	"github.com/nfnt/resize"
)

// IconSize iOS 设备所需Icon尺寸
type IconSize struct {
	BaseWidth   uint
	BasegHeight uint
	Multiply    uint
}

var iconSizes = []*IconSize{
	&IconSize{
		BaseWidth:   29,
		BasegHeight: 29,
		Multiply:    2,
	},
	&IconSize{
		BaseWidth:   29,
		BasegHeight: 29,
		Multiply:    3,
	},
	&IconSize{
		BaseWidth:   40,
		BasegHeight: 40,
		Multiply:    2,
	},
	&IconSize{
		BaseWidth:   40,
		BasegHeight: 40,
		Multiply:    3,
	},
	&IconSize{
		BaseWidth:   60,
		BasegHeight: 60,
		Multiply:    2,
	},
	&IconSize{
		BaseWidth:   60,
		BasegHeight: 60,
		Multiply:    3,
	},
}

func (is *IconSize) factualImageSize() (width, height uint) {
	return is.BaseWidth * is.Multiply, is.BasegHeight * is.Multiply
}

func (is *IconSize) imageName(base string) string {
	width, height := is.factualImageSize()
	return fmt.Sprintf("%s_%dx%d_@%d.png", base, width, height, is.Multiply)
}

// Resizer Resizer结构体
type Resizer struct{}

// IconResizer Resizer对象
var IconResizer = &Resizer{}

// ResizeIcon 转换输入的Icon
func (rs *Resizer) ResizeIcon(info *global.ImageResizeInfo) error {

	file, err := os.Open(info.Input)

	if err != nil {
		return errors.New("图片文件打开失败")
	}

	img, err := png.Decode(file)

	for _, size := range iconSizes {

		w, h := size.factualImageSize()

		fn := size.imageName(info.FileName)

		m := resize.Resize(w, h, img, resize.Lanczos3)

		newIconName := fmt.Sprintf(fn)

		out, err := os.Create(newIconName)
		if err != nil {
			return errors.New("文件创建失败")
		}
		defer out.Close()

		png.Encode(out, m)
	}

	return nil
}
