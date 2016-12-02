package generator

import (
	"fmt"
	"os"
	"path/filepath"
	"sync"

	"CodeFly/model"
	"CodeFly/printer"
	"CodeFly/template"
)

// GeneratSwiftCode Swift代码生成函数
func GeneratSwiftCode(stc *model.SwiftThriftComponents) {
	op := stc.OutputPath
	if err := os.MkdirAll(op, 0755); err != nil {
		panic(fmt.Sprintf("无法创建文件夹:%s", op))
	}

	enumTmplName := template.SwiftEnumTemplateName
	structTmplName := template.SwiftStructTemplateName
	serviceTmplName := template.SwiftServiceTemplateName

	enumTmpl := printer.InitTemplate(enumTmplName, template.EnumTemplate)
	structTmpl := printer.InitTemplate(structTmplName, template.StructTemplate)
	serviceTmpl := printer.InitTemplate(serviceTmplName, template.ServiceTemplate)

	wg := sync.WaitGroup{}

	wg.Add(1)
	go func(op string) {
		defer wg.Done()
		enums := stc.SwiftThrift.Enums
		for _, e := range enums {
			name := fmt.Sprintf("%s.swift", e.Name)
			path, _ := filepath.Abs(filepath.Join(op, name))
			if err := printer.PrintFile(path, enumTmpl, enumTmplName, e); err != nil {
				panic(err.Error())
			}
		}
	}(op)

	wg.Add(1)
	go func(op string) {
		defer wg.Done()
		structs := stc.SwiftThrift.Structs
		for _, s := range structs {
			name := fmt.Sprintf("%s.swift", s.Name)
			path, _ := filepath.Abs(filepath.Join(op, name))
			if err := printer.PrintFile(path, structTmpl, structTmplName, s); err != nil {
				panic(err.Error())
			}
		}
	}(op)

	wg.Add(1)
	go func(op string) {
		defer wg.Done()
		services := stc.SwiftThrift.Services
		for _, s := range services {
			name := fmt.Sprintf("%s.swift", s.Name)
			path, _ := filepath.Abs(filepath.Join(op, name))
			if err := printer.PrintFile(path, serviceTmpl, serviceTmplName, s); err != nil {
				panic(err.Error())
			}
		}
	}(op)

	wg.Wait()
}
