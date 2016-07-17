package resizer

import (
	"errors"
	"fmt"
	"image/png"
	"os"

	"CodeFly/global"

	"github.com/nfnt/resize"
)

// Resizer Resizer结构体
type Resizer struct{}

// IconResizer Resizer对象
var IconResizer = &Resizer{}

// ResizeIcon 转换输入的Icon
func (rs *Resizer) ResizeIcon(info *global.ImageResizeInfo, width, height uint) error {

	file, err := os.Open(info.Input)

	if err != nil {
		return errors.New("图片文件打开失败")
	}

	img, err := png.Decode(file)

	m := resize.Resize(width, height, img, resize.Lanczos3)

	newIconName := fmt.Sprintf("%s_%dx%d.png", info.FileName, width, height)

	out, err := os.Create(newIconName)
	if err != nil {
		return errors.New("文件创建失败")
	}
	defer out.Close()

	png.Encode(out, m)

	return nil
}
