// {{ $ss := . }}
// {{ $ss.Name }}.swift
//
// 此文件由 CodeFly 生成，请不要手动修改
//

import Foundation
import Mappable

public struct {{ $ss.Name }}: Mappable {
    {{ range $i, $f := .Fields }} {{ $n := $ss.Contextwrapper.FormatedFiledName $f.Name }} {{ $result := $ss.Contextwrapper.ParserType $f.Type }}
    public var {{ $n }}: {{ $result.Type }}{{ if $f.Optional }}?{{ else }} = {{ $result.Default }}{{ end }}{{ end }}

    public init() {}

    public init?(any: Any?) {
        
        guard let wrapper = MapWrapper(any) else { return nil }
        {{ range $i, $f := .Fields }} {{ $name := $ss.Contextwrapper.FormatedFiledName $f.Name }}
        {{ $name }} = wrapper[CodingKeys.{{ $name }}]{{ end }}
    }
    
    public var json: Any {
        return MapWrapper.exportAny { wrapper in {{ range $i, $f := .Fields }} {{ $name := $ss.Contextwrapper.FormatedFiledName $f.Name }}
            wrapper[CodingKeys.{{ $name }}] = {{ $name }}{{ end }}
        }
    }

    private enum CodingKeys: String, CodingKey { {{ range $i, $f := .Fields }}
        case {{ $ss.Contextwrapper.FormatedFiledName $f.Name }} = "{{ $f.Name }}"{{ end }}
    }
}
