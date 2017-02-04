package context

import (
	"text/template"

	"CodeFly/templates"
)

// SwiftContext Swift context
type SwiftContext struct {
	Context

	EnumTemplateName    string
	StructTemplateName  string
	ServiceTemplateName string

	EmunTemplate    *template.Template
	StructTemplate  *template.Template
	ServiceTemplate *template.Template
}

// InitSwiftContext Swift context init
func InitSwiftContext(ctx Context) SwiftContext {

	sCtx := SwiftContext{}

	sCtx.Context = ctx

	enumName := "Enum"
	sCtx.EnumTemplateName = enumName
	sCtx.EmunTemplate = templates.InitTemplate(enumName, "templates/swift/enum.tpl")

	structName := "Struct"
	sCtx.StructTemplateName = structName
	sCtx.StructTemplate = templates.InitTemplate(structName, "templates/swift/struct.tpl")

	serviceName := "Service"
	sCtx.ServiceTemplateName = serviceName
	sCtx.ServiceTemplate = templates.InitTemplate(serviceName, "templates/swift/service.tpl")

	return sCtx
}
