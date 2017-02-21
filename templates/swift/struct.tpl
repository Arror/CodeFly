// {{ $ss := . }}
// {{ $ss.Name }}.swift
//
// 此文件由 CodeFly 生成，请不要手动修改
//

import Foundation

public struct {{ $ss.Name }}: Base {
    {{ range $i, $f := .Fields }}
    public var {{ $f.Name }}: {{ $ss.SCA.TypeString $f.Type }}?{{ end }}

    public init?(json: Any?) {
        
        guard let dict = json as? [String: Any] else { return nil }
        {{ range $i, $f := .Fields }}
        {{ $f.Name }} = dict <- "{{ $f.Name }}"{{ end }}
    }
    
    public var json: Any {
        
        var dict = [String: Any]()
        {{ range $i, $f := .Fields }}
        dict["{{ $f.Name }}"] = {{ $f.Name }}?.json{{ end }}

        return dict
    }
}
