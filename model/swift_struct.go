package model

import "fmt"

// SwiftStruct Swift Struct类型定义
type SwiftStruct struct {
	Name   string
	Fields []*SwiftField
}

// ValueTypeFormat Swift 字段类型
func (s *SwiftStruct) ValueTypeFormat(f *SwiftField, optional bool) string {

	var optionalString string

	if optional {
		optionalString = "?"
	} else {
		optionalString = ""
	}

	var typeString string

	switch f.Type.Type {
	case ListType:
		typeString = fmt.Sprintf("[%s]", f.Type.InnerType)
	default:
		typeString = fmt.Sprintf("%s", f.Type.Name)
	}

	return typeString + optionalString
}

// ToDict 创建JSON
func (s *SwiftStruct) ToDict(f *SwiftField) string {
	return fmt.Sprintf("self.%s?.toJSON()", f.Name)
}
