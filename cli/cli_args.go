package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli"
)

// c.Args() go run cli_args.go panqd panqd2 panqd3
func main() {
	app := cli.NewApp()
	app.Action = func(c *cli.Context) error {
		fmt.Printf("hello %q\n", c.Args().Get(0)) // hello panqd
		for _, args := range c.Args() {
			fmt.Println(args) // panqd panqd2 panqd3
		}
		return nil
	}
	app.Run(os.Args)
}
