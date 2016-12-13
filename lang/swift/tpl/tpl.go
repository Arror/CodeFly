package tpl

import (
	"io/ioutil"
	"log"
)

const (
	enumTplPath    = "./lang/swift/tpl/enum.tpl"
	structTplPath  = "./lang/swift/tpl/struct.tpl"
	serviceTplPath = "./lang/swift/tpl/service.tpl"
)

func fileToString(fp string) string {

	buffer, err := ioutil.ReadFile(fp)

	if err != nil {
		log.Fatalln(err.Error())
	}

	return string(buffer)
}

// SwiftEnumTplName Swift Enum 模板名称
const SwiftEnumTplName = "EnumTemplate"

// SwiftStructTplName Swift Struct 模板名称
const SwiftStructTplName = "StructTemplate"

// SwiftServiceTpleName Swift Service 模板名称
const SwiftServiceTpleName = "ServiceTemplate"

// SwiftEnumTpl Swift Enum 模板
var SwiftEnumTpl = fileToString(enumTplPath)

// SwiftStructTpl Swift Struct 模板
var SwiftStructTpl = fileToString(structTplPath)

// SwiftServiceTpl Swift Service 模板
var SwiftServiceTpl = fileToString(serviceTplPath)
