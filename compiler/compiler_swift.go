package compiler

import (
	"os"
	"strings"
	"sync"

	"github.com/samuel/go-thrift/parser"

	"github.com/Arror/CodeFly/context"
	"github.com/Arror/CodeFly/types"
)

const (
	enumName    = "Enum"
	structName  = "Struct"
	serviceName = "Service"

	enumTplPath    = "templates/swift/enum.tpl"
	structTplPath  = "templates/swift/struct.tpl"
	serviceTplPath = "templates/swift/service.tpl"
)

var (
	enumTemplate    = initTemplate(enumName, enumTplPath)
	structTemplate  = initTemplate(structName, structTplPath)
	serviceTemplate = initTemplate(serviceName, serviceTplPath)
)

// SwiftCompiler Swift Code Compiler
type SwiftCompiler struct{}

// SwiftThriftAssistant Swift thrift assistant
type SwiftThriftAssistant struct {
	Ctx context.Context
}

// SwiftEnum Swift Enum
type SwiftEnum struct {
	*parser.Enum
	STA SwiftThriftAssistant
}

// SwiftStruct Swift Struct
type SwiftStruct struct {
	*parser.Struct
	STA SwiftThriftAssistant
}

// SwiftService Swift Service
type SwiftService struct {
	*parser.Service
	STA SwiftThriftAssistant
}

// Name Enum name
func (se *SwiftEnum) Name() string {
	return se.STA.Ctx.Thrift.Namespaces[se.STA.Ctx.Lang] + se.Enum.Name
}

// Name Struct name
func (ss *SwiftStruct) Name() string {
	return ss.STA.Ctx.Thrift.Namespaces[ss.STA.Ctx.Lang] + ss.Struct.Name
}

// Name Service name
func (ss *SwiftService) Name() string {
	return ss.Service.Name + "Service"
}

// MethodName Method Name
func (ss *SwiftService) MethodName(m *parser.Method) string {
	return strings.ToLower(m.Name[:1]) + m.Name[1:]
}

func (sc *SwiftCompiler) compile(ctx context.Context) {

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
				STA: SwiftThriftAssistant{
					Ctx: ctx,
				},
			}
			path := assembleFilePath(ctx.Output, se.Name()+".swift")
			exportFiles(path, enumTemplate, enumName, se)
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		for _, s := range ctx.Thrift.Structs {
			ss := &SwiftStruct{
				Struct: s,
				STA: SwiftThriftAssistant{
					Ctx: ctx,
				},
			}
			path := assembleFilePath(ctx.Output, ss.Name()+".swift")
			exportFiles(path, structTemplate, structName, ss)
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		for _, s := range ctx.Thrift.Services {
			ss := &SwiftService{
				Service: s,
				STA: SwiftThriftAssistant{
					Ctx: ctx,
				},
			}
			path := assembleFilePath(ctx.Output, ss.Name()+".swift")
			exportFiles(path, serviceTemplate, serviceName, ss)
		}
	}()

	wg.Wait()
}

// TypeString Type string
func (STA SwiftThriftAssistant) TypeString(t *parser.Type) string {

	if t == nil {
		return swiftVoid
	}

	switch t.Name {
	case types.ThriftList:
		switch t.ValueType.Name {
		case types.ThriftList, types.ThriftSet, types.ThriftMap:
			panic("Unsupported inner container type.")
		}
		return "[" + STA.TypeString(t.ValueType) + "]"
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
		_thrift = STA.Ctx.Thrift
		_type = typeComponents[0]
	case 2:
		if key := STA.Ctx.Thrift.Includes[typeComponents[0]]; key != "" {
			_thrift = STA.Ctx.Thrifts[key]
			_type = typeComponents[1]
		} else {
			panic(typeComponents[0] + ".thrift not find in file include.")
		}
	}

	if _thrift != nil && _type != "" {
		return _thrift.Namespaces[STA.Ctx.Lang] + _type
	}

	panic("Unsupported type.")
}

const (
	swiftInt    = "Int"
	swiftInt64  = "Int64"
	swiftDouble = "Double"
	swiftBool   = "Bool"
	swiftString = "String"
	swiftVoid   = "Void"
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
