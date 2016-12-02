package types

import tp "github.com/arrors/go-thrift/parser"

// SwiftEnum 枚举类型
type SwiftEnum struct {
	*tp.Enum
	Namespace string
}

// Name 枚举名称
func (e *SwiftEnum) Name() string {
	return e.Namespace + e.Enum.Name[1:]
}

/*
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
*/
