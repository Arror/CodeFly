package context

import "github.com/samuel/go-thrift/parser"

// Context Generate context
type Context struct {
	Lang    string
	Input   string
	Output  string
	Thrifts map[string]*parser.Thrift
}
