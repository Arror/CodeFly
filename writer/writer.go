package writer

import (
	"log"
	"os"
	"path/filepath"
	"text/template"

	"CodeFly/templates"
)

// InitTemplate Init template
func InitTemplate(name string, path string) *template.Template {

	buffer := templates.MustAsset(path)

	template, err := template.New(name).Parse(string(buffer))
	if err != nil {
		log.Fatal(err.Error())
	}
	return template
}

// WriteFile Write file
func WriteFile(fp string, t *template.Template, tplname string, data interface{}) {

	file, err := os.OpenFile(fp, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer file.Close()
	if err := t.ExecuteTemplate(file, tplname, data); err != nil {
		log.Fatal(err.Error())
	}
}

// AssembleFilePath Assemble file path
func AssembleFilePath(op string, fn string) string {

	p, err := filepath.Abs(filepath.Join(op, fn))

	if err != nil {
		log.Fatalln(err.Error())
	}

	return p
}
