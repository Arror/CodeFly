package parameter

import (
	"fmt"
	"os"
	"path/filepath"

	"CodeFly/global"
)

// Parameter 命令参数结构
type Parameter struct {
	Lang   string
	Input  string
	Output string
}

// CheckParameter 检查参数信息
func (param *Parameter) CheckParameter() error {

	lang, err := checkLang(param.Lang)
	if err != nil {
		return err
	}
	param.Lang = lang

	input, err := checkInputPath(param.Input)
	if err != nil {
		return err
	}
	param.Input = input

	output, err := checkOutputPath(param.Output)
	if err != nil {
		return err
	}
	param.Output = output

	return nil
}

func validLang(lang string) bool {
	switch lang {
	case global.Swift:
		return true
	default:
		return false
	}
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
