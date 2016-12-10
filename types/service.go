package types

import p "github.com/arrors/go-thrift/parser"
import "strings"

// SwiftService 服务类型
type SwiftService struct {
	*p.Service
	Thrift           *p.Thrift
	Thrifts          map[string]*p.Thrift
	Lang             string
	Namespace        string
	NamespaceMapping map[string]string
}

// Path 服务的路径
func (s *SwiftService) Path(m *p.Method) string {
	return strings.ToLower(s.Service.Name + "/" + m.Name)
}
