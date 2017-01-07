package compiler

import (
	"os"
	"sync"

	"github.com/samuel/go-thrift/parser"

	"CodeFly/context"
)

// Compiler Compiler interface
type Compiler interface {
	genEnumCode(ctx *context.Context, e *parser.Enum)
	genStructCode(ctx *context.Context, s *parser.Struct)
	genServiceCode(ctx *context.Context, s *parser.Service)
}

// GenCode Generate code
func GenCode(ctx *context.Context, compiler Compiler) {

	if err := os.MkdirAll(ctx.Output, 0755); err != nil {
		panic(err.Error())
	}

	wg := sync.WaitGroup{}

	wg.Add(1)
	go func() {
		defer wg.Done()
		for _, e := range ctx.Thrift.Enums {
			compiler.genEnumCode(ctx, e)
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		for _, s := range ctx.Thrift.Structs {
			compiler.genStructCode(ctx, s)
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		for _, s := range ctx.Thrift.Services {
			compiler.genServiceCode(ctx, s)
		}
	}()

	wg.Wait()
}
