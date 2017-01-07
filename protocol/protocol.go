package protocol

import (
	"CodeFly/global"
	"CodeFly/lang/swift"
)

// Generator Generator protocol
type Generator interface {
	Generate()
}

// GeneratorMapping generator mapping
var GeneratorMapping = map[string]Generator{
	global.Swift: &swift.GenContext{},
}
