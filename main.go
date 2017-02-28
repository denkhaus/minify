package main

import (
	"bufio"
	"os"

	"github.com/codegangsta/cli"
	"github.com/denkhaus/htmlmin"
	"github.com/juju/errors"
)

func main() {
	app := cli.NewApp()
	app.Name = "minify"

	app.Usage = "A minifier"
	app.Flags = []cli.Flag{
		cli.BoolFlag{
			Name:  "html",
			Usage: "minify html",
		},
	}

	app.Action = cli.ActionFunc(func(ctx *cli.Context) error {
		if ctx.Bool("html") {
			scanner := bufio.NewScanner(os.Stdin)
			for scanner.Scan() {
				out, err := htmlmin.Minify(scanner.Bytes(), nil)
				if err != nil {
					return errors.Annotate(err, "minify html")
				}
				_, err = os.Stdout.Write(out)
				if err != nil {
					return errors.Annotate(err, "write to stdout")
				}
			}

			return nil
		}

		return errors.New("no flag defined")
	})

	app.Run(os.Args)
}
