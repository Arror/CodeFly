// {{ $ss := . }}
// {{ $ss.Generator.StructName $ss.Struct }}.swift
//
// 此文件由 CodeFly 生成，请不要手动修改
//

import Foundation

public class {{ $ss.Generator.StructName $ss.Struct }}: Base {
    {{ range $i, $f := .Fields }}
    public var {{ $f.Name }}: {{ $ss.Generator.PropertyType $f }} {{ $ss.Generator.DefaultValue $f }}
    {{ end }}

    override public func from(json: Any) -> Bool {
        
        guard let dict = json as? [String: Any] else { return false }
        {{ range $i, $f := .Fields }}
        {{ $f.Name }} = {{ $ss.Generator.ValueFromJSON $f }}{{ end }}

        return true
    }
    
    override public var json: Any {
        
        var dict = [String: Any]()
        {{ range $i, $f := .Fields }}
        dict["{{ $f.Name }}"] = {{ $ss.Generator.ValueToJSON $f }}{{ end }}

        return dict
    }
}
