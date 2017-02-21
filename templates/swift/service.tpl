// {{ $ss := . }}
// {{ $ss.Name }}.swift
//
// 此文件由 CodeFly 生成，请不要手动修改
//

import Foundation

public struct {{ $ss.Name }} {
    {{ range $i, $m := $ss.Service.Methods }}{{ $returnType := $ss.STA.TypeString $m.ReturnType }}
    @discardableResult
    public static func {{ $ss.MethodName $m }}({{ range $i, $f := $m.Arguments }}{{ $f.Name }}: {{ $ss.STA.TypeString $f.Type }} ,{{end}}completion: @escaping ({{ $returnType }}) -> Void, failure: @escaping (Error) -> Void) -> Bool {

        guard let engine = WebAPIEngine.engine else { return false }

        let path = "{{ $ss.Service.Name }}/{{ $m.Name }}"{{ $argumentCount := $m.Arguments | len }}
        {{ if ne $argumentCount 0 }}
        var params = [String: Any](){{ range $i, $f := $m.Arguments }}
        params["{{ $f.Name }}"] = {{ $f.Name }}.json{{ end }}{{ end }}
        {{  if ne $argumentCount 0  }}
        debugPrint(path, "Request: ", params){{ else }}debugPrint(path, "Request:", [:]){{ end }}
        
        engine.post(path: path, params: {{ if ne $argumentCount 0 }}params{{ else }}[:]{{ end }}, completion: { response in

            debugPrint(path, "Response:", response)
            {{ if ne $returnType "Void" }}
            if let result = {{ $returnType }}(json: response) {
                completion(result)
            } else {
                let error = WebAPIEngineError.invalidResponse
                debugPrint(path, "Error:", error)
                failure(error)
            }
            {{ else }}
            debugPrint(path, "Response", [:])
            completion()
            {{ end }}
        }, failure: { error in
            debugPrint(path, "Error:", error)
            failure(error)
        })

        return true
    }
    {{ end }}
}
