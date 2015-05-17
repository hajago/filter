package main

import (
	"fmt"
	"os"

	"github.com/codegangsta/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "filter"
	app.Usage = "extracts text of a document"
	app.Version = "0.0.1"
	app.Author = "dave"
	app.Email = "goodoi09@gmail.com"
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "output, o",
			Usage: "output file",
		},
	}

	app.Action = func(c *cli.Context) {
		args := c.Args()
		if len(args) != 1 {
			cli.ShowAppHelp(c)
			return
		}
		doc, err := detect(args[0])
		if err != nil {
			cli.ShowAppHelp(c)
			return
		}
		defer doc.Close()

		outF := os.Stdout
		if c.String("output") != "" {
			var err error
			if outF, err = os.OpenFile(c.String("output"), os.O_WRONLY|os.O_CREATE, 0666); err != nil {
				fmt.Printf("filter error: %s\n", err.Error())
				return
			}
			defer outF.Close()
		}
		if err := doc.Filter(outF); err != nil {
			fmt.Printf("filter error: %s\n", err.Error())
			return
		}
	}

	app.Run(os.Args)
}
