package compiler

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"text/template"

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

func exportFiles(fp string, t *template.Template, tplname string, data interface{}) {

	file, err := os.OpenFile(fp, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer file.Close()
	if err := t.ExecuteTemplate(file, tplname, data); err != nil {
		log.Fatal(err.Error())
	}
}

func assembleFilePath(op string, fn string) string {

	p, err := filepath.Abs(filepath.Join(op, fn))

	if err != nil {
		log.Fatalln(err.Error())
	}

	return p
}
