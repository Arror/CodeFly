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
		return fmt.Sprintf(": [%s]?", f.Type.InnerType)
	default:
		return fmt.Sprintf(": %s?", f.Type.Name)
	}
}

// FromDict 从JSON中初始化
func (es *EmbenStruct) FromDict(f *parser.SwiftField) string {
	switch f.Type.Type {
	case parser.ListType:
		return fmt.Sprintf("[%s](json: dict[\"%s\"])", f.Type.InnerType, f.Name)
	default:
		return fmt.Sprintf("%s(json: dict[\"%s\"])", f.Type.Name, f.Name)
	}
}

func toDict(f *parser.SwiftField, isNeedSelf bool) string {
	return fmt.Sprintf("self.%s?.toJSON()", f.Name)
}

// ToDict 创建JSON
func (es *EmbenStruct) ToDict(f *parser.SwiftField) string {
	return toDict(f, true)
}

// ToDict 获取方法的返回值
func (es *EmbenService) ToDict(f *parser.SwiftField) string {
	return toDict(f, false)
}

// ReturnType 获取方法的返回值
func (es *EmbenService) ReturnType(m *parser.SwiftMethod) string {

	switch m.ValueType.Type {
	case parser.ListType:
		return fmt.Sprintf("[%s]", m.ValueType.InnerType)
	case parser.Void:
		return ""
	default:
		return m.ValueType.Name
	}
}

// GetParam 拼接参数
func (es *EmbenService) GetParam(f *parser.SwiftField) string {

	if f.Type.Type == parser.ListType {
		return fmt.Sprintf("%s: [%s]", f.Name, f.Type.InnerType)
	}

	return fmt.Sprintf("%s: %s", f.Name, f.Type.Name)
}
