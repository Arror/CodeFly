package context

import (
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

func initTemplate(name string, path string) (*template.Template, error) {

	buffer, err := templates.Asset(path)

	if err != nil {
		return nil, err
	}

	return template.New(name).Parse(string(buffer))
}

// ExportFiles Export files
func (ctx Context) ExportFiles(fn string, tplname string, tplPath string, data interface{}) error {

	fp, err := filepath.Abs(filepath.Join(ctx.Output, fn))
	if err != nil {
		return err
	}

	file, err := os.OpenFile(fp, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}

	defer file.Close()

	t, err := initTemplate(tplname, tplPath)
	if err != nil {
		return err
	}

	err = t.ExecuteTemplate(file, tplname, data)
	if err != nil {
		return err
	}

	return nil
}
