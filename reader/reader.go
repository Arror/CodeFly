package reader

import (
	"fmt"

	"github.com/arrors/go-thrift/parser"
)

// ReadThrift 读取Thrift文件信息
func ReadThrift(ip string) (map[string]*parser.Thrift, error) {

	p := parser.Parser{}

	thrifts, _, err := p.ParseFile(ip)

	if err != nil {
		return nil, err
	}
	return thrifts, nil
}

// CheckLanguageNameSpace 检查Namespace信息
func CheckLanguageNameSpace(lang string, ts map[string]*parser.Thrift) error {

	for n, t := range ts {
		if t.Namespaces[lang] == "" {
			return fmt.Errorf("%s language namespace info not found in %s.thrift", lang, n)
		}
	}
	return nil
}
