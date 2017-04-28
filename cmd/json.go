package cmd

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/samuel/go-thrift/parser"
	"github.com/urfave/cli"

	"github.com/Arror/CodeFly/compiler"
	"github.com/Arror/CodeFly/context"
)

// JSONCommand Json generate command
var JSONCommand = cli.Command{
	Name:      "json",
	ShortName: "json",
	Usage:     "Command of generate the target language code.",
	Flags: []cli.Flag{
		cli.StringFlag{
			Name:        "lang, l",
			Usage:       "target language",
			Destination: &lang,
		},
		cli.StringFlag{
			Name:        "input, i",
			Usage:       "input thrift file path",
			Destination: &input,
		},
		cli.StringFlag{
			Name:        "output, o",
			Usage:       "file output path",
			Destination: &output,
		},
	},
	Action: func(c *cli.Context) error {

		lang, err := verifyLanguage()
		if err != nil {
			log.Fatalln(err.Error())
		}

		input, err := verifyInputPath()
		if err != nil {
			log.Fatalln(err.Error())
		}

		output, err := verifyOutputPath()
		if err != nil {
			log.Fatalln(err.Error())
		}

		p := parser.Parser{}
		ts, _, err := p.ParseFile(input)
		if err != nil {
			log.Fatalln(err.Error())
		}

		for n, t := range ts {
			if t.Namespaces[lang] == "" {
				log.Fatalln("%s language namespace not found in %s.thrift", lang, n)
			}
		}

		ctx := context.CreateContext(lang, input, output, ts)

		err = compiler.Compile(ctx)
		if err != nil {
			log.Fatalln(err.Error())
		}

		return nil
	},
}

var (
	lang   string
	input  string
	output string
)

func verifyLanguage() (string, error) {

	language := strings.ToLower(lang)

	switch language {
	case "swift":
		return language, nil
	default:
		return "", fmt.Errorf("unsupported language: %s", language)
	}
}

func verifyInputPath() (string, error) {

	if input == "" {
		return "", fmt.Errorf("thrift file input path is empty")
	}

	p, err := filepath.Abs(input)
	if err != nil {
		return "", fmt.Errorf("thrift file not exist")
	}

	info, err := os.Stat(p)
	if err != nil {
		return "", fmt.Errorf("thrift file not exist")
	}

	if info.IsDir() {
		return "", fmt.Errorf("input is a directory")
	}

	return p, nil
}

func verifyOutputPath() (string, error) {

	if output == "" {
		return "", fmt.Errorf("output path is empty")
	}

	p, err := filepath.Abs(output)
	if err != nil {
		return "", fmt.Errorf("output path error")
	}

	return p, nil
}
