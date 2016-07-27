package generator

import (
	"fmt"
	"os"
	"text/template"

	"CodeFly/parser"

	p "github.com/samuel/go-thrift/parser"
)

// Distributor 任务分配者
func Distributor(ts map[string]*p.Thrift, lang string, ip string, op string) error {

	switch lang {
	case parser.Swift:
		stc := parser.Parser(ts, ip, op)
		generatSwiftCode(stc)
	default:
		return fmt.Errorf("未被支持的语言")
	}
	return nil
}

func initTemplate(name string, tmpl string) *template.Template {

	template, err := template.New(name).Parse(tmpl)
	if err != nil {
		panic(err.Error())
	}
	return template
}

func outputFile(fp string, t *template.Template, tplname string, data interface{}) error {

	file, err := os.OpenFile(fp, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	defer file.Close()
	return t.ExecuteTemplate(file, tplname, data)
}
