package generator

import (
	"fmt"
	"os"
	"strings"

	"github.com/Arror/CodeFly/context"
)

type generator interface {
	generate(ctx *context.Context)
}

var generatorMapping = make(map[string]generator)

func enroll(generator generator, lang string) {
	generatorMapping[strings.ToLower(lang)] = generator
}

// Generate generate code with context
func Generate(ctx *context.Context) error {

	lang := strings.ToLower(ctx.Args.Lang)

	generator, exist := generatorMapping[lang]
	if !exist {
		return fmt.Errorf("generator for %s language not found", lang)
	}

	err := os.MkdirAll(ctx.Args.Output, 0755)
	if err != nil {
		return err
	}

	generator.generate(ctx)

	return nil
}
