package context

import "text/template"

// SwiftContext Swift context
type SwiftContext struct {
	Context

	EnumTemplateName    string
	StructTemplateName  string
	ServiceTemplateName string

	EnumTemplate    *template.Template
	StructTemplate  *template.Template
	ServiceTemplate *template.Template
}

// InitSwiftContext Swift context init
func InitSwiftContext(ctx Context) SwiftContext {

	sCtx := SwiftContext{}

	sCtx.Context = ctx

	enumName := "Enum"
	sCtx.EnumTemplateName = enumName
	sCtx.EnumTemplate = initTemplate(enumName, "templates/swift/enum.tpl")

	structName := "Struct"
	sCtx.StructTemplateName = structName
	sCtx.StructTemplate = initTemplate(structName, "templates/swift/struct.tpl")

	serviceName := "Service"
	sCtx.ServiceTemplateName = serviceName
	sCtx.ServiceTemplate = initTemplate(serviceName, "templates/swift/service.tpl")

	return sCtx
}
