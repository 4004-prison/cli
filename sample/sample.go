package main

import (
	"fmt"
	"os"

	"github.com/4004-prison/cli"
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
					fmt.Printf("%d\n", ctx.Value("int"))
					fmt.Printf("%v\n", ctx.Value("bool"))
					fmt.Printf("%#v\n", ctx.Value("sliceint"))
					fmt.Printf("%#v\n", ctx.String("abc"))
				},
				Flags: []cli.Flag{
					&cli.String{
						Name: "string",
					},
					&cli.SliceString{
						Name: "slicestring",
					},
					&cli.Int{
						Name: "int",
					},
					&cli.Bool{
						Name: "bool",
					},
					&cli.SliceInt{
						Name: "sliceint",
					},
				},
			},
		},
	}
	app.Run(os.Args)
}
