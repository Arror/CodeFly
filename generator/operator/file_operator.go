package operator

import (
	"fmt"
	"os"
	"text/template"

	"CodeFly/generator/reader"
	"CodeFly/generator/reader/readerSwift"
	"CodeFly/global"
	"CodeFly/validity"
)

// Print 输出文件
func Print(r *reader.ThriftReader, gci *validity.GenerateCommandInfo) error {

	switch gci.Lang {
	case global.Swift:
		str := &readerSwift.SwiftThriftReader{}
		str.InitSwiftThrift(r)
		RenderSwift(str)
		return nil
	default:
		return fmt.Errorf("不支持的语言")
	}
}

// InitTemplate 模板初始化
func InitTemplate(name string, tmpl string) *template.Template {

	template, err := template.New(name).Parse(tmpl)
	if err != nil {
		panic(err.Error())
	}
	return template
}

// OutputFile 文件输出
func OutputFile(fp string, t *template.Template, tplname string, data interface{}) error {

	file, err := os.OpenFile(fp, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	defer file.Close()
	return t.ExecuteTemplate(file, tplname, data)
}
