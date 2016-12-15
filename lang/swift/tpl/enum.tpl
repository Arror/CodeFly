//
// {{ .Generator.EnumName .Enum }}.swift
//
// 此文件由 CodeFly 生成，请不要手动修改
//

import Foundation

public enum {{ .Generator.EnumName .Enum }}: Int, Enum {
    {{ range $i, $f := .Enum.Values }}
    case {{ $f.Name }} = {{ $f.Value }} {{ end }}
}
