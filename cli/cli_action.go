package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "this is a app name"
	app.Usage = "this is a usage information"
	app.Action = func(c *cli.Context) error {
		fmt.Println("this is inside app action method")
		return nil
	}
	app.Run(os.Args)
}
