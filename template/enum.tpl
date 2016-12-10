//
// {{ .Namespace }}{{ .Enum.Name }}.swift
//
// 此文件由 CodeFly 生成，请不要手动修改
//

import Foundation

public enum {{ .Namespace }}{{ .Enum.Name }}: Int, Enum {
    {{ range $i, $f := .Enum.Values }}
    case {{ $f.Name }} = {{ $f.Value }} {{ end }}
}
