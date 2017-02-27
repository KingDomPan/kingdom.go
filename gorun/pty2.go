package main

import (
	"io"
	"os"
	"os/exec"

	"github.com/kr/pty"
)

func main() {
	c := exec.Command("/bin/bash")
	f, err := pty.Start(c)
	if err != nil {
		panic(err)
	}
	ch := make(chan struct{}, 1)
	go func() {
		io.Copy(os.Stdout, f)
		ch <- struct{}{}
	}()
	go func() {
		io.Copy(f, os.Stdin)
		ch <- struct{}{}
	}()
	<-ch
}
