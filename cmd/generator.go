package cmd

import "github.com/urfave/cli"

// Gen Code generate command
var Gen = cli.Command{
	Name:      "gen",
	ShortName: "g",
	Usage:     "Code generator.",
	UsageText: "Generate target language file using thrift file",
	Flags: []cli.Flag{
		cli.StringFlag{
			Name:  "lang, l",
			Usage: "Target language name",
		},
		cli.StringFlag{
			Name:  "input, i",
			Usage: "Input thrift file path",
		},
		cli.StringFlag{
			Name:  "output, o",
			Usage: "Target language file output path",
		},
	},
	Action: func(c *cli.Context) error {
		return nil
	},
}
