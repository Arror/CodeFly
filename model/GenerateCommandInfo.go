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
		return fmt.Errorf("The target language name is empty")
	}

	if gci.Input == "" {
		return fmt.Errorf("The thrift file input path is empty")
	}
	p, err := filepath.Abs(gci.Input)
	if err != nil {
		return fmt.Errorf("The input thrift file path error")
	}
	gci.Input = p

	if gci.Output == "" {
		return fmt.Errorf("File output path is empty")
	}
	p, err = filepath.Abs(gci.Output)
	if err != nil {
		return fmt.Errorf("File output path error")
	}
	gci.Output = p

	return nil
}
