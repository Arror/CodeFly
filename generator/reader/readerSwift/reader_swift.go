package readerSwift

import (
	"path/filepath"
	"strconv"
	"strings"

	"CodeFly/generator/reader"
	"CodeFly/global"

	"github.com/samuel/go-thrift/parser"
)

// TypeMapping 类型映射
var TypeMapping = map[string]string{
	global.TTI16:    STInt16,
	global.TTI32:    STInt,
	global.TTI64:    STInt64,
	global.TTDouble: STDouble,
	global.TTBool:   STBool,
	global.TTString: STString,
}

// InitSwiftThrift Swift Thrift Reader初始化
func (str *SwiftThriftReader) InitSwiftThrift(reader *reader.ThriftReader) {

	str.ThriftReader = reader
	str.SwiftThriftMap = &SwiftThrift{}

	t := reader.Thrifts[reader.InputPath]

	structs := make(map[string]*SwiftStruct)
	for n, s := range t.Structs {
		_struct := &SwiftStruct{}
		_struct.Name = AssembleName(t.Namespaces[global.Swift], n)
		_struct.Fields = make([]*SwiftField, 0, 10)

		for _, f := range s.Fields {
			_f := &SwiftField{}
			_f.Name = f.Name
			_f.Type = str.GetSwiftType(f.Type)
			_struct.Fields = append(_struct.Fields, _f)
		}
		structs[_struct.Name] = _struct
	}
	str.SwiftThriftMap.Structs = structs

	enums := make(map[string]*SwiftEnum)
	for n, e := range t.Enums {
		_enum := &SwiftEnum{}
		_enum.Name = AssembleName(t.Namespaces[global.Swift], n)
		_enum.Fields = make([]*SwiftField, 0, 10)

		for _, v := range e.Values {
			_f := &SwiftField{}
			_f.Name = v.Name
			_f.Value = strconv.Itoa(v.Value)
			_f.Type = &SwiftType{
				Type:      EnumType,
				Name:      _enum.Name,
				InnerType: "",
			}
			_enum.Fields = append(_enum.Fields, _f)
		}
		enums[_enum.Name] = _enum
	}
	str.SwiftThriftMap.Enums = enums

	services := make(map[string]*SwiftService)
	for n, s := range t.Services {
		_service := &SwiftService{}
		_service.Name = AssembleServiceName(t.Namespaces[global.Swift], n)
		_service.Methods = make([]*SwiftMethod, 0)

		for mn, m := range s.Methods {
			_Method := &SwiftMethod{}
			_Method.Name = mn
			if m.ReturnType == nil {
				_Method.ValueType = &SwiftType{
					Type:      Void,
					Name:      Void,
					InnerType: "",
				}
			} else {
				_Method.ValueType = str.GetSwiftType(m.ReturnType)
			}
			_Method.Fields = make([]*SwiftField, 0, 10)

			for _, f := range m.Arguments {
				_f := &SwiftField{}
				_f.Name = f.Name
				_f.Type = str.GetSwiftType(f.Type)
				_Method.Fields = append(_Method.Fields, _f)
			}
			_service.Methods = append(_service.Methods, _Method)
		}
		services[_service.Name] = _service
	}
	str.SwiftThriftMap.Services = services
}

// GetSwiftType 通过 Field 获取 Type
func (str *SwiftThriftReader) GetSwiftType(t *parser.Type) *SwiftType {

	st := &SwiftType{}

	if b, tn := str.IsPlainType(t); b {
		st.Name = tn
		st.Type = PlainType
		st.InnerType = ""
		return st
	}

	if b, tn := str.IsEnumType(t); b {
		st.Name = tn
		st.Type = EnumType
		st.InnerType = ""
		return st
	}

	if b, tn := str.IsCustomerType(t); b {
		st.Name = tn
		st.Type = CustomerType
		st.InnerType = ""
		return st
	}

	if b, tn, innerType := str.IsListType(t); b {
		st.Name = tn
		st.Type = ListType
		st.InnerType = innerType
		return st
	}

	return st
}

// IsPlainType 如果是基本数据类型，返回该类型名
func (str *SwiftThriftReader) IsPlainType(t *parser.Type) (bool, string) {
	n := t.Name
	switch n {
	case global.TTI16:
		return true, TypeMapping[global.TTI16]
	case global.TTI32:
		return true, TypeMapping[global.TTI32]
	case global.TTI64:
		return true, TypeMapping[global.TTI64]
	case global.TTDouble:
		return true, TypeMapping[global.TTDouble]
	case global.TTBool:
		return true, TypeMapping[global.TTBool]
	case global.TTString:
		return true, TypeMapping[global.TTString]
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
				return true, AssembleName(thrift.Namespaces[global.Swift], s.Name)
			}
		}
	}

	if len(components) == 2 {
		for k, t := range thrifts {
			f := strings.Split(filepath.Base(k), ".")[0]
			if components[0] == f {
				for _, s := range t.Enums {
					if s.Name == components[1] {
						return true, AssembleName(t.Namespaces[global.Swift], components[1])
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
				return true, AssembleName(thrift.Namespaces[global.Swift], s.Name)
			}
		}
	}

	if len(components) == 2 {
		for k, t := range thrifts {
			f := strings.Split(filepath.Base(k), ".")[0]
			if components[0] == f {
				for _, s := range t.Structs {
					if s.Name == components[1] {
						return true, AssembleName(t.Namespaces[global.Swift], components[1])
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
