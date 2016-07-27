package generator

import (
	"fmt"
	"os"
	"path/filepath"
	"sync"

	"CodeFly/parser"
	"CodeFly/template"
)

func generatSwiftCode(stc *parser.SwiftThriftComponents) {
	op := stc.OutputPath
	if err := os.MkdirAll(op, 0755); err != nil {
		panic(fmt.Sprintf("无法创建文件夹:%s", op))
	}

	enumTmplName := template.SwiftEnumTemplateName
	structTmplName := template.SwiftStructTemplateName
	// serviceTmplName := template.SwiftServiceTemplateName

	enumTmpl := initTemplate(enumTmplName, template.EnumTemplate)
	structTmpl := initTemplate(structTmplName, template.StructTemplate)
	// serviceTmpl := initTemplate(serviceTmplName, template.ServiceTemplate)

	wg := sync.WaitGroup{}

	wg.Add(1)
	go func(op string) {
		defer wg.Done()
		enums := stc.SwiftThrift.Enums
		for _, e := range enums {
			name := fmt.Sprintf("%s.swift", e.Name)
			path, _ := filepath.Abs(filepath.Join(op, name))
			if err := outputFile(path, enumTmpl, enumTmplName, e); err != nil {
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
			es := &EmbenStruct{
				s,
			}
			if err := outputFile(path, structTmpl, structTmplName, es); err != nil {
				panic(err.Error())
			}
		}
	}(op)

	// wg.Add(1)
	// go func(op string) {
	// 	defer wg.Done()
	// 	services := stc.SwiftThrift.Services
	// 	for _, s := range services {
	// 		name := fmt.Sprintf("%s.swift", s.Name)
	// 		path, _ := filepath.Abs(filepath.Join(op, name))
	// 		if err := outputFile(path, serviceTmpl, serviceTmplName, s); err != nil {
	// 			panic(err.Error())
	// 		}
	// 	}
	// }(op)

	wg.Wait()
}
