package template

// EnumTpl Swift Enum 模板
const EnumTpl = `//
// {{ .Name }}.swift
//
// 此文件由 CodeFly 生成，请不要手动修改
//

import Foundation

public enum {{ .Name }}: Int, Enum {
    {{ range $i, $f := .Enum.Values }}
    case {{ $f.Name }} = {{ $f.Value }} {{ end }}
}`
