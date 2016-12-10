package reader

import (
	"fmt"

	"CodeFly/command"

	"github.com/arrors/go-thrift/parser"
)

// ReadThrift 读取Thrift文件信息
func ReadThrift(genInfo *command.Command) (map[string]*parser.Thrift, error) {

	p := parser.Parser{}

	ts, _, err := p.ParseFile(genInfo.Input)

	if err != nil {
		return nil, err
	}

	for n, t := range ts {
		if t.Namespaces[genInfo.Lang] == "" {
			return nil, fmt.Errorf("%s language namespace info not found in %s.thrift", genInfo.Lang, n)
		}
	}

	return ts, nil
}
