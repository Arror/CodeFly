package swift

import (
	"os"
	"sync"

	"CodeFly/lang/swift/tpl"
	"CodeFly/parameter"
	"CodeFly/writer"

	"github.com/samuel/go-thrift/parser"
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
func (gen *Generator) Generate(ts map[string]*parser.Thrift, param *parameter.Parameter) {

	op := param.Output
	if err := os.MkdirAll(op, 0755); err != nil {
		panic(err.Error())
	}

	gen.ts = ts
	gen.t = ts[param.Input]
	gen.lang = param.Lang
	gen.input = param.Input
	gen.output = param.Output

	wg := sync.WaitGroup{}

	wg.Add(1)
	go func() {
		defer wg.Done()

		enumTplName := tpl.SwiftEnumTplName
		enumTpl := writer.InitTemplate(enumTplName, tpl.SwiftEnumTpl)

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

		structTplName := tpl.SwiftStructTplName
		structTpl := writer.InitTemplate(structTplName, tpl.SwiftStructTpl)

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

		serviceTplName := tpl.SwiftServiceTpleName
		serviceTpl := writer.InitTemplate(serviceTplName, tpl.SwiftServiceTpl)

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
	return "PT -> " + f.Name
}

// DefaultValue default value
func (gen *Generator) DefaultValue(f *parser.Field) string {
	return "DV -> " + f.Name
}

// ValueFromJSON value from json string
func (gen *Generator) ValueFromJSON(f *parser.Field) string {
	return "VFJ -> " + f.Name
}

// ValueToJSON value to json string
func (gen *Generator) ValueToJSON(f *parser.Field) string {
	return "VTJ -> " + f.Name
}
