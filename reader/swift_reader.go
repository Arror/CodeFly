package reader

import (
	"fmt"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/samuel/go-thrift/parser"
)

// Swift Swift语言
const Swift = "swift"

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

	structs := make(map[string]*SwiftStruct)
	for n, s := range t.Structs {
		_struct := &SwiftStruct{}
		_struct.Name = AssembleName(t.Namespaces[Swift], n)
		_struct.Fields = make(map[string]*SwiftField)

		for _, f := range s.Fields {
			_f := &SwiftField{}
			_f.Name = f.Name
			_f.Type = str.GetSwiftType(f)
			_struct.Fields[_f.Name] = _f
		}
		structs[_struct.Name] = _struct
	}
	str.SwiftThriftMap.Structs = structs

	enums := make(map[string]*SwiftEnum)
	for n, e := range t.Enums {
		_enum := &SwiftEnum{}
		_enum.Name = AssembleName(t.Namespaces[Swift], n)
		_enum.Fields = make(map[string]*SwiftField)

		for _, v := range e.Values {
			_f := &SwiftField{}
			_f.Name = v.Name
			_f.Value = strconv.Itoa(v.Value)
			_f.Type = &SwiftType{
				Type:      EnumType,
				Name:      _enum.Name,
				InnerType: "",
			}
			_enum.Fields[_f.Name] = _f
		}
		enums[_enum.Name] = _enum
	}
	str.SwiftThriftMap.Enums = enums
}

// GetSwiftType 通过 Field 获取 Type
func (str *SwiftThriftReader) GetSwiftType(f *parser.Field) *SwiftType {

	st := &SwiftType{}

	if b, tn := str.IsPlainType(f.Type); b {
		st.Name = tn
		st.Type = PlainType
		st.InnerType = ""
		fmt.Println(st)
		return st
	}

	if b, tn := str.IsEnumType(f.Type); b {
		st.Name = tn
		st.Type = EnumType
		st.InnerType = ""
		fmt.Println(st)
		return st
	}

	if b, tn := str.IsCustomerType(f.Type); b {
		st.Name = tn
		st.Type = EnumType
		st.InnerType = ""
		fmt.Println(st)
		return st
	}

	if b, tn, innerType := str.IsListType(f.Type); b {
		st.Name = tn
		st.Type = ListType
		st.InnerType = innerType
		fmt.Println(st)
		return st
	}

	return st
}

// IsPlainType 如果是基本数据类型，返回该类型名
func (str *SwiftThriftReader) IsPlainType(t *parser.Type) (bool, string) {
	n := t.Name
	switch n {
	case TTI16:
		return true, TypeMapping[TTI16]
	case TTI32:
		return true, TypeMapping[TTI32]
	case TTI64:
		return true, TypeMapping[TTI64]
	case TTDouble:
		return true, TypeMapping[TTDouble]
	case TTBool:
		return true, TypeMapping[TTBool]
	case TTString:
		return true, TypeMapping[TTString]
	default:
		return false, ""
	}
}

// IsEnumType 如果是枚举数据类型，返回该类型名
func (str *SwiftThriftReader) IsEnumType(t *parser.Type) (bool, string) {

	thrift := str.ThriftReader.Thrift
	thrifts := str.ThriftReader.Thrifts

	n := t.Name
	components := strings.Split(n, ".")

	if len(components) == 1 {
		for _, s := range thrift.Enums {
			if s.Name == components[0] {
				return true, AssembleName(thrift.Namespaces[Swift], s.Name)
			}
		}
	}

	if len(components) == 2 {
		for k, t := range thrifts {
			f := strings.Split(filepath.Base(k), ".")[0]
			if components[0] == f {
				for _, s := range t.Enums {
					if s.Name == components[1] {
						return true, AssembleName(t.Namespaces[Swift], components[1])
					}
				}
			}
		}
	}
	return false, ""
}

// IsCustomerType 如果是自定义数据类型，返回该类型名
func (str *SwiftThriftReader) IsCustomerType(t *parser.Type) (bool, string) {

	thrift := str.ThriftReader.Thrift
	thrifts := str.ThriftReader.Thrifts

	n := t.Name
	components := strings.Split(n, ".")

	if len(components) == 1 {
		for _, s := range thrift.Structs {
			if s.Name == components[0] {
				return true, AssembleName(thrift.Namespaces[Swift], s.Name)
			}
		}
	}

	if len(components) == 2 {
		for k, t := range thrifts {
			f := strings.Split(filepath.Base(k), ".")[0]
			if components[0] == f {
				for _, s := range t.Structs {
					if s.Name == components[1] {
						return true, AssembleName(t.Namespaces[Swift], components[1])
					}
				}
			}
		}
	}
	return false, ""
}

// IsListType 如果是数组数据类型，返回该数据类型及内部类型名
func (str *SwiftThriftReader) IsListType(t *parser.Type) (bool, string, string) {

	if t.Name == "list" {
		innerType := t.ValueType
		if b, tn := str.IsPlainType(innerType); b {
			return b, t.Name, tn
		}
		if b, tn := str.IsEnumType(innerType); b {
			return b, t.Name, tn
		}
		if b, tn := str.IsCustomerType(innerType); b {
			return b, t.Name, tn
		}
	}
	return false, "", ""
}

// AssembleName 名称组装
func AssembleName(namespace string, name string) string {
	return fmt.Sprintf("%s%s", namespace, name[1:])
}
