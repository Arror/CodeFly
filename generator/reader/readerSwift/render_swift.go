package readerSwift

import "fmt"

// AssembleName 名称组装
func AssembleName(namespace string, name string) string {
	return fmt.Sprintf("%s%s", namespace, name[1:])
}

// AssembleServiceName Service名称组装
func AssembleServiceName(namespace string, name string) string {
	return fmt.Sprintf("%sService", name)
}

// DefaultValue Swift 字段默认值
func (str *SwiftStruct) DefaultValue(f *SwiftField) string {

	switch f.Type.Type {
	case ListType:
		return fmt.Sprintf(": [%s] = []", f.Type.InnerType)
	case EnumType:
		return fmt.Sprintf(": %s?", f.Type.Name)
	case PlainType:
		switch f.Type.Name {
		case STInt16, STInt, STInt64:
			return fmt.Sprintf(": %s = 0", f.Type.Name)
		case STDouble:
			return fmt.Sprintf(": %s = 0.0", f.Type.Name)
		case STBool:
			return fmt.Sprintf(": %s = false", f.Type.Name)
		case STString:
			return fmt.Sprintf(": %s?", f.Type.Name)
		default:
			return fmt.Sprintf(": %s?", f.Type.Name)
		}
	case CustomerType:
		return fmt.Sprintf(": %s?", f.Type.Name)
	default:
		return fmt.Sprintf(": %s?", f.Type.Name)
	}
}

// FromDict 从JSON中初始化
func (str *SwiftStruct) FromDict(f *SwiftField) string {
	switch f.Type.Type {
	case ListType:
		return fmt.Sprintf("[%s].fromJSON(json: dict[\"%s\"])", f.Type.InnerType, f.Name)
	case EnumType:
		return fmt.Sprintf("%s(code: dict[\"%s\"] as? Int)", f.Type.Name, f.Name)
	case PlainType:
		switch f.Type.Name {
		case STInt16, STInt, STInt64:
			return fmt.Sprintf("dict[\"%s\"] as? %s ?? 0", f.Name, f.Type.Name)
		case STDouble:
			return fmt.Sprintf("dict[\"%s\"] as? %s ?? 0.0", f.Name, f.Type.Name)
		case STBool:
			return fmt.Sprintf("dict[\"%s\"] as? %s ?? false", f.Name, f.Type.Name)
		case STString:
			return fmt.Sprintf("dict[\"%s\"] as? %s", f.Name, f.Type.Name)
		default:
			return fmt.Sprintf("%s.fromJSON(json: dict[\"%s\"])", f.Type.InnerType, f.Name)
		}
	case CustomerType:
		return fmt.Sprintf("%s.fromJSON(json: dict[\"%s\"])", f.Type.Name, f.Name)
	default:
		return fmt.Sprintf("%s.fromJSON(json: dict[\"%s\"])", f.Type.Name, f.Name)
	}
}

// ToDict 创建JSON
func (str *SwiftStruct) ToDict(f *SwiftField) string {

	switch f.Type.Type {
	case ListType, CustomerType:
		return fmt.Sprintf("self.%s.toJSON()", f.Name)
	case EnumType:
		return fmt.Sprintf("self.%s.rawValue ?? 0", f.Name)
	case PlainType:
		switch f.Type.Name {
		case STInt16, STInt, STInt64:
			return fmt.Sprintf("self.%s ?? 0", f.Name)
		case STDouble:
			return fmt.Sprintf("self.%s ?? 0.0", f.Name)
		case STBool:
			return fmt.Sprintf("self.%s ?? false", f.Name)
		case STString:
			return fmt.Sprintf("self.%s ?? \"\"", f.Name)
		default:
			return fmt.Sprintf("self.%s.toJSON()", f.Name)
		}
	default:
		return fmt.Sprintf("self.%s.toJSON()", f.Name)
	}
}

// ReturnType 获取方法的返回值
func (m *SwiftMethod) ReturnType() string {

	switch m.ValueType.Type {
	case ListType:
		return fmt.Sprintf("[%s]", m.ValueType.InnerType)
	case Void:
		return ""
	default:
		return m.ValueType.Name
	}
}

// GetParam 拼接参数
func (f *SwiftField) GetParam() string {

	if f.Type.Type == ListType {
		return fmt.Sprintf("%s: [%s]", f.Name, f.Type.InnerType)
	}

	return fmt.Sprintf("%s: %s", f.Name, f.Type.Name)
}
