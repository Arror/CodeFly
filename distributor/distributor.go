package distributor

import (
	"CodeFly/generator"
	"CodeFly/model"

	"github.com/samuel/go-thrift/parser"
)

// Distribute 任务分配者
func Distribute(ts map[string]*parser.Thrift, genInfo *model.GenerateCommandInfo) {

	switch genInfo.Lang {
	case model.Swift:
		stc := &model.SwiftThriftComponents{}
		stc.InitWith(ts, genInfo)
		generator.GeneratSwiftCode(stc)
	}
}
