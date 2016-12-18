package main

import (
	"os"

	"github.com/urfave/cli"
)

// cli_run [global options] command [command options] [arguments...]
func main() {
	cli.NewApp().Run(os.Args)
}
