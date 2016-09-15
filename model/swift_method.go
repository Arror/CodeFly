package model

// SwiftMethod Swift Method类型定义
type SwiftMethod struct {
	Name      string
	URL       string
	Fields    []*SwiftField
	ValueType *SwiftType
}
