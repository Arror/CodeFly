package generator

import (
	"fmt"

	"CodeFly/lang"
	"CodeFly/lang/swift/model"
	"CodeFly/parameter"

	"github.com/arrors/go-thrift/parser"
)

// Generator Generator协议
type Generator interface {
	Generate(ts map[string]*parser.Thrift, param *parameter.Parameter)
}

// MakeGenerator 创建Generator对象
func MakeGenerator(param *parameter.Parameter) (Generator, error) {

	switch param.Lang {
	case lang.Swift:
		return &model.SwiftGenerator{}, nil
	default:
		return nil, fmt.Errorf("Create %s language generator failed", param.Lang)
	}
}
