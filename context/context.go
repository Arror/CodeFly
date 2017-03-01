package context

import (
	"log"
	"os"
	"path/filepath"
	"text/template"

	"github.com/Arror/CodeFly/templates"
	"github.com/samuel/go-thrift/parser"
)

// Context Generate context
type Context struct {
	Lang   string
	Input  string
	Output string

	Thrift *parser.Thrift

	Thrifts map[string]*parser.Thrift
}

// InitContext Context init
func InitContext(lang string, input string, output string, thrifts map[string]*parser.Thrift) Context {

	ctx := Context{}

	ctx.Lang = lang
	ctx.Input = input
	ctx.Output = output

	ctx.Thrifts = thrifts

	ctx.Thrift = thrifts[input]

	return ctx
}

func assembleFilePath(op string, fn string) string {

	p, err := filepath.Abs(filepath.Join(op, fn))

	if err != nil {
		log.Fatalln(err.Error())
	}

	return p
}

func initTemplate(name string, path string) *template.Template {

	buffer := templates.MustAsset(path)

	template, err := template.New(name).Parse(string(buffer))
	if err != nil {
		log.Fatal(err.Error())
	}
	return template
}

// MakeOutputFolder Make output folder
func (ctx Context) MakeOutputFolder() error {

	err := os.MkdirAll(ctx.Output, 0755)

	return err
}

// ExportFiles Export files
func (ctx Context) ExportFiles(fn string, tplname string, tplPath string, data interface{}) {

	fp := assembleFilePath(ctx.Output, fn)

	file, err := os.OpenFile(fp, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer file.Close()

	t := initTemplate(tplname, tplPath)

	if err := t.ExecuteTemplate(file, tplname, data); err != nil {
		log.Fatal(err.Error())
	}
}
