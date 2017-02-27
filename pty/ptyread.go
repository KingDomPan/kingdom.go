package main

import (
	"io"
	"os"
	"os/exec"
)

func main() {
	c := exec.Command("/usr/local/bin/zsh")

	file, err := os.Create("panqd.log")
	if err != nil {
		panic("can not create filename: " + "panqd.log" + " ,err is: " + err.Error())
	}
	defer file.Close()

	mw := io.MultiWriter(os.Stdout, file)

	c.Stdin = os.Stdin
	c.Stdout = mw

	if err := c.Start(); err != nil {
		panic(err)
	}

	if err := c.Wait(); err != nil {
		panic(err)
	}
}
