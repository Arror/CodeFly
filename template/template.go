package template

import (
	"io/ioutil"
	"log"
)

const (
	SwiftEnumTplName     = "EnumTemplate"
	SwiftStructTplName   = "StructTemplate"
	SwiftServiceTpleName = "ServiceTemplate"
)

const (
	enumTplPath    = "./template/enum.tpl"
	structTplPath  = "./template/struct.tpl"
	serviceTplPath = "./template/service.tpl"
)

func fileToString(fp string) string {

	buffer, err := ioutil.ReadFile(fp)

	if err != nil {
		log.Fatalln(err.Error())
	}

	return string(buffer)
}

func SwiftEnumTpl() string {
	return fileToString(enumTplPath)
}

func SwiftStructTpl() string {
	return fileToString(structTplPath)
}

func SwiftServiceTpl() string {
	return fileToString(serviceTplPath)
}
