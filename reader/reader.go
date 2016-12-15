package reader

import (
	"fmt"

	"CodeFly/parameter"

	"github.com/samuel/go-thrift/parser"
)

// ReadThrift Read thrift file via input infomation
func ReadThrift(param *parameter.Parameter) (map[string]*parser.Thrift, error) {

	p := parser.Parser{}

	ts, _, err := p.ParseFile(param.Input)

	if err != nil {
		return nil, err
	}

	for n, t := range ts {
		if t.Namespaces[param.Lang] == "" {
			return nil, fmt.Errorf("%s language namespace info not found in %s.thrift", param.Lang, n)
		}
	}

	return ts, nil
}
