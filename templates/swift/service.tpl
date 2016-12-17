// {{ $ss := . }}
// {{ $ss.Generator.ServiceName $ss.Service }}.swift
//
// 此文件由 CodeFly 生成，请不要手动修改
//

import Foundation

public final class {{ $ss.Generator.ServiceName $ss.Service }}: NSObject {
    {{ range $i, $m := $ss.Service.Methods }}
    public static func {{ $m.Name }}({{ range $i, $f := $m.Arguments }}{{ $f.Name }}: {{ $f.Type.String }} ,{{end}}completion: () -> Void, failure: (Error) -> Void) -> Bool {

        return true
    }
    {{ end }}
}
