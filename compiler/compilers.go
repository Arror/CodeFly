package compiler

import (
	"fmt"
	"strings"

	"github.com/Arror/CodeFly/context"
)

const swift = "swift"

// Compiler Compiler interface
type Compiler interface {
	compile(ctx context.Context)
}

func compilerOfLang(lang string) (Compiler, error) {

	switch lang {
	case swift:
		return &SwiftCompiler{}, nil
	}

	return nil, fmt.Errorf("compiler of %s not found", lang)
}

// Compile Compile code
func Compile(ctx context.Context) error {

	compiler, err := compilerOfLang(strings.ToLower(ctx.Lang))
	if err != nil {
		return err
	}

	err = ctx.MakeOutputFolder()
	if err != nil {
		return err
	}

	compiler.compile(ctx)

	return nil
}
