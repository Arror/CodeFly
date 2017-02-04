package compiler

import (
	"os"
	"strings"
	"sync"

	"github.com/samuel/go-thrift/parser"

	"CodeFly/context"
	"CodeFly/types"
)

var sCtx context.SwiftContext

// SwiftCompiler Swift Code Compiler
type SwiftCompiler struct{}

// SwiftEnum Swift Enum
type SwiftEnum struct {
	*parser.Enum
}

// SwiftStruct Swift Struct
type SwiftStruct struct {
	*parser.Struct
}

// SwiftService Swift Service
type SwiftService struct {
	*parser.Service
}

// SwiftCommon Swift common protocol
type SwiftCommon interface {
	Name() string
	TypeString(t *parser.Type) string
	ExportSwiftFiles(ctx context.SwiftContext)
}

// Name Enum name
func (se *SwiftEnum) Name() string {
	return sCtx.Thrift.Namespaces[sCtx.Lang] + se.Enum.Name
}

// Name Struct name
func (ss *SwiftStruct) Name() string {
	return sCtx.Thrift.Namespaces[sCtx.Lang] + ss.Struct.Name
}

// Name Service name
func (ss *SwiftService) Name() string {
	return ss.Service.Name + "Service"
}

// TypeString Enum type string
func (se *SwiftEnum) TypeString(t *parser.Type) string {
	return typeString(t)
}

// TypeString Struct type string
func (ss *SwiftStruct) TypeString(t *parser.Type) string {
	return typeString(t)
}

// TypeString Struct type string
func (ss *SwiftService) TypeString(t *parser.Type) string {
	return typeString(t)
}

// ExportSwiftFiles Export swift enum files
func (se *SwiftEnum) ExportSwiftFiles(ctx context.SwiftContext) {
	path := assembleFilePath(ctx.Output, se.Name())
	exportFiles(path, ctx.EnumTemplate, ctx.EnumTemplateName, se)
}

// ExportSwiftFiles Export swift struct files
func (ss *SwiftStruct) ExportSwiftFiles(ctx context.SwiftContext) {
	path := assembleFilePath(ctx.Output, ss.Name())
	exportFiles(path, ctx.StructTemplate, ctx.StructTemplateName, ss)
}

// ExportSwiftFiles Export swift service files
func (ss *SwiftService) ExportSwiftFiles(ctx context.SwiftContext) {
	path := assembleFilePath(ctx.Output, ss.Name())
	exportFiles(path, ctx.ServiceTemplate, ctx.ServiceTemplateName, ss)
}

// MethodName Method Name
func (ss *SwiftService) MethodName(m *parser.Method) string {
	return strings.ToLower(m.Name[:1]) + m.Name[1:]
}

func (sc *SwiftCompiler) compile(ctx context.Context) {

	sCtx = context.InitSwiftContext(ctx)

	if err := os.MkdirAll(ctx.Output, 0755); err != nil {
		panic(err.Error())
	}

	wg := sync.WaitGroup{}

	wg.Add(1)
	go func() {
		defer wg.Done()
		for _, e := range ctx.Thrift.Enums {
			se := &SwiftEnum{
				Enum: e,
			}
			se.ExportSwiftFiles(sCtx)
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		for _, s := range ctx.Thrift.Structs {
			ss := &SwiftStruct{
				Struct: s,
			}
			ss.ExportSwiftFiles(sCtx)
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		for _, s := range ctx.Thrift.Services {
			ss := &SwiftService{
				Service: s,
			}
			ss.ExportSwiftFiles(sCtx)
		}
	}()

	wg.Wait()
}

func typeString(t *parser.Type) string {

	if t == nil {
		return swiftVoid
	}

	switch t.Name {
	case types.ThriftList:
		switch t.ValueType.Name {
		case types.ThriftList, types.ThriftSet, types.ThriftMap:
			panic("Unsupported inner container type.")
		}
		return "[" + typeString(t.ValueType) + "]"
	case types.ThriftMap, types.ThriftSet:
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
		_thrift = sCtx.Thrift
		_type = typeComponents[0]
	case 2:
		if key := sCtx.Thrift.Includes[typeComponents[0]]; key != "" {
			_thrift = sCtx.Thrifts[key]
			_type = typeComponents[1]
		} else {
			panic(typeComponents[0] + ".thrift not find in file include.")
		}
	}

	if _thrift != nil && _type != "" {
		return _thrift.Namespaces[sCtx.Lang] + _type
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
	types.ThriftI16:    swiftInt,
	types.ThriftI32:    swiftInt,
	types.ThriftI64:    swiftInt64,
	types.ThriftBool:   swiftBool,
	types.ThriftDouble: swiftDouble,
	types.ThriftString: swiftString,
	types.ThriftByte:   types.UnsupportedType,
	types.ThriftBinary: types.UnsupportedType,
}
