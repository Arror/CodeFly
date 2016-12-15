package maker

import (
	"fmt"

	"CodeFly/global"
	"CodeFly/lang/swift"
	"CodeFly/parameter"
	"CodeFly/protocol"
)

// MakeGenerator Make generator
func MakeGenerator(param *parameter.Parameter) (protocol.Generator, error) {

	switch param.Lang {
	case global.Swift:
		return &swift.Generator{}, nil
	default:
		return nil, fmt.Errorf("Create %s language generator failed", param.Lang)
	}
}
