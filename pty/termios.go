package main

import (
	"io"
	"os"
	"os/exec"
	"syscall"

	"github.com/pkg/term/termios"
)

func main() {

	c := exec.Command("/usr/local/bin/zsh")

	ptmx, tty, err := termios.Pty()
	if err != nil {
		panic(err)
	}
	defer ptmx.Close()
	defer tty.Close()

	c.Stdout = tty
	c.Stdin = tty
	c.Stderr = tty
	c.SysProcAttr = &syscall.SysProcAttr{Setctty: true, Setsid: true}

	err = c.Start()
	if err != nil {
		panic(err)
	}

	go func() {
		io.Copy(ptmx, os.Stdin)
	}()

	go func() {
		io.Copy(os.Stdout, ptmx)
	}()

	err = c.Wait()
	if err != nil {
		panic(err)
	}
}
