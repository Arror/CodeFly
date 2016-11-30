package model

import (
	"fmt"
	"os"
	"path/filepath"
)

const (
	// Swift Swift 语言
	Swift = "swift"
	// ObjectiveC ObjectiveC 语言
	ObjectiveC = "objectivec"
	// Go Go 语言
	Go = "go"
	// Java Java语言
	Java = "java"
	// JavaScript JavaScript语言
	JavaScript = "javascript"
)

var langs = [...]string{Swift}

func validLang(lang string) bool {
	for _, l := range langs {
		if l == lang {
			return true
		}
		continue
	}
	return false
}

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

	if !validLang(gci.Lang) {
		return fmt.Errorf("Unsupported language")
	}

	if gci.Input == "" {
		return fmt.Errorf("The thrift file input path is empty")
	}
	p, err := filepath.Abs(gci.Input)
	if err != nil {
		return fmt.Errorf("The input thrift file path not exist")
	}
	inputFileInfo, err := os.Stat(p)
	if err != nil {
		return fmt.Errorf("The input thrift file path not exist")
	}
	if inputFileInfo.IsDir() {
		return fmt.Errorf("The input thrift file path not exist")
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
