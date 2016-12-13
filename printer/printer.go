package printer

import (
	"log"
	"os"
	"path/filepath"
	"sync"
	"text/template"

	"CodeFly/command"
	"CodeFly/lang/swift/model"
	"CodeFly/lang/swift/tpl"

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
		enumTpl := initTemplate(enumTplName, tpl.SwiftEnumTpl)

		for _, e := range t.Enums {

			se := &model.Enum{}
			se.Enum = e
			se.Namespace = namespace

			name := se.Namespace + se.Enum.Name

			printFile(assembleFilePath(op, name+".swift"), enumTpl, enumTplName, se)
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()

		structTplName := tpl.SwiftStructTplName
		structTpl := initTemplate(structTplName, tpl.SwiftStructTpl)

		for _, s := range t.Structs {

			ss := &model.Struct{}
			ss.Struct = s
			ss.Namespace = namespace
			ss.NamespaceMapping = namespaceMapping

			name := ss.Namespace + ss.Struct.Name

			printFile(assembleFilePath(op, name+".swift"), structTpl, structTplName, ss)
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()

		serviceTplName := tpl.SwiftServiceTpleName
		serviceTpl := initTemplate(serviceTplName, tpl.SwiftServiceTpl)

		for _, s := range t.Services {

			ss := &model.Service{}
			ss.Service = s
			ss.Namespace = namespace
			ss.NamespaceMapping = namespaceMapping

			name := s.Name + "Service"

			printFile(assembleFilePath(op, name+".swift"), serviceTpl, serviceTplName, ss)
		}
	}()

	wg.Wait()
}

func initTemplate(name string, tmpl string) *template.Template {

	template, err := template.New(name).Parse(tmpl)
	if err != nil {
		log.Fatal(err.Error())
	}
	return template
}

func printFile(fp string, t *template.Template, tplname string, data interface{}) {

	file, err := os.OpenFile(fp, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer file.Close()
	if err := t.ExecuteTemplate(file, tplname, data); err != nil {
		log.Fatal(err.Error())
	}
}

func assembleFilePath(op string, fn string) string {

	p, err := filepath.Abs(filepath.Join(op, fn))

	if err != nil {
		log.Fatalln(err.Error())
	}

	return p
}
