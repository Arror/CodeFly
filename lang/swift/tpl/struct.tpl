//
// {{ .Namespace }}{{ .Struct.Name }}.swift
//
// 此文件由 CodeFly 生成，请不要手动修改
//

import Foundation
{{ $ss := . }}
public class {{ .Namespace }}{{ .Struct.Name }}: Base {
    {{ range $i, $f := .Fields }}
    public var {{ $f.Name }}: {{ $f.Type.PrintTypeString $ss.Namespace $ss.NamespaceMapping }}{{ $f.Type.DefaultValueString }}
    {{ end }}

    override public func from(json: Any) -> Bool {
        
        guard let dict = json as? [String: Any] else { return false }
        {{ range $i, $f := .Fields }}
        {{ $f.Name }} = {{ $f.ValueFromDict  $ss.Namespace $ss.NamespaceMapping }}{{ end }}

        return true
    }
    
    override public var json: Any {
        
        var dict = [String: Any]()
        {{ range $i, $f := .Fields }}
        dict["{{ $f.Name }}"] = {{ $f.Name }}{{ $f.Type.ToDictValue }}{{ end }}

        return dict
    }
}
