package types

import p "github.com/arrors/go-thrift/parser"

// SwiftStruct 结构类型
type SwiftStruct struct {
	*p.Struct
	Thrift    *p.Thrift
	Thrifts   map[string]*p.Thrift
	Lang      string
	Namespace string
}

// Name 结构名称
func (e *SwiftStruct) Name() string {
	return e.Namespace + e.Struct.Name[1:]
}

/*
public protocol JSONParser: NSObjectProtocol {

    init()
    func from(json: Any) -> Bool
    var json: Any { get }
}

extension JSONParser {

    public init?(json: Any?) {
        guard let json = json else { return nil }
        self.init()
        guard !self.from(json: json) else { return nil }
    }
}

extension Array where Element: JSONParser {

    public init?(json: Any?) {
        guard let elements = json as? [Any] else { return nil }
        self = elements.flatMap { Element(json: $0) }
    }

    public mutating func from(json: Any) -> Bool {
        if let elements = Array<Element>(json: json) {
            self = elements
            return true
        } else {
            return false
        }
    }

    public var json: Any { return self.flatMap { $0.json } }
}

public class JSON: NSObject, JSONParser {

    override public required init() { super.init() }

    @discardableResult
    public func from(json: Any) -> Bool { return false }

    public var json: Any { fatalError("Not implemented.") }
}

extension JSON: NSCopying {

    public func copy(with zone: NSZone? = nil) -> Any {
        let obj = type(of: self).init()
        obj.from(json: self.json)
        return obj
    }
}
*/
