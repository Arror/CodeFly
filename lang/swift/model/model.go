package model

import (
	"os"
	"sync"

	"CodeFly/lang/swift/tpl"
	"CodeFly/parameter"
	"CodeFly/writer"

	"github.com/arrors/go-thrift/parser"
)

// SwiftGenerator Swift 代码生器
type SwiftGenerator struct{}

// Generate 遵守Generate协议
func (sg *SwiftGenerator) Generate(ts map[string]*parser.Thrift, param *parameter.Parameter) {

	op := param.Output
	if err := os.MkdirAll(op, 0755); err != nil {
		panic(err.Error())
	}

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
			se.Namespace = namespace

			name := se.Namespace + se.Enum.Name

			writer.WriteFile(writer.AssembleFilePath(op, name+".swift"), enumTpl, enumTplName, se)
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
			ss.Namespace = namespace
			ss.NamespaceMapping = namespaceMapping

			name := ss.Namespace + ss.Struct.Name

			writer.WriteFile(writer.AssembleFilePath(op, name+".swift"), structTpl, structTplName, ss)
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
	Namespaces
}

// Struct Swift结构类型
type Struct struct {
	*parser.Struct
	Namespaces
}

// Service Swift服务类型
type Service struct {
	*parser.Service
	Namespaces
}
