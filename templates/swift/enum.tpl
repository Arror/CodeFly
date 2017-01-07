//
// {{ .Name }}.swift
//
// 此文件由 CodeFly 生成，请不要手动修改
//

import Foundation

public enum {{ .Name }}: Int, Enum {

    public typealias E = Int
    {{ range $i, $f := .Enum.Values }}
    case {{ $f.Name }} = {{ $f.Value }} {{ end }}
}
