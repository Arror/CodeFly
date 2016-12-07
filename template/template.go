package template

import (
	"io/ioutil"
	"log"
)

const (
	// SwiftEnumTplName Swift 枚举模板名称
	SwiftEnumTplName = "EnumTemplate"
	// SwiftStructTemplateName Swift Strcut Template名字
	SwiftStructTemplateName = "StructTemplate"
	// SwiftServiceTemplateName Swift Service Template名字
	SwiftServiceTemplateName = "ServiceTemplate"
)

const (
	enumPath    = "./template/enum.tpl"
	structPath  = "./template/struct.tpl"
	servicePath = "./template/service.tpl"
)

func file2string(fp string) string {

	buffer, err := ioutil.ReadFile(fp)

	if err != nil {
		log.Fatalln(err.Error())
	}

	return string(buffer)
}

// SwiftEnumTpl Swif枚举模板
func SwiftEnumTpl() string {
	return file2string(enumPath)
}

// SwiftStructTpl Swif结构模板
func SwiftStructTpl() string {
	return file2string(structPath)
}

// SwiftServiceTpl Swif服务模板
func SwiftServiceTpl() string {
	return file2string(servicePath)
}
