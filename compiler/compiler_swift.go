package compiler

import (
	"os"
	"strings"
	"sync"
	"text/template"

	"github.com/samuel/go-thrift/parser"

	"CodeFly/context"
	"CodeFly/global"
)

// SwiftCompiler Swift Code Compiler
type SwiftCompiler struct{}

// SwiftContext Swift code context
type SwiftContext struct {
	Ctx *context.Context

	EnumTemplateName    string
	StructTemplateName  string
	ServiceTemplateName string

	EmunTemplate    *template.Template
	StructTemplate  *template.Template
	ServiceTemplate *template.Template
}

func initSwiftContext(ctx *context.Context) *SwiftContext {

	sCtx := &SwiftContext{}

	sCtx.Ctx = ctx

	enumName := "Enum"
	sCtx.EnumTemplateName = enumName
	sCtx.EmunTemplate = initTemplate(enumName, "templates/swift/enum.tpl")

	structName := "Struct"
	sCtx.StructTemplateName = structName
	sCtx.StructTemplate = initTemplate(structName, "templates/swift/struct.tpl")

	serviceName := "Service"
	sCtx.ServiceTemplateName = serviceName
	sCtx.ServiceTemplate = initTemplate(serviceName, "templates/swift/service.tpl")

	return sCtx
}

// SwiftEnum Swift Enum
type SwiftEnum struct {
	SCtx *SwiftContext
	*parser.Enum
}

// SwiftStruct Swift Struct
type SwiftStruct struct {
	SCtx *SwiftContext
	*parser.Struct
}

// SwiftService Swift Service
type SwiftService struct {
	SCtx *SwiftContext
	*parser.Service
}

// Name Enum name
func (se *SwiftEnum) Name() string {
	return se.SCtx.Ctx.Thrift.Namespaces[se.SCtx.Ctx.Lang] + se.Enum.Name
}

// Name Struct name
func (ss *SwiftStruct) Name() string {
	return ss.SCtx.Ctx.Thrift.Namespaces[ss.SCtx.Ctx.Lang] + ss.Struct.Name
}

// Name Service name
func (ss *SwiftService) Name() string {
	return ss.Service.Name + "Service"
}

// MethodName Method Name
func (ss *SwiftService) MethodName(m *parser.Method) string {
	return strings.ToLower(m.Name[:1]) + m.Name[1:]
}

func (sc *SwiftCompiler) genCodes(ctx *context.Context) {

	sCtx := initSwiftContext(ctx)

	if err := os.MkdirAll(ctx.Output, 0755); err != nil {
		panic(err.Error())
	}

	wg := sync.WaitGroup{}

	wg.Add(1)
	go func() {
		defer wg.Done()
		for _, e := range ctx.Thrift.Enums {

			se := &SwiftEnum{
				SCtx: sCtx,
				Enum: e,
			}

			path := assembleFilePath(ctx.Output, se.Name()+".swift")

			writeFile(path, sCtx.EmunTemplate, sCtx.EnumTemplateName, se)
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		for _, s := range ctx.Thrift.Structs {

			ss := &SwiftStruct{
				SCtx:   sCtx,
				Struct: s,
			}

			path := assembleFilePath(ctx.Output, ss.Name()+".swift")

			writeFile(path, sCtx.StructTemplate, sCtx.StructTemplateName, ss)
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		for _, s := range ctx.Thrift.Services {

			ss := &SwiftService{
				SCtx:    sCtx,
				Service: s,
			}

			path := assembleFilePath(ctx.Output, ss.Name()+".swift")

			writeFile(path, sCtx.ServiceTemplate, sCtx.ServiceTemplateName, ss)
		}
	}()

	wg.Wait()
}

// TypeString Type string
func (sCtx *SwiftContext) TypeString(t *parser.Type) string {

	if t == nil {
		return swiftVoid
	}

	switch t.Name {
	case global.ThriftList:
		switch t.ValueType.Name {
		case global.ThriftList, global.ThriftSet, global.ThriftMap:
			panic("Unsupported inner container type.")
		}
		return "[" + sCtx.TypeString(t.ValueType) + "]"
	case global.ThriftMap, global.ThriftSet:
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
		_thrift = sCtx.Ctx.Thrift
		_type = typeComponents[0]
	case 2:
		if key := sCtx.Ctx.Thrift.Includes[typeComponents[0]]; key != "" {
			_thrift = sCtx.Ctx.Thrifts[key]
			_type = typeComponents[1]
		} else {
			panic(typeComponents[0] + ".thrift not find in file include.")
		}
	}

	if _thrift != nil && _type != "" {
		return _thrift.Namespaces[sCtx.Ctx.Lang] + _type
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
	global.ThriftByte:   global.UnsupportedType,
	global.ThriftBinary: global.UnsupportedType,
}
