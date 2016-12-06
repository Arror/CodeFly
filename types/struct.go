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
func (s *SwiftStruct) Name() string {
	return s.Namespace + s.Struct.Name
}
