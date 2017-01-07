package swift

import (
	"os"
	"strings"
	"sync"

	"CodeFly/global"
	"CodeFly/writer"

	"github.com/samuel/go-thrift/parser"
)

const (
	enumTplPath    = "templates/swift/enum.tpl"
	structTplPath  = "templates/swift/struct.tpl"
	serviceTplPath = "templates/swift/service.tpl"
)

const (
	enumTplName    = "EnumTemplate"
	structTplName  = "StructTemplate"
	serviceTplName = "ServiceTemplate"
)

// GenContext Generate context
type GenContext struct {
	lang    string
	input   string
	output  string
	thrifts map[string]*parser.Thrift
	thrift  *parser.Thrift
}

// EnumGenContext Enum generate context
type EnumGenContext struct {
	*parser.Enum
	*GenContext
}

// StructGenContext Struct generate context
type StructGenContext struct {
	*parser.Struct
	*GenContext
}

// ServiceGenContext Service generate context
type ServiceGenContext struct {
	*parser.Service
	*GenContext
}

// Generate GenContext implement Generator interface
func (ctx *GenContext) Generate() {

	if err := os.MkdirAll(global.Output, 0755); err != nil {
		panic(err.Error())
	}

	ctx.lang = global.Lang
	ctx.input = global.Input
	ctx.output = global.Output
	ctx.thrifts = global.ThriftMapping
	ctx.thrift = global.ThriftMapping[ctx.input]

	wg := sync.WaitGroup{}

	wg.Add(1)
	go func() {
		defer wg.Done()

		enumTpl := writer.InitTemplate(enumTplName, enumTplPath)

		for _, e := range ctx.thrift.Enums {

			eCtx := &EnumGenContext{}
			eCtx.Enum = e
			eCtx.GenContext = ctx

			writer.WriteFile(writer.AssembleFilePath(eCtx.GenContext.output, eCtx.Name()+".swift"), enumTpl, enumTplName, eCtx)
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()

		structTpl := writer.InitTemplate(structTplName, structTplPath)

		for _, s := range ctx.thrift.Structs {

			sCtx := &StructGenContext{}
			sCtx.Struct = s
			sCtx.GenContext = ctx

			writer.WriteFile(writer.AssembleFilePath(sCtx.GenContext.output, sCtx.Name()+".swift"), structTpl, structTplName, sCtx)
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()

		serviceTpl := writer.InitTemplate(serviceTplName, serviceTplPath)

		for _, s := range ctx.thrift.Services {

			sCtx := &ServiceGenContext{}
			sCtx.Service = s
			sCtx.GenContext = ctx

			writer.WriteFile(writer.AssembleFilePath(sCtx.GenContext.output, sCtx.Name()+".swift"), serviceTpl, serviceTplName, sCtx)
		}
	}()

	wg.Wait()
}

// Name Enum name
func (ctx *EnumGenContext) Name() string {
	return ctx.GenContext.thrift.Namespaces[ctx.GenContext.lang] + ctx.Enum.Name
}

// Name Struct name
func (ctx *StructGenContext) Name() string {
	return ctx.GenContext.thrift.Namespaces[ctx.GenContext.lang] + ctx.Struct.Name
}

// Name Service name
func (ctx *ServiceGenContext) Name() string {
	return ctx.Service.Name + "Service"
}

// MethodName Method name
func (ctx *ServiceGenContext) MethodName(m *parser.Method) string {
	return strings.ToLower(m.Name[:1]) + m.Name[1:]
}

// TypeString Type string
func (ctx *GenContext) TypeString(t *parser.Type) string {

	if t == nil {
		return swiftVoid
	}

	switch t.Name {
	case global.List:
		switch t.ValueType.Name {
		case global.List, global.Set, global.Map:
			panic("Unsupported inner container type.")
		}
		return "[" + ctx.TypeString(t.ValueType) + "]"
	case global.Map, global.Set:
		panic("Unsupported container type.")
	}

	if base := mapping[t.Name]; base != "" {
		return base
	}

	typeComponents := strings.Split(t.Name, ".")

	componentCount := len(typeComponents)

	var _thrift *parser.Thrift
	var _type string

	switch componentCount {
	case 1:
		_thrift = ctx.thrift
		_type = typeComponents[0]
	case 2:
		if key := ctx.thrift.Includes[typeComponents[0]]; key != "" {
			_thrift = ctx.thrifts[key]
			_type = typeComponents[1]
		} else {
			panic(typeComponents[0] + ".thrift not find in file include.")
		}
	}

	if _thrift != nil && _type != "" {
		return _thrift.Namespaces[global.Swift] + _type
	}

	panic("Unsupported type.")
}

const (
	swiftInt    = "Int"
	swiftInt64  = "Int64"
	swiftDouble = "Double"
	swiftBool   = "Bool"
	swiftString = "String"
)

const (
	swiftVoid = "Void"
)

var mapping = map[string]string{
	global.ThriftI16:    swiftInt,
	global.ThriftI32:    swiftInt,
	global.ThriftI64:    swiftInt64,
	global.ThriftBool:   swiftBool,
	global.ThriftDouble: swiftDouble,
	global.ThriftString: swiftString,
	global.ThriftByte:   global.Unsupported,
	global.ThriftBinary: global.Unsupported,
}
