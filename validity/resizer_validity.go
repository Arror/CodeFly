package validity

import (
	"errors"
	"path/filepath"
	"strings"
)

// ImageResizeInfo 命令信息结构
type ImageResizeInfo struct {
	Input         string
	Dir           string
	File          string
	FileName      string
	FileExtension string
}

// CheckImageResizeInfoInfoValidity 检查输入命令合法性
func (iri *ImageResizeInfo) CheckImageResizeInfoInfoValidity() error {

	if iri.Input == "" {
		return errors.New("Icon文件路径为空")
	}

	p, err := filepath.Abs(iri.Input)

	if err != nil {
		return errors.New("Icon文件路径错误")
	}

	_, file := filepath.Split(p)

	if file == "" {
		return errors.New("Icon文件路径错误")
	}

	components := strings.Split(file, ".")

	if len(components) < 2 {
		return errors.New("不是一个文件的路径")
	}

	if components[1] != "png" {
		return errors.New("图片必须是*.png的")
	}

	iri.Input = p
	iri.File = file
	iri.FileName = components[0]
	iri.FileExtension = components[1]

	return nil
}
