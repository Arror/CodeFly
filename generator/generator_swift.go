package generator

import (
	"errors"
	"strings"
	"sync"

	"github.com/samuel/go-thrift/parser"

	"github.com/Arror/CodeFly/context"
	"github.com/Arror/CodeFly/types"
	"github.com/Arror/CodeFly/utils"
)

const (
	enumTplName    = "SwiftEnum"
	structTplName  = "SwiftStruct"
	serviceTplName = "SwiftService"

	enumTplPath    = "templates/swift/enum.tpl"
	structTplPath  = "templates/swift/struct.tpl"
	serviceTplPath = "templates/swift/service.tpl"

	swiftInt    = "Int"
	swiftInt64  = "Int64"
	swiftDouble = "Double"
	swiftBool   = "Bool"
	swiftString = "String"
	swiftVoid   = "Void"
)

var typeMapping = map[string]string{
	types.ThriftI16:    swiftInt,
	types.ThriftI32:    swiftInt,
	types.ThriftI64:    swiftInt64,
	types.ThriftBool:   swiftBool,
	types.ThriftDouble: swiftDouble,
	types.ThriftString: swiftString,
	types.ThriftByte:   types.Unsupported,
	types.ThriftBinary: types.Unsupported,
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

// Name enum name
func (se *SwiftEnum) Name() string {
	return se.contextwrapper.ctx.Thrift.Namespaces[se.contextwrapper.ctx.Args.Lang] + se.Enum.Name
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
		for _, e := range ctx.Thrift.Enums {
			se := &SwiftEnum{
				Enum: e,
				contextwrapper: &contextwrapper{
					ctx: ctx,
				},
			}
			fn := se.Name() + ".swift"
			err := ctx.GenerateFile(fn, enumTplName, enumTplPath, se)
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
func (ctxW contextwrapper) FormatedFiledName(n string) string {

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

func (ctxW *contextwrapper) DefaultValue(t string) string {

	switch t {
	case swiftInt, swiftInt64:
		return "0"
	case swiftDouble:
		return "0.0"
	case swiftBool:
		return "false"
	case swiftString:
		return "\"\""
	}

	if strings.HasPrefix(t, "[") && strings.HasSuffix(t, "]") {
		return "[]"
	}
	errMes := "type of " + t + " default value not implemented"
	panic(errMes)
}

// TypeString parse type string
func (ctxW *contextwrapper) TypeString(t *parser.Type) (string, error) {

	if t == nil {
		return swiftVoid, nil
	}

	switch t.Name {
	case types.ThriftList:
		switch t.ValueType.Name {
		case types.ThriftList, types.ThriftSet, types.ThriftMap:
			return "", errors.New("Array inner type [Array、Set、Map] not implemented")
		}
		innerType, err := ctxW.TypeString(t.ValueType)
		if err != nil {
			return "", err
		}
		return "[" + innerType + "]", nil
	case types.ThriftMap, types.ThriftSet:
		return "", errors.New("Type [Set、Map] are not implemented")
	}

	if baseType := typeMapping[t.Name]; baseType != "" {
		return baseType, nil
	}

	components := strings.Split(t.Name, ".")

	count := len(components)

	var _thrift *parser.Thrift
	var _type string

	switch count {
	case 1:
		_thrift = ctxW.ctx.Thrift
		_type = components[0]
	case 2:
		_thrift = ctxW.ctx.Thrifts[ctxW.ctx.Thrift.Includes[components[0]]]
		_type = components[1]
	}

	if _thrift == nil || _type == "" {
		return "", errors.New("")
	}

	typeStr := _thrift.Namespaces[ctxW.ctx.Args.Lang] + _type

	return typeStr, nil
}
