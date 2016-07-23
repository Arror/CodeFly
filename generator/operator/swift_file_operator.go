package operator

import (
	"fmt"
	"os"
	"path/filepath"
	"sync"

	"CodeFly/generator/reader/readerSwift"
	"CodeFly/generator/templates"
	"CodeFly/generator/templates/templateSwift"
)

// RenderSwift 渲染Swift
func RenderSwift(str *readerSwift.SwiftThriftReader) {

	op := str.ThriftReader.OutputPath
	if err := os.MkdirAll(op, 0755); err != nil {
		panic(fmt.Sprintf("无法创建文件夹:%s", op))
	}

	enumTplName := templates.SwiftEnumTemplate
	structTplName := templates.SwiftStructTemplate
	servicelName := templates.SwiftServiceTemplate

	enumTpl := InitTemplate(enumTplName, templateSwift.SwiftEnumTemplate)
	structTpl := InitTemplate(structTplName, templateSwift.SwiftStructTemplate)
	serviceTpl := InitTemplate(servicelName, templateSwift.SwiftServiceTemplate)

	wg := sync.WaitGroup{}

	wg.Add(1)
	go func(op string) {
		defer wg.Done()
		enums := str.SwiftThriftMap.Enums
		for _, e := range enums {
			name := fmt.Sprintf("%s.swift", e.Name)
			path, _ := filepath.Abs(filepath.Join(op, name))
			if err := OutputFile(path, enumTpl, enumTplName, e); err != nil {
				panic(err.Error())
			}
		}
	}(op)

	wg.Add(1)
	go func(op string) {
		defer wg.Done()
		structs := str.SwiftThriftMap.Structs
		for _, s := range structs {
			name := fmt.Sprintf("%s.swift", s.Name)
			path, _ := filepath.Abs(filepath.Join(op, name))
			if err := OutputFile(path, structTpl, structTplName, s); err != nil {
				panic(err.Error())
			}
		}
	}(op)

	wg.Add(1)
	go func(op string) {
		defer wg.Done()
		services := str.SwiftThriftMap.Services
		for _, s := range services {
			name := fmt.Sprintf("%s.swift", s.Name)
			path, _ := filepath.Abs(filepath.Join(op, name))
			if err := OutputFile(path, serviceTpl, servicelName, s); err != nil {
				panic(err.Error())
			}
		}
	}(op)

	wg.Wait()
}
