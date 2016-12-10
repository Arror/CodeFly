package types

import p "github.com/arrors/go-thrift/parser"

// SwiftStruct 结构类型
type SwiftStruct struct {
	*p.Struct
	Namespace        string
	NamespaceMapping map[string]string
}
