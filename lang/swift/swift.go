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

// Generator Swift generator
type Generator struct {
	t      *parser.Thrift
	ts     map[string]*parser.Thrift
	lang   string
	input  string
	output string
}

// Generate Generate implement Generator interface
func (gen *Generator) Generate() {

	if err := os.MkdirAll(global.Output, 0755); err != nil {
		panic(err.Error())
	}

	gen.ts = global.ThriftMapping
	gen.t = global.ThriftMapping[global.Input]
	gen.lang = global.Lang
	gen.input = global.Input
	gen.output = global.Output

	wg := sync.WaitGroup{}

	wg.Add(1)
	go func() {
		defer wg.Done()

		enumTpl := writer.InitTemplate(enumTplName, enumTplPath)

		for _, e := range gen.t.Enums {

			se := &Enum{}
			se.Enum = e
			se.Generator = gen

			writer.WriteFile(writer.AssembleFilePath(gen.output, gen.EnumName(e)+".swift"), enumTpl, enumTplName, se)
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()

		structTpl := writer.InitTemplate(structTplName, structTplPath)

		for _, s := range gen.t.Structs {

			ss := &Struct{}
			ss.Struct = s
			ss.Generator = gen

			writer.WriteFile(writer.AssembleFilePath(gen.output, gen.StructName(s)+".swift"), structTpl, structTplName, ss)
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()

		serviceTpl := writer.InitTemplate(serviceTplName, serviceTplPath)

		for _, s := range gen.t.Services {

			ss := &Service{}
			ss.Service = s
			ss.Generator = gen

			writer.WriteFile(writer.AssembleFilePath(gen.output, gen.ServiceName(s)+".swift"), serviceTpl, serviceTplName, ss)
		}
	}()

	wg.Wait()
}

// Enum Swift enum type
type Enum struct {
	*parser.Enum
	*Generator
}

// Struct Swift struct type
type Struct struct {
	*parser.Struct
	*Generator
}

// Service Swift service type
type Service struct {
	*parser.Service
	*Generator
}

// EnumName enum name
func (gen *Generator) EnumName(e *parser.Enum) string {
	return gen.t.Namespaces[gen.lang] + e.Name
}

// StructName struct name
func (gen *Generator) StructName(s *parser.Struct) string {
	return gen.t.Namespaces[gen.lang] + s.Name
}

// ServiceName service name
func (gen *Generator) ServiceName(s *parser.Service) string {
	return s.Name + "Service"
}

// PropertyType property type
func (gen *Generator) PropertyType(f *parser.Field) string {
	return gen.typeString(f.Type)
}

func (gen *Generator) typeString(t *parser.Type) string {

	switch t.Name {
	case global.List:
		return "[" + gen.typeString(t.ValueType) + "]"
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
		_thrift = gen.t
		_type = t.Name
	case 2:
		if key := gen.t.Includes[typeComponents[0]]; key != "" {
			_thrift = gen.ts[key]
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
