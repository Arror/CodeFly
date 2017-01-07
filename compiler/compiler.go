package compiler

import (
	"log"
	"os"
	"path/filepath"
	"sync"
	"text/template"

	"github.com/samuel/go-thrift/parser"

	"CodeFly/context"
	"CodeFly/global"
)

// Compiler Compiler interface
type Compiler interface {
	genEnumCode(ctx *context.Context, e *parser.Enum)
	genStructCode(ctx *context.Context, s *parser.Struct)
	genServiceCode(ctx *context.Context, s *parser.Service)
}

var compilerMapping = map[string]Compiler{
	global.Swift: &SwiftCompiler{},
}

// GenCode Generate code
func GenCode(ctx *context.Context) {

	compiler := compilerMapping[ctx.Lang]

	if err := os.MkdirAll(ctx.Output, 0755); err != nil {
		panic(err.Error())
	}

	wg := sync.WaitGroup{}

	wg.Add(1)
	go func() {
		defer wg.Done()
		for _, e := range ctx.Thrift.Enums {
			compiler.genEnumCode(ctx, e)
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		for _, s := range ctx.Thrift.Structs {
			compiler.genStructCode(ctx, s)
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		for _, s := range ctx.Thrift.Services {
			compiler.genServiceCode(ctx, s)
		}
	}()

	wg.Wait()
}

func writeFile(fp string, t *template.Template, tplname string, data interface{}) {

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
