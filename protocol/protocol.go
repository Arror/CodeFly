package protocol

import (
	"CodeFly/parameter"

	"github.com/samuel/go-thrift/parser"
)

// Generator Generator protocol
type Generator interface {
	Generate(ts map[string]*parser.Thrift, param *parameter.Parameter)
}
