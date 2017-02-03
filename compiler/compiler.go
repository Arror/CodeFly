package compiler

import (
	"log"
	"os"
	"path/filepath"
	"text/template"

	"CodeFly/context"
	"CodeFly/global"
	"CodeFly/templates"
)

// Compiler Compiler interface
type Compiler interface {
	genCodes(ctx *context.Context)
}

var compilerMapping = map[string]Compiler{
	global.Swift: &SwiftCompiler{},
}

// GenCode Generate code
func GenCode(ctx *context.Context) {

	compiler := compilerMapping[ctx.Lang]

	compiler.genCodes(ctx)
}

func initTemplate(name string, path string) *template.Template {

	buffer := templates.MustAsset(path)

	template, err := template.New(name).Parse(string(buffer))
	if err != nil {
		log.Fatal(err.Error())
	}
	return template
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
