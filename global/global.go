package global

import (
	"errors"
	"path/filepath"
	"strings"
)

const (
	// Swfit Swift语言
	Swfit = "swift"
)

// GenerateCommandInfo 命令信息结构
type GenerateCommandInfo struct {
	Lang   string
	Input  string
	Output string
}

// GenCmdInfo 命令信息对象
var GenCmdInfo = &GenerateCommandInfo{}

// CheckGenerateCommandInfoValidity 检查输入命令合法性
func (gci *GenerateCommandInfo) CheckGenerateCommandInfoValidity() error {

	if gci.Lang == "" {
		return errors.New("语言名称为空")
	}
	switch gci.Lang {
	case Swfit:
		break
	default:
		return errors.New("未被支持的语言")
	}

	if gci.Input == "" {
		return errors.New("thrift文件路径为空")
	}
	p, err := filepath.Abs(gci.Input)
	if err != nil {
		return errors.New("thrift文件路径错误")
	}
	gci.Input = p

	if gci.Output == "" {
		return errors.New("输出文件夹路径为空")
	}
	p, err = filepath.Abs(gci.Output)
	if err != nil {
		return errors.New("输出文件路径错误")
	}
	gci.Output = p

	return nil
}

// ImageResizeInfo 命令信息结构
type ImageResizeInfo struct {
	Input         string
	Dir           string
	File          string
	FileName      string
	FileExtension string
}

// ImageResize 命令信息对象
var ImageResize = &ImageResizeInfo{}

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
