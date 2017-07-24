// {{ $ss := . }}
// {{ $ss.Name }}.swift
//
// 此文件由 CodeFly 生成，请不要手动修改
//

import Foundation
import Mappable

public struct {{ $ss.Name }}: Mappable {
    {{ "" }}
    {{- range $i, $f := .Fields }}
        {{- $name := $ss.contextwrapper.FormatedFiledName $f.Name }}
        {{- $result := $ss.contextwrapper.ParserType $f.Type }}
    public var {{ $name }}: {{ $result.Type }}{{ if $f.Optional }}?{{ end }}
    {{- end }}

    public init() {
        {{ "" }}
        {{- range $i, $f := .Fields }}
        {{- $name := $ss.contextwrapper.FormatedFiledName $f.Name }}
            {{- $result := $ss.contextwrapper.ParserType $f.Type }}
            {{- if eq $f.Optional false }}
        {{ $name }} = {{ $result.Default }}
            {{- end }}
        {{- end }}
    }

    public init(any: Any?) { {{ $count := len .Fields }}
        {{ "" }}
        {{- if ne $count 0 }}
        guard let map = any as? [String: Any] else { throw MappableError.dictionaryConvertFailed(any) }
        {{ "" }}
            {{- range $i, $f := .Fields }}
                {{- $name := $ss.contextwrapper.FormatedFiledName $f.Name }}
                {{- $result := $ss.contextwrapper.ParserType $f.Type }}
                {{- $isBase := $ss.contextwrapper.IsBaseType $result.Type }}
                {{- if $f.Optional }}
        {{ $name }} = {{ if $isBase }}map["{{ $name }}"] as? {{ $result.Type }}{{ else }}try? {{ $result.Type }}(any: map["{{ $name }}"]){{ end }}
                {{- else }}
                    {{- if $isBase }}
        if let p = map["{{ $name }}"] as? {{ $result.Type }} {
            {{ $name }} = p
        } else {
            throw MappableError.propertyConvertFailed(map["{{ $name }}"])
        }
                    {{- else }}
        if let p = {{ $result.Type }}(any: map["{{ $name }}"]) {
            {{ $name }} = p
        } else {
            throw MappableError.propertyConvertFailed(map["{{ $name }}"])
        }
                    {{- end }}
                {{- end }}
            {{- end }}
        {{- end }}
    }
    
    public var any: Any { {{ $count := len .Fields }}
        {{- if eq $count 0 }}
        return [:]
        {{- else }}
        {{ "" }}
        var dict = [String: Any]()
        {{ "" }}
            {{- range $i, $f := .Fields }}
                {{- $name := $ss.contextwrapper.FormatedFiledName $f.Name }}
                {{- $result := $ss.contextwrapper.ParserType $f.Type }}
                {{- $isBase := $ss.contextwrapper.IsBaseType $result.Type }}
        dict["{{ $f.Name }}"] = {{ if $isBase }}{{ $name }}{{ else }}{{ $name }}{{ if $f.Optional }}?{{ end }}.any{{ end }}
            {{- end }}
        {{ "" }}
        return dict
        {{- end }}
    }
}
