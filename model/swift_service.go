package model

import "fmt"

// SwiftService Swift Service类型定义
type SwiftService struct {
	Name    string
	Methods []*SwiftMethod
}

// ToDict 渲染方法的返回值
func (s *SwiftService) ToDict(f *SwiftField) string {
	return fmt.Sprintf("%s?.toJSON()", f.Name)
}

// ReturnType 获取方法的返回类型
func (s *SwiftService) ReturnType(m *SwiftMethod) string {

	switch m.ValueType.Type {
	case ListType:
		return fmt.Sprintf("[%s]?", m.ValueType.InnerType)
	case Void:
		return ""
	default:
		return fmt.Sprintf("%s?", m.ValueType.Name)
	}
}

// GetParam 拼接参数
func (s *SwiftService) GetParam(f *SwiftField) string {

	if f.Type.Type == ListType {
		return fmt.Sprintf("%s: [%s]?", f.Name, f.Type.InnerType)
	}

	return fmt.Sprintf("%s: %s?", f.Name, f.Type.Name)
}
