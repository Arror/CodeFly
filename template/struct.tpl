//
// {{ .Name }}.swift
//
// 此文件由 CodeFly 生成，请不要手动修改
//

import Foundation
{{ $ss := . }}
public class {{ .Name }}: JSON {
    {{ range $i, $f := .Fields }}
    {{ $f.PropertDefine $ss.Thrifts $ss.Thrift $ss.Lang }}
    {{ end }}

    override func from(json: Any) -> Bool {
        
        guard let dict = json as? [String: Any] else { return false }
        {{ range $i, $f := .Fields }}
        {{ $f.FromJSON $ss.Thrifts $ss.Thrift $ss.Lang }}{{ end }}

        return true
    }
    
    override var json: Any {
        
        var dict = [String: Any]()
        {{ range $i, $f := .Fields }}
        {{ $f.ToJSON $ss.Thrifts $ss.Thrift $ss.Lang }}{{ end }}

        return dict
    }
}
