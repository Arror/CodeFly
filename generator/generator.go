package generator

import (
	"fmt"
	"os"
	"path/filepath"
	"sync"

	"CodeFly/model"
	"CodeFly/printer"
	"CodeFly/template"
	tp "CodeFly/types"

	"github.com/arrors/go-thrift/parser"
)

// GeneratSwiftCode Swift代码生成函数
func GeneratSwiftCode(stc *model.SwiftThriftComponents) {
	op := stc.OutputPath
	if err := os.MkdirAll(op, 0755); err != nil {
		panic(fmt.Sprintf("无法创建文件夹:%s", op))
	}

	structTmplName := template.SwiftStructTemplateName
	serviceTmplName := template.SwiftServiceTemplateName

	structTmpl := printer.InitTemplate(structTmplName, template.StructTemplate)
	serviceTmpl := printer.InitTemplate(serviceTmplName, template.ServiceTemplate)

	wg := sync.WaitGroup{}

	wg.Add(1)
	go func(op string) {
		defer wg.Done()
		structs := stc.SwiftThrift.Structs
		for _, s := range structs {
			name := fmt.Sprintf("%s.swift", s.Name)
			path, _ := filepath.Abs(filepath.Join(op, name))
			printer.PrintFile(path, structTmpl, structTmplName, s)
		}
	}(op)

	wg.Add(1)
	go func(op string) {
		defer wg.Done()
		services := stc.SwiftThrift.Services
		for _, s := range services {
			name := fmt.Sprintf("%s.swift", s.Name)
			path, _ := filepath.Abs(filepath.Join(op, name))
			printer.PrintFile(path, serviceTmpl, serviceTmplName, s)
		}
	}(op)

	wg.Wait()
}

// Generate 代码生成
func Generate(ts map[string]*parser.Thrift, genInfo *model.GenerateCommandInfo) {

	op := genInfo.Output
	if err := os.MkdirAll(op, 0755); err != nil {
		panic(err.Error())
	}

	t := ts[genInfo.Input]

	enumTplName := template.SwiftEnumTplName
	enumTmpl := printer.InitTemplate(enumTplName, template.EnumTpl)

	wg := sync.WaitGroup{}

	wg.Add(1)
	go func() {
		defer wg.Done()

		for _, enum := range t.Enums {

			e := &tp.SwiftEnum{}
			e.Enum = enum
			e.Namespace = t.Namespaces[genInfo.Lang]

			name := e.Name()

			path, _ := filepath.Abs(filepath.Join(op, name+".swift"))
			printer.PrintFile(path, enumTmpl, enumTplName, e)
		}
	}()

	wg.Wait()
}
