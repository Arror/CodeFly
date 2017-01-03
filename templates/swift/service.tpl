// {{ $ss := . }}
// {{ $ss.Generator.ServiceName $ss.Service }}.swift
//
// 此文件由 CodeFly 生成，请不要手动修改
//

import Foundation

public struct {{ $ss.Generator.ServiceName $ss.Service }} {
    {{ range $i, $m := $ss.Service.Methods }}{{ $returnType := $ss.Generator.TypeString $m.ReturnType }}
    @discardableResult
    public static func {{ $ss.Generator.MethodName $m }}({{ range $i, $f := $m.Arguments }}{{ $f.Name }}: {{ $ss.Generator.TypeString $f.Type }} ,{{end}}completion: @escaping ({{ $returnType }}) -> Void, failure: @escaping (Error) -> Void) -> Bool {

        guard let caller = Invokers.caller else { return false }

        let path = "{{ $ss.Name }}/{{ $m.Name }}"{{ $argumentCount := $m.Arguments | len }}
        {{ if ne $argumentCount 0 }}
        var params = [String: Any]()
        {{ range $i, $f := $m.Arguments }}
        params["{{ $f.Name }}"] = {{ $f.Name }}.json{{ end }}{{ end }}
        {{  if ne $argumentCount 0  }}
        debugPrint("API: \(path)", "Request: ", params){{ else }}debugPrint("API: \(path)", "Request: [:]"){{ end }}
        
        caller.invoke(path: path, params: {{ if ne $argumentCount 0 }}params{{ else }}[:]{{ end }}, completion: { response in
            {{ if ne $returnType "Void" }}
            if let result = {{ $returnType }}(json: response) {

                debugPrint("API: \(path)", "Response: ", response)
                
                completion(result)

            } else {

                let error = InvokeError.invalidResponse
                
                debugPrint("API: \(path)", "Error: ", error)
                
                failure(error)
            }
            {{ else }}
            debugPrint("API: \(path)", "Response: [:]")

            completion()
            {{ end }}
        }, failure: { error in

            debugPrint("API: \(path)", "Error: ", error)
            
            failure(error)
        })

        return true
    }
    {{ end }}
}
