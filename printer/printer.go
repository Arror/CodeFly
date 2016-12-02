package printer

import (
	"html/template"
	"os"
)

// InitTemplate 初始化模板文件
func InitTemplate(name string, tmpl string) *template.Template {

	template, err := template.New(name).Parse(tmpl)
	if err != nil {
		panic(err.Error())
	}
	return template
}

// PrintFile 输出文件
func PrintFile(fp string, t *template.Template, tplname string, data interface{}) error {

	file, err := os.OpenFile(fp, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	defer file.Close()
	return t.ExecuteTemplate(file, tplname, data)
}
