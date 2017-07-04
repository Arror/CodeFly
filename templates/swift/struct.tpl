// {{ $ss := . }}
// {{ $ss.Name }}.swift
//
// 此文件由 CodeFly 生成，请不要手动修改
//

import Foundation
import Mappable

public struct {{ $ss.Name }}: Mappable {
    {{ range $i, $f := .Fields }} {{ $name := $ss.contextwrapper.FormatedFiledName $f.Name }} {{ $type := $ss.contextwrapper.TypeString $f.Type }}
    public var {{ $name }}: {{ $type }}{{ if $f.Optional }}?{{ else }} = {{ $ss.contextwrapper.DefaultValue $type }}{{ end }}{{ end }}

    public init() {}

    public init?(any: Any?) {
        
        guard let wrapper = MapWrapper(any) else { return nil }
        {{ range $i, $f := .Fields }} {{ $name := $ss.contextwrapper.FormatedFiledName $f.Name }}
        {{ $name }} = wrapper[CodingKeys.{{ $name }}]{{ end }}
    }
    
    public var json: Any {
        return MapWrapper.exportAny { wrapper in {{ range $i, $f := .Fields }} {{ $name := $ss.contextwrapper.FormatedFiledName $f.Name }}
            wrapper[CodingKeys.{{ $name }}] = {{ $name }}{{ end }}
        }
    }

    private enum CodingKeys: String, CodingKey { {{ range $i, $f := .Fields }}
        case {{ $ss.contextwrapper.FormatedFiledName $f.Name }} = "{{ $f.Name }}"{{ end }}
    }
}
