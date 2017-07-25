package generator

import (
	"strings"
	"sync"

	"github.com/samuel/go-thrift/parser"

	"github.com/Arror/CodeFly/context"
	"github.com/Arror/CodeFly/types"
	"github.com/Arror/CodeFly/utils"
)

const (
	structTplName  = "SwiftStruct"
	serviceTplName = "SwiftService"

	structTplPath  = "templates/swift/struct.tpl"
	serviceTplPath = "templates/swift/service.tpl"

	sInt     = "Int"
	sInt64   = "Int64"
	sDouble  = "Double"
	sBool    = "Bool"
	sString  = "String"
	sVoid    = "Void"
	sUnknown = "Unknown"
)

var typeMapping = map[string]string{
	types.ThriftI16:    sInt,
	types.ThriftI32:    sInt,
	types.ThriftI64:    sInt64,
	types.ThriftBool:   sBool,
	types.ThriftDouble: sDouble,
	types.ThriftString: sString,
	types.ThriftByte:   sUnknown,
	types.ThriftBinary: sUnknown,
}

func init() {
	enroll(&swiftgenerator{}, "swift")
}

type swiftgenerator struct{}

type contextwrapper struct {
	ctx *context.Context
}

// SwiftEnum swift Enum
type SwiftEnum struct {
	*parser.Enum
	*contextwrapper
}

// SwiftStruct swift Struct
type SwiftStruct struct {
	*parser.Struct
	*contextwrapper
}

// Name struct name
func (ss *SwiftStruct) Name() string {
	return ss.contextwrapper.ctx.Thrift.Namespaces[ss.contextwrapper.ctx.Args.Lang] + ss.Struct.Name
}

// SwiftService swift Service
type SwiftService struct {
	*parser.Service
	*contextwrapper
}

// Name service name
func (ss *SwiftService) Name() string {
	return ss.Service.Name + "Service"
}

// MethodName method Name
func (ss *SwiftService) MethodName(m *parser.Method) string {
	return strings.ToLower(m.Name[:1]) + m.Name[1:]
}

func (sc *swiftgenerator) generate(ctx *context.Context) {

	wg := sync.WaitGroup{}

	wg.Add(1)
	go func() {
		defer wg.Done()
		for _, s := range ctx.Thrift.Structs {
			ss := &SwiftStruct{
				Struct: s,
				contextwrapper: &contextwrapper{
					ctx: ctx,
				},
			}
			fn := ss.Name() + ".swift"
			err := ctx.GenerateFile(fn, structTplName, structTplPath, ss)
			if err != nil {
				panic(err.Error())
			}
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		for _, s := range ctx.Thrift.Services {
			ss := &SwiftService{
				Service: s,
				contextwrapper: &contextwrapper{
					ctx: ctx,
				},
			}
			fn := ss.Name() + ".swift"
			err := ctx.GenerateFile(fn, serviceTplName, serviceTplPath, ss)
			if err != nil {
				panic(err.Error())
			}
		}
	}()

	wg.Wait()
}

// FormatedFiledName format filed name
func (ctxW *contextwrapper) FormatedFiledName(n string) string {

	if !strings.Contains(n, "_") {
		return n
	}

	components := strings.Split(n, "_")

	components = utils.Filter(components, func(str string) bool {
		return str == ""
	})

	if components == nil || len(components) == 0 {
		panic("invaild filed name: " + n)
	}

	name := ""

	for idx, component := range components {

		if idx == 0 {
			name += component
			continue
		}

		if component != "" {
			name += (strings.Title(component))
		}
	}

	return name
}

// Result parse type result
type Result struct {
	Type    string
	Default string
}

// GetPath get method path
func (ctxW *contextwrapper) GetPath(m *parser.Method) string {

	if m == nil {
		panic("Mehtod nil...")
	}

	for _, a := range m.Annotations {

		if strings.ToLower(a.Name) == "path" {
			return a.Value
		}
	}

	panic("Mehtod path not define")
}

func (ctxW *contextwrapper) IsBaseType(t string) bool {

	switch t {
	case sInt, sInt64, sDouble, sBool, sString:
		return true
	default:
		return false
	}
}

func (ctxW *contextwrapper) ParserType(t *parser.Type) *Result {

	if t == nil {
		return &Result{
			Type:    sVoid,
			Default: "",
		}
	}

	switch t.Name {
	case types.ThriftList:
		switch t.ValueType.Name {
		case types.ThriftList:
			panic("Array not implement for Array element")
		case types.ThriftMap:
			panic("Dictionary not implement for Array element")
		case types.ThriftSet:
			panic("Set not implement for Array element")
		case types.ThriftByte:
			panic("Byte not implement for Array element")
		case types.ThriftBinary:
			panic("Binary not implement for Array element")
		}
		result := ctxW.ParserType(t.ValueType)
		return &Result{
			Type:    "[" + result.Type + "]",
			Default: "[]",
		}
	case types.ThriftMap:
		panic("Dictionary not implement")
	case types.ThriftSet:
		panic("Set not implement")
	case types.ThriftByte:
		panic("Byte not implement")
	case types.ThriftBinary:
		panic("Binary not implement")
	case types.ThriftI16, types.ThriftI32:
		return &Result{
			Type:    sInt,
			Default: "0",
		}
	case types.ThriftI64:
		return &Result{
			Type:    sInt64,
			Default: "0",
		}
	case types.ThriftDouble:
		return &Result{
			Type:    sDouble,
			Default: "0.0",
		}
	case types.ThriftBool:
		return &Result{
			Type:    sBool,
			Default: "false",
		}
	case types.ThriftString:
		return &Result{
			Type:    sString,
			Default: "\"\"",
		}
	}

	if _thrift, _type := func() (*parser.Thrift, string) {
		components := strings.Split(t.Name, ".")
		switch len(components) {
		case 1:
			return ctxW.ctx.Thrift, components[0]
		case 2:
			return ctxW.ctx.Thrifts[ctxW.ctx.Thrift.Includes[components[0]]], components[1]
		}
		return nil, ""
	}(); _thrift != nil && _type != "" {

		for _, s := range _thrift.Structs {
			if s.Name == _type {
				structName := _thrift.Namespaces[ctxW.ctx.Args.Lang] + _type
				return &Result{
					Type:    structName,
					Default: structName + "()",
				}
			}
		}
	}

	panic("Undefine error, info: " + t.Name)
}
