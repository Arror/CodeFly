//
// {{ .Namespace }}{{ .Service.Name }}.swift
//
// 此文件由 CodeFly 生成，请不要手动修改
//

import Foundation
{{ $ss := . }}
public final class {{ .Namespace }}{{ .Service.Name }}: NSObject {
    {{ range $i, $m := .Service.Methods }}{{ $void := $m.ReturnIsVoid }}
    public static func {{ $m.Name }}({{ range $i, $f := $m.Arguments }}{{ $f.Name }}: {{ $f.Type.PrintTypeString $ss.Namespace $ss.NamespaceMapping }}{{ $f.Type.Option }}, {{end}}completion: ({{ $m.ReturnTypeString $ss.Namespace $ss.NamespaceMapping }}) -> Void, failure: (Error) -> Void) -> Bool {
        
        let path = "{{ $ss.Name }}/{{ $m.Name }}"{{ $len := len $m.Arguments }}{{ if ne $len 0 }}

        var params = [String: Any](){{ range $i, $f := $m.Arguments }}
        params["{{ $f.Name }}"] = {{ $f.Name }}{{ $f.Type.ToDictValue }}{{ end }}{{ end }}

        session.request(path: path, method: {{ $m.AlamofireHTTPMothodEnum }}{{ if ne $len 0 }}, parameters: params{{end}}).responseJSON { response in
        
            {{ if ne $void true }}switch response.result {
            case .success(let value):
                if let result = {{ $m.ReturnType.ValueFromReturn $ss.Namespace $ss.NamespaceMapping }} {
                    completion(result)
                } else {
                    failure(AppError.invalidResponse)
                }
            case .failure(let error):
                failure(error)
            }{{ else }}switch response.result {
            case .success(_):
                completion(result)
            case .failure(let error):
                failure(error)
            }{{ end }}
        }

        return true
    }
    {{ end }}
}
