//
// {{ .Name }}.swift
//
// 此文件由 CodeFly 生成，请不要手动修改
//

import Foundation
{{ $ss := . }}
public class {{ .Name }}: JSON {
    {{ range $i, $f := .Fields }}
    public var {{ $f.Name }}: {{ $f.Type.SwiftString $ss.Thrifts $ss.Thrift $ss.Lang }}
    {{ end }}
}
