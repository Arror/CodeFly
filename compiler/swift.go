package compiler

import (
	"github.com/samuel/go-thrift/parser"

	"CodeFly/context"
	"CodeFly/writer"
)

// SwiftCompiler Swift Code Compiler
type SwiftCompiler struct{}

// Swift Swift
type Swift struct {
	Ctx *context.Context
}

// SwiftEnum Swift Enum
type SwiftEnum struct {
	*Swift
	*parser.Enum
}

// SwiftStruct Swift Struct
type SwiftStruct struct {
	*Swift
	*parser.Struct
}

// SwiftService Swift Service
type SwiftService struct {
	*Swift
	*parser.Service
}

var swift = &Swift{}

func (sc *SwiftCompiler) genEnumCode(ctx *context.Context, e *parser.Enum) {

	swift.Ctx = ctx

	swiftEnum := &SwiftEnum{}
	swiftEnum.Swift = swift
	swiftEnum.Enum = e

	writer.WriteFile("", ctx.EmunTemplate, ctx.EnumTemplateName, swiftEnum)
}

func (sc *SwiftCompiler) genStructCode(ctx *context.Context, s *parser.Struct) {

	swift.Ctx = ctx

	swiftStruct := &SwiftStruct{}
	swiftStruct.Swift = swift
	swiftStruct.Struct = s

	writer.WriteFile("", ctx.StructTemplate, ctx.StructTemplateName, nil)
}

func (sc *SwiftCompiler) genServiceCode(ctx *context.Context, s *parser.Service) {

	swift.Ctx = ctx

	swiftService := &SwiftService{}
	swiftService.Swift = swift
	swiftService.Service = s

	writer.WriteFile("", ctx.ServiceTemplate, ctx.ServiceTemplateName, nil)
}
