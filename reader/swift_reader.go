package reader

import (
	"fmt"
	"strconv"
	"strings"
)

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
	Type  *SwiftType
	Name  string
	Value string
}

// SwiftStruct Swift Struct类型定义
type SwiftStruct struct {
	Name   string
	Fields map[string]*SwiftField
}

// SwiftEnum Swift Enum类型定义
type SwiftEnum struct {
	Name   string
	Fields map[string]*SwiftField
}

// SwiftService Swift Service类型定义
type SwiftService struct {
	Name      string
	Fields    map[string]*SwiftField
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
	ThriftReader   *ThriftReader
	SwiftThriftMap *SwiftThrift
}

// SwiftReader Swift Thrift Reader对象
var SwiftReader = &SwiftThriftReader{}

// InitSwiftThrift Swift Thrift Reader初始化
func (str *SwiftThriftReader) InitSwiftThrift(reader *ThriftReader) {

	str.ThriftReader = reader
	str.SwiftThriftMap = &SwiftThrift{}

	t := reader.Thrifts[reader.InputPath]

	enums := make(map[string]*SwiftEnum)
	for n, e := range t.Enums {
		enum := &SwiftEnum{}
		enum.Name = str.AssembleEnumName(n)
		enum.Fields = make(map[string]*SwiftField)

		for _, v := range e.Values {
			f := &SwiftField{}
			f.Name = v.Name
			f.Value = strconv.Itoa(v.Value)
			f.Type = &SwiftType{
				Type:      EnumType,
				Name:      enum.Name,
				InnerType: "",
			}
			enum.Fields[f.Name] = f
		}
		enums[enum.Name] = enum
	}

	str.SwiftThriftMap.Enums = enums
}

// AssembleEnumName 配置枚举的名称
func (str *SwiftThriftReader) AssembleEnumName(name string) string {

	components := strings.Split(name, ".")

	t := str.ThriftReader.Thrift

	if len(components) == 1 {
		ns := t.Namespaces["swift"]
		return AssembleName(ns, components[0])
	}

	return ""
}

// AssembleName 名称组装
func AssembleName(namespace string, name string) string {
	return fmt.Sprintf("%s%s", namespace, name[1:])
}
