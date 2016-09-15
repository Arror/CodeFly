package model

import (
	"fmt"
	"path/filepath"
)

// GenerateCommandInfo 命令信息结构
type GenerateCommandInfo struct {
	Lang   string
	Input  string
	Output string
}

// CheckGenerateCommandInfo 检查命令信息
func (gci *GenerateCommandInfo) CheckGenerateCommandInfo() error {

	if gci.Lang == "" {
		return fmt.Errorf("语言名称为空")
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
