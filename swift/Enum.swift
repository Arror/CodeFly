//
// Enum.swift
//
// 模板生成的 enum 必须遵守 Enum 协议, RawValue 为 Int 类型
//

protocol Enum {

    associatedtype Value

    var rawValue: Value { get }

    init?(rawValue: Value)
}

extension Enum {

    public init?(json: Any?) {

        guard let value = json as? Value else { return nil }

        self.init(rawValue: value)
    }

    public var json: Any {

        return self.rawValue
    }
}