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

/*
//
// {{ .Name }}.swift
//
// 此文件由 CodeFly 生成，请不要手动修改
//

import Foundation
{{ $ss := . }}
public final class {{ .Name }}: NSObject {
    {{ range $i, $m := .Methods }}
    public class func {{ .Name }}({{ range $i, $f := .Fields }}{{ $ss.GetParam $f }}, {{ end }}completion: ({{ $ss.ReturnType $m }}?) -> Void, failure: (NSError?) -> Void) -> Bool {

        guard let apiServerEngine = AppDelegate.current?.apiServiceEngine else { return false }

        let url: String = "{{ .URL }}"{{ $len := .Fields|len }}
        {{ if ne $len 0 }}
        var params = [String: AnyObject]()
        {{ range $i, $f := .Fields }}
        params["{{ $f.Name }}"] = {{ $ss.ToDict $f }}{{ end }}

        debugPrint("req: ", url, "\n", "params: ", params)
        {{ else }}
        debugPrint("req: ", url, "\n", "params: []")
        {{ end }}
        apiServerEngine.Post(url, parameters: {{ if ne $len 0 }}params{{ else }}nil{{ end }}, completion: { data in

            debugPrint("req: ", url, "\n", "rsp: ", data)

            let value = {{ $ss.ReturnType $m }}(json: data)

            completion(value)

            }, failure: { error in

                debugPrint("req: ", url, "\n", "rsp: ", error?.localizedDescription)

                failure(error)
        })

        return true
    }
    {{ end }}
}
*/
