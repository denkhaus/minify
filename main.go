package main

import (
	"bufio"
	"io"
	"os"

	"github.com/codegangsta/cli"
	"github.com/dchest/cssmin"
	"github.com/dchest/jsmin"
	"github.com/denkhaus/htmlmin"
	"github.com/juju/errors"
)

func process(in io.Reader, out io.Writer, fn func(buf []byte) ([]byte, error)) error {
	scanner := bufio.NewScanner(in)
	for scanner.Scan() {
		min, err := fn(scanner.Bytes())
		if err != nil {
			return errors.Annotate(err, "minify")
		}
		_, err = out.Write(min)
		if err != nil {
			return errors.Annotate(err, "write output")
		}
	}

	return nil
}

func main() {
	app := cli.NewApp()
	app.Name = "minify"

	app.Usage = "A minifier that  accepts input from Stdin and puts the result to Stdout"
	app.Flags = []cli.Flag{
		cli.BoolFlag{
			Name:  "html",
			Usage: "minify html",
		},
		cli.BoolFlag{
			Name:  "css",
			Usage: "minify css",
		},
		cli.BoolFlag{
			Name:  "js",
			Usage: "minify js",
		},
	}

	app.Action = cli.ActionFunc(func(ctx *cli.Context) error {
		if ctx.Bool("html") {
			return process(os.Stdin, os.Stdout, func(buf []byte) ([]byte, error) {
				return htmlmin.Minify(buf, nil)
			})
		}
		if ctx.Bool("css") {
			return process(os.Stdin, os.Stdout, func(buf []byte) ([]byte, error) {
				return cssmin.Minify(buf), nil
			})
		}
		if ctx.Bool("js") {
			return process(os.Stdin, os.Stdout, func(buf []byte) ([]byte, error) {
				return jsmin.Minify(buf)
			})
		}

		return errors.New("error: no flag defined")
	})

	app.Run(os.Args)
}
