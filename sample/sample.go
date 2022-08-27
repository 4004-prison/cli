package main

import (
	"cli"
	"fmt"
	"os"
)

func main() {
	app := &cli.App{
		Name:  "123",
		Usage: "123",
		Commands: []*cli.Command{
			{
				Name: "run",
				Action: func(ctx *cli.Context) {
					fmt.Printf("%s\n", ctx.String("string"))
					fmt.Printf("%#v\n", ctx.Value("slicestring"))
				},
				Flags: []cli.Flag{
					&cli.String{
						Name: "string",
					},
					&cli.SliceString{
						Name: "slicestring",
					},
				},
			},
		},
	}
	app.Run(os.Args)
}
