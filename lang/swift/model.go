package swift

import "github.com/arrors/go-thrift/parser"

// Enum Swift枚举类型
type Enum struct {
	*parser.Enum
	Namespace string
}

// Struct Swift结构类型
type Struct struct {
	*parser.Struct
	Namespace        string
	NamespaceMapping map[string]string
}

// Service Swift服务类型
type Service struct {
	*parser.Service
	Namespace        string
	NamespaceMapping map[string]string
}
