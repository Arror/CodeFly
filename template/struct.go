package template

// StructTemplate Swift Struct 模板
const StructTemplate = `//
// {{ .Name }}.swift
//
// 此文件由 CodeFly 生成，请不要手动修改
//

import Foundation
{{ $ss := . }}
public final class {{ .Name }}: JSONItem {
    {{ range $i, $f := .Fields }}
    pubilc var {{ $f.Name }}{{ $ss.DefaultValue $f }}{{ end }}
    
    public override var allKeys: Set<String> {
        return [{{ range $i, $f := .Fields }}{{ if ne $i 0 }}, {{ end }}"{{ $f.Name }}"{{ end }}]
    } 

    public override func fromJSON(json: AnyObject?) -> Bool {

        guard super.fromJSON(json) else { return false }
        guard let dict = json as? [String: AnyObject] else { return false }
        {{ range $i, $f := .Fields }}
        self.{{ $f.Name }} = {{ $ss.FromDict $f }}{{ end }}

        return true
    }

    public override func toJSON() -> AnyObject {

        var dict = [String: AnyObject]()
        {{ range $i, $f := .Fields }}
        dict["{{ $f.Name }}"] = {{ $ss.ToDict $f }}{{ end }}

        return dict
    }
}`
