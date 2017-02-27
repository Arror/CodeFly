// {{ $ss := . }}
// {{ $ss.Name }}.swift
//
// 此文件由 CodeFly 生成，请不要手动修改
//

import Foundation

public struct {{ $ss.Name }}: Base {
    {{ range $i, $f := .Fields }}
    public var {{ $ss.SCA.FormatFiledName $f.Name }}: {{ $ss.SCA.TypeString $f.Type }}?{{ end }}

    public init?(json: Any?) {
        
        guard let dict = json as? [String: Any] else { return nil }
        {{ range $i, $f := .Fields }}
        {{ $ss.SCA.FormatFiledName $f.Name }} = dict <- "{{ $f.Name }}"{{ end }}
    }
    
    public var json: Any {
        
        var dict = [String: Any]()
        {{ range $i, $f := .Fields }}
        dict["{{ $f.Name }}"] = {{ $ss.SCA.FormatFiledName $f.Name }}?.json{{ end }}

        return dict
    }
}
