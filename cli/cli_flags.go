package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli"
)

// go run cli_flags.go --lang spanish panqd
func main() {
	app := cli.NewApp()

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "lang",
			Value: "english",
			Usage: "language for the you",
		},
	}

	app.Action = func(c *cli.Context) error {
		name := "Nefertiti"
		if c.NArg() > 0 {
			name = c.Args().Get(0)
		}
		if c.String("lang") == "spanish" {
			fmt.Println("spanish", name)
		}
		if c.String("lang") == "english" {
			fmt.Println("english", name)
		}
		return nil
	}

	app.Run(os.Args)
}
