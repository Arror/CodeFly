package template

// StructTemplate Swift Struct 模板
const StructTemplate = `//
// {{ .Name }}.swift
//
// 此文件由 CodeFly 生成，请不要手动修改
//

import Foundation
{{ $ss := . }}
public class {{ .Name }}: NSObject, JSONItem {
    {{ range $i, $f := .Fields }}
    var {{ $f.Name }}{{ $ss.DefaultValue $f }}{{ end }}
    
    public var allKeys: Set<String> {
        return [{{ range $i, $f := .Fields }}{{ if ne $i 0 }}, {{ end }}"{{ $f.Name }}"{{ end }}]
    } 

    public required init?(json: AnyObject?) {

        super.init()

        guard let json = json as? [String: AnyObject] else { return nil }
        {{ range $i, $f := .Fields }}
        self.{{ $f.Name }} = {{ $ss.FromDict $f }}{{ end }}
    }

    public func toJSON() -> AnyObject? {

        var json = [String: AnyObject]()
        {{ range $i, $f := .Fields }}
        json["{{ $f.Name }}"] = {{ $ss.ToDict $f }}{{ end }}

        return json
    }
}`
