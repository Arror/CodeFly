package model

// SwiftThrift Swift Thrift类型定义
type SwiftThrift struct {
	Structs  map[string]*SwiftStruct
	Services map[string]*SwiftService
}
