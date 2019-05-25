// {{ $ss := . }}
// {{ .Name }}.swift
//
// 此文件由 CodeFly 生成，请不要手动修改
//

import Foundation
import Mappable

public enum {{ .Name }}: Int, EnumMappable {

    public typealias E = Int
    {{ range $i, $f := .Enum.Values }}
    case {{ $ss.Contextwrapper.FormatedFiledName $f.Name }} = {{ $f.Value }} {{ end }}

    public static var `default`: {{ .Name }} { {{ $result := $ss.Contextwrapper.EnumDefaultValue .Enum }}
        return {{ $result.Default }}
    }
}
