package generator

import (
	"os"
	"sync"

	"CodeFly/command"
	"CodeFly/lang/swift/model"
	"CodeFly/lang/swift/tpl"
	"CodeFly/writer"

	"github.com/arrors/go-thrift/parser"
)

// Generate 代码生成
func Generate(ts map[string]*parser.Thrift, cmd *command.Command) {

	op := cmd.Output
	if err := os.MkdirAll(op, 0755); err != nil {
		panic(err.Error())
	}

	t := ts[cmd.Input]

	namespaceMapping := make(map[string]string)

	for fn, fp := range t.Includes {
		for p, t := range ts {
			if p == fp {
				namespaceMapping[fn] = t.Namespaces[cmd.Lang]
				break
			}
		}
	}

	namespace := t.Namespaces[cmd.Lang]

	wg := sync.WaitGroup{}

	wg.Add(1)
	go func() {
		defer wg.Done()

		enumTplName := tpl.SwiftEnumTplName
		enumTpl := writer.InitTemplate(enumTplName, tpl.SwiftEnumTpl)

		for _, e := range t.Enums {

			se := &model.Enum{}
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

			ss := &model.Struct{}
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

			ss := &model.Service{}
			ss.Service = s
			ss.Namespace = namespace
			ss.NamespaceMapping = namespaceMapping

			name := s.Name + "Service"

			writer.WriteFile(writer.AssembleFilePath(op, name+".swift"), serviceTpl, serviceTplName, ss)
		}
	}()

	wg.Wait()
}
