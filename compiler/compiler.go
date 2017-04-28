package compiler

import (
	"fmt"
	"os"
	"strings"

	"github.com/Arror/CodeFly/context"
)

type compiler interface {
	compile(ctx *context.Context)
}

var compilerMapping = make(map[string]compiler)

func register(complier compiler, lang string) {
	compilerMapping[strings.ToLower(lang)] = complier
}

// Compile Compile code
func Compile(ctx *context.Context) error {

	lang := strings.ToLower(ctx.Lang)

	compiler, exist := compilerMapping[lang]
	if !exist {
		return fmt.Errorf("compliler of %s not found", lang)
	}

	err := os.MkdirAll(ctx.Output, 0755)
	if err != nil {
		return err
	}

	compiler.compile(ctx)

	return nil
}
