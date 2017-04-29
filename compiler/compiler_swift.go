package compiler

import (
	"strings"
	"sync"

	"github.com/samuel/go-thrift/parser"

	"github.com/Arror/CodeFly/context"
	"github.com/Arror/CodeFly/types"
)

const (
	enumTplName    = "SwiftEnum"
	structTplName  = "SwiftStruct"
	serviceTplName = "SwiftService"

	enumTplPath    = "templates/swift/enum.tpl"
	structTplPath  = "templates/swift/struct.tpl"
	serviceTplPath = "templates/swift/service.tpl"
)

func init() {
	register(&swiftcompiler{}, "swift")
}

var (
	_ctx *context.Context
)

type swiftcompiler struct{}

type assistant struct{}

// SwiftEnum swift Enum
type SwiftEnum struct {
	*parser.Enum
	Ass assistant
}

// SwiftStruct swift Struct
type SwiftStruct struct {
	*parser.Struct
	Ass assistant
}

// SwiftService swift Service
type SwiftService struct {
	*parser.Service
	Ass assistant
}

// Name enum name
func (se *SwiftEnum) Name() string {
	return _ctx.Thrift.Namespaces[_ctx.Lang] + se.Enum.Name
}

// Name struct name
func (ss *SwiftStruct) Name() string {
	return _ctx.Thrift.Namespaces[_ctx.Lang] + ss.Struct.Name
}

// Name service name
func (ss *SwiftService) Name() string {
	return ss.Service.Name + "Service"
}

// MethodName method Name
func (ss *SwiftService) MethodName(m *parser.Method) string {
	return strings.ToLower(m.Name[:1]) + m.Name[1:]
}

func (sc *swiftcompiler) compile(ctx *context.Context) {

	_ctx = ctx

	wg := sync.WaitGroup{}

	wg.Add(1)
	go func() {
		defer wg.Done()
		for _, e := range ctx.Thrift.Enums {
			se := &SwiftEnum{
				Enum: e,
				Ass:  assistant{},
			}
			fn := se.Name() + ".swift"
			err := ctx.ExportFile(fn, enumTplName, enumTplPath, se)
			if err != nil {
				panic(err.Error())
			}
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		for _, s := range ctx.Thrift.Structs {
			ss := &SwiftStruct{
				Struct: s,
				Ass:    assistant{},
			}
			fn := ss.Name() + ".swift"
			err := ctx.ExportFile(fn, structTplName, structTplPath, ss)
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
				Ass:     assistant{},
			}
			fn := ss.Name() + ".swift"
			err := ctx.ExportFile(fn, serviceTplName, serviceTplPath, ss)
			if err != nil {
				panic(err.Error())
			}
		}
	}()

	wg.Wait()
}

// FormatFiledName format filed name
func (ass assistant) FormatFiledName(n string) string {

	name := n

	if !strings.Contains(name, "_") {
		return name
	}

	components := strings.Split(name, "_")

	name = ""

	for _, component := range components {
		if component != "" {
			name += (strings.ToUpper(component[:1]) + component[1:])
		}
	}

	if name == "" {
		panic("invaild filed name: " + n)
	}

	return strings.ToLower(name[:1]) + name[1:]
}

// TypeString type string
func (ass assistant) TypeString(t *parser.Type) string {

	if t == nil {
		return swiftVoid
	}

	switch t.Name {
	case types.ThriftList:
		switch t.ValueType.Name {
		case types.ThriftList, types.ThriftSet, types.ThriftMap:
			panic("unsupported [[Type]]], [Key : Value] or Set<Type>")
		}
		return "[" + ass.TypeString(t.ValueType) + "]"
	case types.ThriftMap, types.ThriftSet:
		panic("unsupported [Key : Value] or Set<Type>")
	}

	if base := mapping[t.Name]; base != "" {
		return base
	}

	components := strings.Split(t.Name, ".")

	count := len(components)

	var _thrift *parser.Thrift
	var _type string

	switch count {
	case 1:
		_thrift = _ctx.Thrift
		_type = components[0]
	case 2:
		key := _ctx.Thrift.Includes[components[0]]
		_thrift = _ctx.Thrifts[key]
		_type = components[1]
	}

	if _thrift == nil || _type == "" {
		panic("unsupported type " + t.Name)
	}

	return _thrift.Namespaces[_ctx.Lang] + _type
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
	types.ThriftByte:   types.Unsupported,
	types.ThriftBinary: types.Unsupported,
}
