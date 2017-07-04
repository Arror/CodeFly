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
    case {{ $ss.contextwrapper.FormatedFiledName $f.Name }} = {{ $f.Value }} {{ end }}
}
