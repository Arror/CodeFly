package context

import (
	"os"
	"path/filepath"
	"text/template"

	"github.com/Arror/CodeFly/templates"
	"github.com/samuel/go-thrift/parser"
)

// CommandArgs command args
type CommandArgs struct {
	Lang   string
	Input  string
	Output string
}

// Context Generate context
type Context struct {
	Args *CommandArgs

	Thrift *parser.Thrift

	Thrifts map[string]*parser.Thrift
}

// Create create context
func Create(lang string, input string, output string, thrifts map[string]*parser.Thrift) *Context {

	return &Context{
		Args: &CommandArgs{
			Lang:   lang,
			Input:  input,
			Output: output,
		},
		Thrifts: thrifts,
		Thrift:  thrifts[input],
	}
}

func initTemplate(name string, path string) (*template.Template, error) {

	buffer, err := templates.Asset(path)

	if err != nil {
		return nil, err
	}

	return template.New(name).Parse(string(buffer))
}

// GenerateFile generate file
func (ctx *Context) GenerateFile(fn string, tplName string, tplPath string, data interface{}) error {

	fp, err := filepath.Abs(filepath.Join(ctx.Args.Output, fn))
	if err != nil {
		return err
	}

	file, err := os.OpenFile(fp, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}

	defer file.Close()

	tpl, err := initTemplate(tplName, tplPath)
	if err != nil {
		return err
	}

	err = tpl.ExecuteTemplate(file, tplName, data)
	if err != nil {
		return err
	}

	return nil
}
