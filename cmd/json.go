package cmd

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/samuel/go-thrift/parser"
	"github.com/urfave/cli"

	"github.com/Arror/CodeFly/context"
	"github.com/Arror/CodeFly/generator"
)

var (
	lang   string
	input  string
	output string

	tp = parser.Parser{}
)

const (
	swift = "swift"
)

var jsonCommand = cli.Command{
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

		if err := verifyCommondArgs(); err != nil {
			log.Fatalln(err.Error())
		}

		ts, _, err := tp.ParseFile(input)
		if err != nil {
			log.Fatalln(err.Error())
		}

		for n, t := range ts {
			if t.Namespaces[lang] == "" {
				log.Fatalln("%s language namespace not found in %s.thrift", lang, n)
			}
		}

		ctx := context.Create(lang, input, output, ts)

		if err = generator.Generate(ctx); err != nil {
			log.Fatalln(err.Error())
		}

		return nil
	},
}

func verifyCommondArgs() error {

	language := strings.ToLower(lang)

	switch language {
	case swift:
		lang = language
	default:
		return fmt.Errorf("unsupported language: %s", language)
	}

	if input == "" {
		return fmt.Errorf("input path is empty")
	}
	ip, err := filepath.Abs(input)
	if err != nil {
		return fmt.Errorf("file not exist")
	}
	info, err := os.Stat(ip)
	if err != nil {
		return fmt.Errorf("file not exist")
	}
	if info.IsDir() {
		return fmt.Errorf("input is a directory")
	}
	input = ip

	if output == "" {
		output = "./"
	}
	op, err := filepath.Abs(output)
	if err != nil {
		return fmt.Errorf("output path error")
	}
	output = op

	return nil
}
