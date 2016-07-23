package reader

import (
	"fmt"

	"CodeFly/validity"

	"github.com/samuel/go-thrift/parser"
)

// ThriftReader Reader结构
type ThriftReader struct {
	Thrifts    map[string]*parser.Thrift
	Thrift     *parser.Thrift
	InputPath  string
	OutputPath string
}

// ReadThrift 读取thrift文件
func (r *ThriftReader) ReadThrift(info *validity.GenerateCommandInfo) error {

	p := parser.Parser{}

	thrifts, _, err := p.ParseFile(info.Input)

	if err != nil {
		return fmt.Errorf("解析thrift文件失败")
	}

	r.Thrifts = thrifts
	r.Thrift = thrifts[info.Input]
	r.InputPath = info.Input
	r.OutputPath = info.Output

	return nil
}

// CheckNameSpace 检查Namespace
func (r *ThriftReader) CheckNameSpace(info *validity.GenerateCommandInfo) error {

	ts := r.Thrifts
	lang := info.Lang

	for n, t := range ts {
		if t.Namespaces[lang] == "" {
			return fmt.Errorf("发现%s.thrift文件没有关于%s语言的Namespace信息", n, lang)
		}
	}
	return nil
}
