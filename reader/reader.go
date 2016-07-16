package reader

import (
	"errors"

	"CodeFly/global"

	"github.com/samuel/go-thrift/parser"
)

// ThriftReader Reader结构
type ThriftReader struct {
	Thrifts    map[string]*parser.Thrift
	OutputPath string
}

// Reader ThriftReader对象
var Reader = &ThriftReader{}

// ReadThrift 读取thrift文件
func (r *ThriftReader) ReadThrift(info *global.GenerateCommandInfo) error {

	p := parser.Parser{}

	thrifts, _, err := p.ParseFile(info.Input)

	if err != nil {
		return errors.New("无法解析thrift文件")
	}

	r.Thrifts = thrifts
	r.OutputPath = info.Output

	return nil
}
