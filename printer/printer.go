package printer

import (
	"log"
	"os"
	"path/filepath"
	"sync"
	"text/template"

	"CodeFly/model"
	tpl "CodeFly/template"
	tps "CodeFly/types"

	"github.com/arrors/go-thrift/parser"
)

// Generate 代码生成
func Generate(ts map[string]*parser.Thrift, genInfo *model.GenerateCommandInfo) {

	op := genInfo.Output
	if err := os.MkdirAll(op, 0755); err != nil {
		panic(err.Error())
	}

	t := ts[genInfo.Input]

	wg := sync.WaitGroup{}

	enumTplName := tpl.SwiftEnumTplName
	enumTmpl := initTemplate(enumTplName, tpl.SwiftEnumTpl())

	wg.Add(1)
	go func() {
		defer wg.Done()

		for _, e := range t.Enums {

			se := &tps.SwiftEnum{}
			se.Enum = e
			se.Namespace = t.Namespaces[genInfo.Lang]

			name := se.Name()

			path, _ := filepath.Abs(filepath.Join(op, name+".swift"))
			printFile(path, enumTmpl, enumTplName, se)
		}
	}()

	structTplName := tpl.SwiftStructTemplateName
	structTmpl := initTemplate(structTplName, tpl.SwiftStructTpl())

	wg.Add(1)
	go func() {
		defer wg.Done()

		for _, s := range t.Structs {

			ss := &tps.SwiftStruct{}
			ss.Struct = s
			ss.Thrifts = ts
			ss.Thrift = t
			ss.Lang = genInfo.Lang
			ss.Namespace = t.Namespaces[genInfo.Lang]

			name := ss.Name()

			path, _ := filepath.Abs(filepath.Join(op, name+".swift"))
			printFile(path, structTmpl, structTplName, ss)
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
