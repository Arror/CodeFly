package printer

import (
	"html/template"
	"log"
	"os"
	"path/filepath"
	"sync"

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

	enumTplName := tpl.SwiftEnumTplName
	enumTmpl := initTemplate(enumTplName, tpl.SwiftEnumTpl())

	wg := sync.WaitGroup{}

	wg.Add(1)
	go func() {
		defer wg.Done()

		for _, enum := range t.Enums {

			e := &tps.SwiftEnum{}
			e.Enum = enum
			e.Namespace = t.Namespaces[genInfo.Lang]

			name := e.Name()

			path, _ := filepath.Abs(filepath.Join(op, name+".swift"))
			printFile(path, enumTmpl, enumTplName, e)
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

// func GeneratSwiftCode(stc *model.SwiftThriftComponents) {
// 	op := stc.OutputPath
// 	if err := os.MkdirAll(op, 0755); err != nil {
// 		panic(fmt.Sprintf("无法创建文件夹:%s", op))
// 	}

// 	structTmplName := template.SwiftStructTemplateName
// 	serviceTmplName := template.SwiftServiceTemplateName

// 	structTmpl := printer.InitTemplate(structTmplName, template.StructTemplate)
// 	serviceTmpl := printer.InitTemplate(serviceTmplName, template.ServiceTemplate)

// 	wg := sync.WaitGroup{}

// 	wg.Add(1)
// 	go func(op string) {
// 		defer wg.Done()
// 		structs := stc.SwiftThrift.Structs
// 		for _, s := range structs {
// 			name := fmt.Sprintf("%s.swift", s.Name)
// 			path, _ := filepath.Abs(filepath.Join(op, name))
// 			printer.PrintFile(path, structTmpl, structTmplName, s)
// 		}
// 	}(op)

// 	wg.Add(1)
// 	go func(op string) {
// 		defer wg.Done()
// 		services := stc.SwiftThrift.Services
// 		for _, s := range services {
// 			name := fmt.Sprintf("%s.swift", s.Name)
// 			path, _ := filepath.Abs(filepath.Join(op, name))
// 			printer.PrintFile(path, serviceTmpl, serviceTmplName, s)
// 		}
// 	}(op)

// 	wg.Wait()
// }
