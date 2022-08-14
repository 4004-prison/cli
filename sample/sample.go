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
					fmt.Println(ctx.String("t"))
				},
				Flags: []cli.Flag{
					&cli.String{
						Name: "t",
					},
				},
			},
		},
	}
	app.Run(os.Args)
}
