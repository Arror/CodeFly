package parser

import (
	"fmt"
	"path/filepath"
	"strings"

	p "github.com/samuel/go-thrift/parser"
)

var typeMapping = map[string]string{
	TTI16:    STInt,
	TTI32:    STInt,
	TTI64:    STInt64,
	TTDouble: STDouble,
	TTBool:   STBool,
	TTString: STString,
}

// Parser 解析生成SwiftThriftComponents
func Parser(ts map[string]*p.Thrift, ip string, op string) *SwiftThriftComponents {

	stc := &SwiftThriftComponents{
		Thrifts:     ts,
		Thrift:      ts[ip],
		InputPath:   ip,
		OutputPath:  op,
		SwiftThrift: &SwiftThrift{},
	}

	parser(stc)

	return stc
}

func parser(stc *SwiftThriftComponents) {

	t := stc.Thrift
	st := stc.SwiftThrift

	structs := make(map[string]*SwiftStruct)
	for sn, s := range t.Structs {
		ss := &SwiftStruct{}
		ss.Name = assembleNamespace(t.Namespaces[Swift], sn)
		ss.Fields = make([]*SwiftField, 0, 10)

		for _, f := range s.Fields {
			sf := &SwiftField{}
			sf.Name = f.Name
			sf.Type = stc.getSwiftType(f.Type)
			ss.Fields = append(ss.Fields, sf)
		}
		structs[ss.Name] = ss
	}
	st.Structs = structs

	enums := make(map[string]*SwiftEnum)
	for en, e := range t.Enums {
		se := &SwiftEnum{}
		se.Name = assembleNamespace(t.Namespaces[Swift], en)
		se.Cases = make([]*SwiftEnumCase, 0, 10)

		for _, v := range e.Values {
			sec := &SwiftEnumCase{}
			sec.Name = v.Name
			sec.Value = v.Value
			sec.Type = &SwiftType{
				Type:      EnumType,
				Name:      se.Name,
				InnerType: "",
			}
			se.Cases = append(se.Cases, sec)
		}
		enums[se.Name] = se
	}
	st.Enums = enums

	services := make(map[string]*SwiftService)
	for sn, s := range t.Services {
		ss := &SwiftService{}
		ss.Name = assembleServiceName(sn)
		ss.Methods = make([]*SwiftMethod, 0)

		for mn, m := range s.Methods {
			sm := &SwiftMethod{}
			sm.Name = mn
			sm.URL = fmt.Sprintf("%s/%s", sn, mn)
			if m.ReturnType == nil {
				sm.ValueType = &SwiftType{
					Type:      Void,
					Name:      Void,
					InnerType: "",
				}
			} else {
				sm.ValueType = stc.getSwiftType(m.ReturnType)
			}
			sm.Fields = make([]*SwiftField, 0, 10)

			for _, f := range m.Arguments {
				sf := &SwiftField{}
				sf.Name = f.Name
				sf.Type = stc.getSwiftType(f.Type)
				sm.Fields = append(sm.Fields, sf)
			}
			ss.Methods = append(ss.Methods, sm)
		}
		services[ss.Name] = ss
	}
	st.Services = services
}

func assembleNamespace(namespace string, name string) string {
	return fmt.Sprintf("%s%s", namespace, name[1:])
}

func assembleServiceName(name string) string {
	return fmt.Sprintf("%sService", name)
}

func (stc *SwiftThriftComponents) getSwiftType(t *p.Type) *SwiftType {

	st := &SwiftType{}

	if tn, b := stc.isPlainType(t); b {
		st.Name = tn
		st.Type = PlainType
		st.InnerType = ""
		return st
	}

	if tn, b := stc.isEnumType(t); b {
		st.Name = tn
		st.Type = EnumType
		st.InnerType = ""
		return st
	}

	if tn, b := stc.isCustomerType(t); b {
		st.Name = tn
		st.Type = CustomerType
		st.InnerType = ""
		return st
	}

	if tn, innerType, b := stc.isListType(t); b {
		st.Name = tn
		st.Type = ListType
		st.InnerType = innerType
		return st
	}

	return st
}

func (stc *SwiftThriftComponents) isPlainType(t *p.Type) (string, bool) {
	n := t.Name
	switch n {
	case TTI16, TTI32:
		return typeMapping[TTI32], true
	case TTI64:
		return typeMapping[TTI64], true
	case TTDouble:
		return typeMapping[TTDouble], true
	case TTBool:
		return typeMapping[TTBool], true
	case TTString:
		return typeMapping[TTString], true
	default:
		return "", false
	}
}

func (stc *SwiftThriftComponents) isEnumType(t *p.Type) (string, bool) {

	thrift := stc.Thrift
	thrifts := stc.Thrifts

	n := t.Name
	components := strings.Split(n, ".")

	if len(components) == 1 {
		for _, s := range thrift.Enums {
			if s.Name == components[0] {
				return assembleNamespace(thrift.Namespaces[Swift], s.Name), true
			}
		}
	}

	if len(components) == 2 {
		for k, t := range thrifts {
			f := strings.Split(filepath.Base(k), ".")[0]
			if components[0] == f {
				for _, s := range t.Enums {
					if s.Name == components[1] {
						return assembleNamespace(t.Namespaces[Swift], components[1]), true
					}
				}
			}
		}
	}
	return "", false
}

func (stc *SwiftThriftComponents) isCustomerType(t *p.Type) (string, bool) {

	thrift := stc.Thrift
	thrifts := stc.Thrifts

	n := t.Name
	components := strings.Split(n, ".")

	if len(components) == 1 {
		for _, s := range thrift.Structs {
			if s.Name == components[0] {
				return assembleNamespace(thrift.Namespaces[Swift], s.Name), true
			}
		}
	}

	if len(components) == 2 {
		for k, t := range thrifts {
			f := strings.Split(filepath.Base(k), ".")[0]
			if components[0] == f {
				for _, s := range t.Structs {
					if s.Name == components[1] {
						return assembleNamespace(t.Namespaces[Swift], components[1]), true
					}
				}
			}
		}
	}
	return "", false
}

func (stc *SwiftThriftComponents) isListType(t *p.Type) (string, string, bool) {

	if t.Name == "list" {
		innerType := t.ValueType
		if tn, b := stc.isPlainType(innerType); b {
			return t.Name, tn, b
		}
		if tn, b := stc.isEnumType(innerType); b {
			return t.Name, tn, b
		}
		if tn, b := stc.isCustomerType(innerType); b {
			return t.Name, tn, b
		}
	}
	return "", "", false
}

// DefaultValue Swift 字段默认值
func (s *SwiftStruct) DefaultValue(f *SwiftField) string {

	switch f.Type.Type {
	case ListType:
		return fmt.Sprintf(": [%s]?", f.Type.InnerType)
	default:
		return fmt.Sprintf(": %s?", f.Type.Name)
	}
}

// FromDict 从JSON中初始化
func (s *SwiftStruct) FromDict(f *SwiftField) string {
	switch f.Type.Type {
	case ListType:
		return fmt.Sprintf("[%s](json: dict[\"%s\"])", f.Type.InnerType, f.Name)
	default:
		return fmt.Sprintf("%s(json: dict[\"%s\"])", f.Type.Name, f.Name)
	}
}

func toDict(f *SwiftField) string {
	return fmt.Sprintf("self.%s?.toJSON()", f.Name)
}

// ToDict 创建JSON
func (s *SwiftStruct) ToDict(f *SwiftField) string {
	return fmt.Sprintf("self.%s?.toJSON()", f.Name)
}

// ToDict 获取方法的返回值
func (s *SwiftService) ToDict(f *SwiftField) string {
	return fmt.Sprintf("%s?.toJSON()", f.Name)
}

// ReturnType 获取方法的返回值
func (s *SwiftService) ReturnType(m *SwiftMethod) string {

	switch m.ValueType.Type {
	case ListType:
		return fmt.Sprintf("[%s]?", m.ValueType.InnerType)
	case Void:
		return ""
	default:
		return m.ValueType.Name + "?"
	}
}

// GetParam 拼接参数
func (s *SwiftService) GetParam(f *SwiftField) string {

	if f.Type.Type == ListType {
		return fmt.Sprintf("%s: [%s]?", f.Name, f.Type.InnerType)
	}

	return fmt.Sprintf("%s: %s?", f.Name, f.Type.Name)
}
