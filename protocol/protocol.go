package protocol

import (
	"CodeFly/parameter"

	"github.com/arrors/go-thrift/parser"
)

// Generator Generator protocol
type Generator interface {
	Generate(ts map[string]*parser.Thrift, param *parameter.Parameter)
}
