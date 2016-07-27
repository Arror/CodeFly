package template

// ServiceTemplate Swift Service 模板
const ServiceTemplate = `//
// {{ .Name }}.swift
//
// 此文件由 codefly 生成，请不要手动修改
//

import Foundation
{{ $ss := . }}
public final class {{ .Name }}: NSObject {
    {{ range $i, $m := .Methods }}
    public class func {{ .Name }}({{ range $i, $f := .Fields }}{{ .GetParam }}, {{ end }}completion: ({{ .ReturnType }}) -> Void, failure: (NSError) -> Void) -> Bool {
        return true
    }
    {{ end }}
}`
