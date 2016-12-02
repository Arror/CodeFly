package model

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/arrors/go-thrift/parser"
)

// SwiftThriftComponents 完整的Thrift信息
type SwiftThriftComponents struct {
	InputPath   string
	OutputPath  string
	Thrift      *parser.Thrift
	Thrifts     map[string]*parser.Thrift
	SwiftThrift *SwiftThrift
}

// InitWith SwiftThriftComponents 初始化
func (stc *SwiftThriftComponents) InitWith(ts map[string]*parser.Thrift, genInfo *GenerateCommandInfo) {

	stc.InputPath = genInfo.Input
	stc.OutputPath = genInfo.Output
	stc.Thrifts = ts
	stc.Thrift = ts[genInfo.Input]
	stc.SwiftThrift = &SwiftThrift{}

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

func (stc *SwiftThriftComponents) getSwiftType(t *parser.Type) *SwiftType {

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

func (stc *SwiftThriftComponents) isPlainType(t *parser.Type) (string, bool) {
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

func (stc *SwiftThriftComponents) isEnumType(t *parser.Type) (string, bool) {

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

func (stc *SwiftThriftComponents) isCustomerType(t *parser.Type) (string, bool) {

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

func (stc *SwiftThriftComponents) isListType(t *parser.Type) (string, string, bool) {

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
