package swift

import (
	"os"
	"sync"

	"CodeFly/lang/swift/tpl"
	"CodeFly/parameter"
	"CodeFly/writer"

	"github.com/arrors/go-thrift/parser"
)

// Generator Swift generator
type Generator struct {
	t      *parser.Thrift
	ts     map[string]*parser.Thrift
	lang   string
	input  string
	output string
}

// Generate 实现Generator协议
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

	t := ts[param.Input]

	namespaceMapping := make(map[string]string)

	for fn, fp := range t.Includes {
		for p, t := range ts {
			if p == fp {
				namespaceMapping[fn] = t.Namespaces[param.Lang]
				break
			}
		}
	}

	namespace := t.Namespaces[param.Lang]

	wg := sync.WaitGroup{}

	wg.Add(1)
	go func() {
		defer wg.Done()

		enumTplName := tpl.SwiftEnumTplName
		enumTpl := writer.InitTemplate(enumTplName, tpl.SwiftEnumTpl)

		for _, e := range t.Enums {

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

		for _, s := range t.Structs {

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

		for _, s := range t.Services {

			ss := &Service{}
			ss.Service = s
			ss.Namespace = namespace
			ss.NamespaceMapping = namespaceMapping

			name := s.Name + "Service"

			writer.WriteFile(writer.AssembleFilePath(op, name+".swift"), serviceTpl, serviceTplName, ss)
		}
	}()

	wg.Wait()
}

// Namespaces 名称空间信息
type Namespaces struct {
	Namespace        string
	NamespaceMapping map[string]string
}

// Enum Swift枚举类型
type Enum struct {
	*parser.Enum
	*Generator
}

// Struct Swift结构类型
type Struct struct {
	*parser.Struct
	*Generator
}

// Service Swift服务类型
type Service struct {
	*parser.Service
	*Generator
	Namespaces
}

// EnumName enum name
func (gen *Generator) EnumName(e *parser.Enum) string {
	return gen.t.Namespaces[gen.lang] + e.Name
}

// StructName struct name
func (gen *Generator) StructName(s *parser.Struct) string {
	return gen.t.Namespaces[gen.lang] + s.Name
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
