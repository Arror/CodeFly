// {{ $ss := . }}
// {{ $ss.Name }}.swift
//
// 此文件由 CodeFly 生成，请不要手动修改
//

import Foundation
import Mappable

public struct {{ $ss.Name }} {
    {{ "" }}
    {{- range $i, $m := $ss.Methods }}
        {{- $count := len $m.Arguments }}
        {{- $response := $ss.contextwrapper.ParserType $m.ReturnType }}
    static func {{ $m.Name }}(
            {{- if ne 0 $count }}
                {{- range $i, $f := $m.Arguments }}
                    {{- $request := $ss.contextwrapper.ParserType $f.Type }}{{ $f.Name }}: {{ $request.Type }},{{ " " }} 
                {{- end }}
            {{- end }}completion: @escaping (Result<{{ $response.Type }}>) -> Void) {
        {{ "" }}
        let path = "{{ $ss.contextwrapper.GetPath $m }}"
        {{ "" }}
        {{- if ne $count 0 }}
        {{- range $i, $f := $m.Arguments }}
        let param = {{ $f.Name }}.any as? [String: Any] ?? [:]
        {{- end }}
        {{- else }}
        let param = [:]
        {{- end }}
        {{ "" }}
        #<SessionInstance>#.#<invoke>#(path: path, param: param, completion: { responeObject in

            if let r = try? $response.Type(any: responeObject) {

                completion(Result.succeed(r))

            } else {

                completion(Result.failure(#<Error>#))
            }

        }, failure: { error in

            completion(Result.failure(error))
        })
    }
    {{ "" }}
    {{- end }}
}