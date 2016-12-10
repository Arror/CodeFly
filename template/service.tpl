//
// {{ .Namespace }}{{ .Service.Name }}.swift
//
// 此文件由 CodeFly 生成，请不要手动修改
//

import Foundation
{{ $ss := . }}
public final class {{ .Namespace }}{{ .Service.Name }}: NSObject {
    {{ range $i, $m := .Service.Methods }}{{ $void := $m.ReturnTypeIsVoid }}
    public static func {{ $m.Name }}({{ range $i, $f := $m.Arguments }}{{ $f.Name }}: {{ $f.Type.SwiftString  $ss.Thrifts $ss.Thrift $ss.Lang }}, {{end}}completion: ({{ if ne $void true }}{{ $m.ReturnType.SwiftString $ss.Thrifts $ss.Thrift $ss.Lang }}{{ end }}) -> Void, failure: (Error) -> Void) -> Bool {

        let path = {{ $ss.Path $m }}
        let url = host.appendingPathComponent(path){{ $len := $m.Arguments|len }}
        {{ if ne $len 0 }}
        var params = [String: Any]()
        {{ range $i, $f := $m.Arguments }}
        params["{{ $f.Name }}"] = {{ $f.ToJSON $ss.Thrifts $ss.Thrift $ss.Lang false }}{{ end }}{{ end }}

        Alamofire.request(url, method: {{ $m.HttpMothod }}{{ if ne $len 0 }}, parameters: params{{end}}).responseJSON { response in
        
            switch response.result {

            case .success(let value):

                {{ if ne $void true }}if let result = {{ $m.ReturnType.ReturnValueFromJSON $ss.Thrifts $ss.Thrift $ss.Lang }} {
                    completion(result)
                } else {
                    failure(AppError.invalidResponse)
                }{{ else }}completion(){{ end }}

            case .failure(let error):

                failure(error)
            }
        }

        return true
    }
    {{ end }}
}
