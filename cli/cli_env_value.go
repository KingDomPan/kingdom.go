package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:   "lang, l",
			Value:  "english",
			Usage:  "language for the greeting",
			EnvVar: "PATH", // 优先级别大于Value
		},
	}

	app.Action = func(c *cli.Context) error {
		fmt.Println(c.String("lang"))
		return nil
	}
	app.Run(os.Args)
}
