// {{ $ss := . }}
// {{ $ss.Generator.ServiceName $ss.Service }}.swift
//
// 此文件由 CodeFly 生成，请不要手动修改
//

import Foundation

public struct {{ $ss.Generator.ServiceName $ss.Service }} {
    {{ range $i, $m := $ss.Service.Methods }}
    public static func {{ $ss.Generator.MethodName $m }}({{ range $i, $f := $m.Arguments }}{{ $f.Name }}: {{ $ss.Generator.TypeString $f.Type }} ,{{end}}completion: @escaping ({{ $ss.Generator.TypeString $m.ReturnType }}) -> Void, failure: @escaping (Error) -> Void) -> Bool {

        return true
    }
    {{ end }}
}
