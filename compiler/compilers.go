package compiler

import (
	"fmt"
	"log"
	"path/filepath"
	"strings"

	"github.com/Arror/CodeFly/context"
)

const swift = "swift"

// Compiler Compiler interface
type Compiler interface {
	compile(ctx context.Context)
}

func compilerOfLang(lang string) Compiler {

	switch lang {
	case swift:
		return &SwiftCompiler{}
	}

	return nil
}

// Compile Compile code
func Compile(ctx context.Context) error {

	if compiler := compilerOfLang(strings.ToLower(ctx.Lang)); compiler != nil {

		compiler.compile(ctx)

		return nil
	}

	return fmt.Errorf("Can't find compiler for language: %s", ctx.Lang)
}

func assembleFilePath(op string, fn string) string {

	p, err := filepath.Abs(filepath.Join(op, fn))

	if err != nil {
		log.Fatalln(err.Error())
	}

	return p
}
