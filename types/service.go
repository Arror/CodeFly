package types

import p "github.com/arrors/go-thrift/parser"

// SwiftService 服务类型
type SwiftService struct {
	*p.Service
	Namespace        string
	NamespaceMapping map[string]string
}
