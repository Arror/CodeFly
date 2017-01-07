package context

import (
	"text/template"

	"github.com/samuel/go-thrift/parser"
)

// Context Generate context
type Context struct {
	Lang   string
	Input  string
	Output string

	Thrift *parser.Thrift

	Thrifts map[string]*parser.Thrift

	EnumTemplateName    string
	StructTemplateName  string
	ServiceTemplateName string

	EmunTemplate    *template.Template
	StructTemplate  *template.Template
	ServiceTemplate *template.Template
}
