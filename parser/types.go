package parser

import p "github.com/samuel/go-thrift/parser"

// SwiftType Swift类型定义
type SwiftType struct {
	Type      string
	Name      string
	InnerType string
}

// SwiftField Swift Field类型定义
type SwiftField struct {
	Type  *SwiftType
	Name  string
	Value string
}

// SwiftStruct Swift Struct类型定义
type SwiftStruct struct {
	Name   string
	Fields []*SwiftField
}

// SwiftEnum Swift Enum类型定义
type SwiftEnum struct {
	Name  string
	Cases []*SwiftField
}

// SwiftService Swift Service类型定义
type SwiftService struct {
	Name    string
	Methods []*SwiftMethod
}

// SwiftMethod Swift Method类型定义
type SwiftMethod struct {
	Name      string
	Fields    []*SwiftField
	ValueType *SwiftType
}

// SwiftThrift Swift Thrift类型定义
type SwiftThrift struct {
	Structs  map[string]*SwiftStruct
	Enums    map[string]*SwiftEnum
	Services map[string]*SwiftService
}

// SwiftThriftComponents 完整的Thrift信息
type SwiftThriftComponents struct {
	Thrifts     map[string]*p.Thrift
	Thrift      *p.Thrift
	InputPath   string
	OutputPath  string
	SwiftThrift *SwiftThrift
}
