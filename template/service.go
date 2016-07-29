package template

// ServiceTemplate Swift Service 模板
const ServiceTemplate = `//
// {{ .Name }}.swift
//
// 此文件由 codefly 生成，请不要手动修改
//

import Foundation
{{ $ss := . }}
public final class {{ .Name }} {
    {{ range $i, $m := .Methods }}
    public class func {{ .Name }}({{ range $i, $f := .Fields }}{{ $ss.GetParam $f }}, {{ end }}completion: ({{ $ss.ReturnType $m }}?) -> Void, failure: (NSError?) -> Void) -> Bool {

        guard let apiServerEngine = AppDelegate.current?.apiServiceEngine else { return false }

        let url: String = "{{ .URL }}"{{ $len := .Fields|len }}
        {{ if ne $len 0 }}
        var params = [String: AnyObject]()
        {{ range $i, $f := .Fields }}
        params["{{ $f.Name }}"] = {{ $ss.ToDict $f }}{{ end }}
         
        debugPrint("req: ", url, "\n", "params: ", params)
        {{ else }}
        debugPrint("req: ", url, "\n", "params: []")
        {{ end }}
        apiServerEngine.Post(url, parameters: {{ if ne $len 0 }}params{{ else }}nil{{ end }}, completion: { data in
            
            debugPrint("req: ", url, "\n", "rsp: ", data)
            
            let value = {{ $ss.ReturnType $m }}(json: data)
            
            completion(value)
            
            }, failure: { error in
                
                debugPrint("req: ", url, "\n", "rsp: ", error?.localizedDescription)
                
                failure(error)
        })

        return true
    }
    {{ end }}
}`
