package context

import (
	"log"
	"text/template"

	"github.com/Arror/CodeFly/templates"
	"github.com/samuel/go-thrift/parser"
)

// Context Generate context
type Context struct {
	Lang   string
	Input  string
	Output string

	Thrift *parser.Thrift

	Thrifts map[string]*parser.Thrift
}

// InitContext Context init
func InitContext(lang string, input string, output string, thrifts map[string]*parser.Thrift) Context {

	ctx := Context{}

	ctx.Lang = lang
	ctx.Input = input
	ctx.Output = output

	ctx.Thrifts = thrifts

	ctx.Thrift = thrifts[input]

	return ctx
}

func initTemplate(name string, path string) *template.Template {

	buffer := templates.MustAsset(path)

	template, err := template.New(name).Parse(string(buffer))
	if err != nil {
		log.Fatal(err.Error())
	}
	return template
}
