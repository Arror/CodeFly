package cmd

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"CodeFly/compiler"
	"CodeFly/context"
	"CodeFly/global"

	"github.com/samuel/go-thrift/parser"
	"github.com/urfave/cli"
)

// JSONCommand Json generate command
var JSONCommand = cli.Command{
	Name:      "json",
	ShortName: "json",
	Usage:     "Command of generate the target language code.",
	Flags: []cli.Flag{
		cli.StringFlag{
			Name:        "lang, l",
			Usage:       "Target language.",
			Destination: &lang,
		},
		cli.StringFlag{
			Name:        "input, i",
			Usage:       "Input thrift file path.",
			Destination: &input,
		},
		cli.StringFlag{
			Name:        "output, o",
			Usage:       "File output path.",
			Destination: &output,
		},
	},
	Action: func(c *cli.Context) error {

		lang, err := checkLanguage()
		if err != nil {
			log.Fatalln(err.Error())
		}

		input, err := checkInputPath()
		if err != nil {
			log.Fatalln(err.Error())
		}

		output, err := checkOutputPath()
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
				log.Fatalln("%s language namespace info not found in %s.thrift", lang, n)
			}
		}

		ctx := context.Init(lang, input, output, ts)

		compiler.GenCode(ctx)

		return nil
	},
}

var lang string
var input string
var output string

func checkLanguage() (string, error) {

	if lang == "" {
		return "", fmt.Errorf("The target language name is empty")
	}

	switch lang {
	case global.Swift:
		return lang, nil
	default:
		return "", fmt.Errorf("Unsupported language")
	}
}

func checkInputPath() (string, error) {

	if input == "" {
		return "", fmt.Errorf("The thrift file input path is empty")
	}

	p, err := filepath.Abs(input)

	if err != nil {
		return "", fmt.Errorf("The input thrift file path not exist")
	}

	info, err := os.Stat(p)

	if err != nil {
		return "", fmt.Errorf("The input thrift file path not exist")
	}

	if info.IsDir() {
		return "", fmt.Errorf("The input path is a directory")
	}

	return p, nil
}

func checkOutputPath() (string, error) {

	if output == "" {
		return "", fmt.Errorf("File output path is empty")
	}

	p, err := filepath.Abs(output)

	if err != nil {
		return "", fmt.Errorf("File output path error")
	}

	return p, nil
}
