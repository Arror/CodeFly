package templates

import (
	"log"
	"text/template"
)

// InitTemplate Template init
func InitTemplate(name string, path string) *template.Template {

	buffer := MustAsset(path)

	template, err := template.New(name).Parse(string(buffer))
	if err != nil {
		log.Fatal(err.Error())
	}
	return template
}
