package printer

import (
	"html/template"
	"log"
	"os"
)

// InitTemplate 初始化模板文件
func InitTemplate(name string, tmpl string) *template.Template {

	template, err := template.New(name).Parse(tmpl)
	if err != nil {
		log.Fatal(err.Error())
	}
	return template
}

// PrintFile 输出文件
func PrintFile(fp string, t *template.Template, tplname string, data interface{}) {

	file, err := os.OpenFile(fp, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer file.Close()
	if err := t.ExecuteTemplate(file, tplname, data); err != nil {
		log.Fatal(err.Error())
	}
}
