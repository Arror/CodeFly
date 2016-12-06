package types

import p "github.com/arrors/go-thrift/parser"

// SwiftEnum 枚举类型
type SwiftEnum struct {
	*p.Enum
	Namespace string
}

// Name 枚举名称
func (e *SwiftEnum) Name() string {
	return e.Namespace + e.Enum.Name
}
