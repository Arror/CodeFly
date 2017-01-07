package context

import (
	"log"
	"text/template"

	"github.com/samuel/go-thrift/parser"

	"CodeFly/global"
	"CodeFly/templates"
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

// InitTemplates Init Templates
func (ctx *Context) InitTemplates() {

	switch ctx.Lang {
	case global.Swift:
		enumName := "Enum"
		enumPath := "templates/swift/enum.tpl"
		ctx.EnumTemplateName = enumName
		ctx.EmunTemplate = initTemplate(enumName, enumPath)

		structName := "Struct"
		structPath := "templates/swift/struct.tpl"
		ctx.StructTemplateName = structName
		ctx.StructTemplate = initTemplate(structName, structPath)

		serviceName := "Service"
		servicePath := "templates/swift/service.tpl"
		ctx.EnumTemplateName = serviceName
		ctx.EmunTemplate = initTemplate(serviceName, servicePath)
	}
}

func initTemplate(name string, path string) *template.Template {

	buffer := templates.MustAsset(path)

	template, err := template.New(name).Parse(string(buffer))
	if err != nil {
		log.Fatal(err.Error())
	}
	return template
}
