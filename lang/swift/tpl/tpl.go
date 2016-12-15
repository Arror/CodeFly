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

// SwiftEnumTplName Swift enum template name
const SwiftEnumTplName = "EnumTemplate"

// SwiftStructTplName Swift struct template name
const SwiftStructTplName = "StructTemplate"

// SwiftServiceTpleName Swift service template name
const SwiftServiceTpleName = "ServiceTemplate"

// SwiftEnumTpl Swift enum template
var SwiftEnumTpl = fileToString(enumTplPath)

// SwiftStructTpl Swift struct template
var SwiftStructTpl = fileToString(structTplPath)

// SwiftServiceTpl Swift service template
var SwiftServiceTpl = fileToString(serviceTplPath)
