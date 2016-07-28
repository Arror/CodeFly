package generator

import (
	"fmt"

	"CodeFly/parser"
)

// EmbenStruct SwiftStruct内嵌结构
type EmbenStruct struct {
	*parser.SwiftStruct
}

// EmbenService SwiftService内嵌结构
type EmbenService struct {
	*parser.SwiftService
}

// AssembleName 名称组装
func AssembleName(namespace string, name string) string {
	return fmt.Sprintf("%s%s", namespace, name[1:])
}

// AssembleServiceName Service名称组装
func AssembleServiceName(namespace string, name string) string {
	return fmt.Sprintf("%sService", name)
}

// DefaultValue Swift 字段默认值
func (es *EmbenStruct) DefaultValue(f *parser.SwiftField) string {

	switch f.Type.Type {
	case parser.ListType:
		return fmt.Sprintf(": [%s] = []", f.Type.InnerType)
	case parser.PlainType:
		switch f.Type.Name {
		case parser.STInt, parser.STInt64:
			return fmt.Sprintf(": %s = 0", f.Type.Name)
		case parser.STDouble:
			return fmt.Sprintf(": %s = 0.0", f.Type.Name)
		case parser.STBool:
			return fmt.Sprintf(": %s = false", f.Type.Name)
		case parser.STString:
			return fmt.Sprintf(": %s?", f.Type.Name)
		default:
			return fmt.Sprintf(": %s?", f.Type.Name)
		}
	case parser.CustomerType, parser.EnumType:
		return fmt.Sprintf(": %s?", f.Type.Name)
	default:
		return fmt.Sprintf(": %s?", f.Type.Name)
	}
}

// FromDict 从JSON中初始化
func (es *EmbenStruct) FromDict(f *parser.SwiftField) string {
	switch f.Type.Type {
	case parser.ListType:
		return fmt.Sprintf("[%s](json: dict[\"%s\"]) ?? []", f.Type.InnerType, f.Name)
	default:
		return fmt.Sprintf("%s(json: dict[\"%s\"])", f.Type.Name, f.Name)
	}
}

// ToDict 创建JSON
func (es *EmbenStruct) ToDict(f *parser.SwiftField) string {

	switch f.Type.Type {
	case parser.CustomerType, parser.EnumType:
		return fmt.Sprintf("self.%s?.toJSON()", f.Name)
	case parser.PlainType:
		switch f.Type.Name {
		case parser.STString:
			return fmt.Sprintf("self.%s?.toJSON()", f.Name)
		default:
			return fmt.Sprintf("self.%s.toJSON()", f.Name)
		}
	default:
		return fmt.Sprintf("self.%s.toJSON()", f.Name)
	}
}

// // ReturnType 获取方法的返回值
// func (m *SwiftMethod) ReturnType() string {

// 	switch m.ValueType.Type {
// 	case ListType:
// 		return fmt.Sprintf("[%s]", m.ValueType.InnerType)
// 	case Void:
// 		return ""
// 	default:
// 		return m.ValueType.Name
// 	}
// }

// // GetParam 拼接参数
// func (f *SwiftField) GetParam() string {

// 	if f.Type.Type == ListType {
// 		return fmt.Sprintf("%s: [%s]", f.Name, f.Type.InnerType)
// 	}

// 	return fmt.Sprintf("%s: %s", f.Name, f.Type.Name)
// }
