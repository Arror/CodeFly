package compiler

import (
	"strings"

	"github.com/samuel/go-thrift/parser"

	"CodeFly/context"
	"CodeFly/global"
	"CodeFly/writer"
)

// SwiftCompiler Swift Code Compiler
type SwiftCompiler struct{}

// Swift Swift
type Swift struct {
	Ctx *context.Context
}

// SwiftEnum Swift Enum
type SwiftEnum struct {
	*Swift
	*parser.Enum
}

// SwiftStruct Swift Struct
type SwiftStruct struct {
	*Swift
	*parser.Struct
}

// SwiftService Swift Service
type SwiftService struct {
	*Swift
	*parser.Service
}

var swift = &Swift{}

// Name Enum name
func (se *SwiftEnum) Name() string {
	return se.Swift.Ctx.Thrift.Namespaces[se.Swift.Ctx.Lang] + se.Enum.Name
}

// Name Struct name
func (ss *SwiftStruct) Name() string {
	return ss.Swift.Ctx.Thrift.Namespaces[ss.Swift.Ctx.Lang] + ss.Struct.Name
}

// Name Service name
func (ss *SwiftService) Name() string {
	return ss.Service.Name + "Service"
}

// MethodName Method Name
func (ss *SwiftService) MethodName(m *parser.Method) string {
	return strings.ToLower(m.Name[:1]) + m.Name[1:]
}

func (sc *SwiftCompiler) genEnumCode(ctx *context.Context, e *parser.Enum) {

	swift.Ctx = ctx

	swiftEnum := &SwiftEnum{}
	swiftEnum.Swift = swift
	swiftEnum.Enum = e

	path := writer.AssembleFilePath(swiftEnum.Ctx.Output, swiftEnum.Name()+".swift")

	writer.WriteFile(path, ctx.EmunTemplate, ctx.EnumTemplateName, swiftEnum)
}

func (sc *SwiftCompiler) genStructCode(ctx *context.Context, s *parser.Struct) {

	swift.Ctx = ctx

	swiftStruct := &SwiftStruct{}
	swiftStruct.Swift = swift
	swiftStruct.Struct = s

	path := writer.AssembleFilePath(swiftStruct.Ctx.Output, swiftStruct.Name()+".swift")

	writer.WriteFile(path, ctx.StructTemplate, ctx.StructTemplateName, swiftStruct)
}

func (sc *SwiftCompiler) genServiceCode(ctx *context.Context, s *parser.Service) {

	swift.Ctx = ctx

	swiftService := &SwiftService{}
	swiftService.Swift = swift
	swiftService.Service = s

	path := writer.AssembleFilePath(swiftService.Ctx.Output, swiftService.Name()+".swift")

	writer.WriteFile(path, ctx.ServiceTemplate, ctx.ServiceTemplateName, swiftService)
}

// TypeString Type string
func (swift *Swift) TypeString(t *parser.Type) string {

	if t == nil {
		return swiftVoid
	}

	switch t.Name {
	case global.List:
		switch t.ValueType.Name {
		case global.List, global.Set, global.Map:
			panic("Unsupported inner container type.")
		}
		return "[" + swift.TypeString(t.ValueType) + "]"
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
		_thrift = swift.Ctx.Thrift
		_type = typeComponents[0]
	case 2:
		if key := swift.Ctx.Thrift.Includes[typeComponents[0]]; key != "" {
			_thrift = swift.Ctx.Thrifts[key]
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
