package command

import (
	"fmt"
	"os"
	"path/filepath"
)

const (
	swift = "swift"
)

func validLang(lang string) bool {
	switch lang {
	case swift:
		return true
	default:
		return false
	}
}

// Command 命令信息结构
type Command struct {
	Lang   string
	Input  string
	Output string
}

// CheckCommand 检查命令信息
func (c *Command) CheckCommand() error {

	lang, err := checkLang(c.Lang)
	if err != nil {
		return err
	}
	c.Lang = lang

	input, err := checkInputPath(c.Input)
	if err != nil {
		return err
	}
	c.Input = input

	output, err := checkOutputPath(c.Output)
	if err != nil {
		return err
	}
	c.Output = output

	return nil
}

func checkLang(lang string) (string, error) {

	if lang == "" {
		return "", fmt.Errorf("The target language name is empty")
	}

	if !validLang(lang) {
		return "", fmt.Errorf("Unsupported language")
	}

	return lang, nil
}

func checkInputPath(ip string) (string, error) {

	if ip == "" {
		return "", fmt.Errorf("The thrift file input path is empty")
	}

	p, err := filepath.Abs(ip)

	if err != nil {
		return "", fmt.Errorf("The input thrift file path not exist")
	}

	info, err := os.Stat(p)

	if err != nil {
		return "", fmt.Errorf("The input thrift file path not exist")
	}

	if info.IsDir() {
		return "", fmt.Errorf("The input path is a directory")
	}

	return p, nil
}

func checkOutputPath(op string) (string, error) {

	if op == "" {
		return "", fmt.Errorf("File output path is empty")
	}

	p, err := filepath.Abs(op)

	if err != nil {
		return "", fmt.Errorf("File output path error")
	}

	return p, nil
}
