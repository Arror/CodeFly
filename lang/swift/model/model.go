package model

import "github.com/arrors/go-thrift/parser"

// Namespaces 名称空间信息
type Namespaces struct {
	Namespace        string
	NamespaceMapping map[string]string
}

// Enum Swift枚举类型
type Enum struct {
	*parser.Enum
	Namespaces
}

// Struct Swift结构类型
type Struct struct {
	*parser.Struct
	Namespaces
}

// Service Swift服务类型
type Service struct {
	*parser.Service
	Namespaces
}
