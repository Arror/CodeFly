package reader

import (
	"errors"

	"github.com/samuel/go-thrift/parser"
)

// ThriftReader Reader结构
type ThriftReader struct {
	Thrifts map[string]*parser.Thrift
	Path    string
}

// Reader ThriftReader对象
var Reader = &ThriftReader{}

// ReadThrift 读取thrift文件
func (r *ThriftReader) ReadThrift(path string) error {

	p := parser.Parser{}

	thrifts, _, err := p.ParseFile(path)

	if err != nil {
		return errors.New("无法解析thrift文件")
	}

	r.Thrifts = thrifts
	r.Path = path

	return nil
}
