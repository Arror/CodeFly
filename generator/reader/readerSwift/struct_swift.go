package readerSwift

import "CodeFly/generator/reader"

const (
	// STInt16 Swift Type Int16
	STInt16 = "Int16"
	// STInt Swift Type Int
	STInt = "Int"
	// STInt64 Swift Type Int64
	STInt64 = "Int64"
	// STDouble Swift Type Double
	STDouble = "Double"
	// STBool Swift Type Bool
	STBool = "Bool"
	// STString Swift Type String
	STString = "String"
)

const (
	// ListType 数组类型
	ListType = "ListType"
	// EnumType 枚举类型
	EnumType = "EnumType"
	// PlainType 基本数据类型
	PlainType = "PlainType"
	// CustomerType 自定义数据类型
	CustomerType = "CustomerType"
	// Void 空类型
	Void = "Void"
)

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
	Name   string
	Fields []*SwiftField
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

// SwiftThriftReader Swift Thrift Reader
type SwiftThriftReader struct {
	ThriftReader   *reader.ThriftReader
	SwiftThriftMap *SwiftThrift
}
