package reader

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
)

// TypeMapping 类型映射
var TypeMapping = map[string]string{
	TTI16:    STInt16,
	TTI32:    STInt,
	TTI64:    STInt64,
	TTDouble: STDouble,
	TTBool:   STBool,
	TTString: STString,
}

// SwiftType Swift类型定义
type SwiftType struct {
	Type      string
	Name      string
	InnerType string
}

// SwiftField Swift Field类型定义
type SwiftField struct {
	Type  SwiftType
	Name  string
	Value string
}

// SwiftStruct Swift Struct类型定义
type SwiftStruct struct {
}

// SwiftEnum Swift Enum类型定义
type SwiftEnum struct {
}

// SwiftService Swift Service类型定义
type SwiftService struct {
}

// SwiftThrift Swift Thrift类型定义
type SwiftThrift struct {
	Structs []SwiftStruct
	Enums   []SwiftEnum
	Service []SwiftService
}
