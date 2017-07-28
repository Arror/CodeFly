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
        var param: [String: Any] = [:]
        {{- range $i, $f := $m.Arguments }}
        {{- $result := $ss.contextwrapper.ParserType $f.Type }}
        {{- $isBase := $ss.contextwrapper.IsBaseType $result.Type }}
        {{- $name := $ss.contextwrapper.FormatedFiledName $f.Name }}
        param["{{ $f.Name }}"] = {{ $name }}{{- if ne $isBase true }}{{- if $f.Optional }}?{{- end }}.any{{- end }}
        {{- end }}
        {{- else }}
        let param: [String: Any] = [:]
        {{- end }}
        {{ "" }}
        <#session#>.<#invoke#>(path: path, param: param, completion: { responeObject in

            do {

                let result = try {{ $response.Type }}(any: responeObject)

                completion(Result.succeed(result))

            } catch let error {

                completion(Result.failure(error))
            }

        }, failure: { error in

            completion(Result.failure(error))
        })
    }
    {{ "" }}
    {{- end }}
}