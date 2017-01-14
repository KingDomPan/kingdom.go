package main

import (
	"fmt"
	"os"

	"golang.org/x/crypto/ssh/terminal"
)

func main() {
	ff, _ := os.OpenFile("/dev/ttys002", os.O_RDWR, 0)
	defer ff.Close()

	ff3, _ := os.OpenFile("/dev/ttys003", os.O_RDWR, 0)
	defer ff3.Close()

	fmt.Println(terminal.IsTerminal(int(ff.Fd())))
	fmt.Println(terminal.IsTerminal(int(os.Stdout.Fd())))
	fmt.Println(terminal.IsTerminal(int(os.Stderr.Fd())))
	fmt.Println(terminal.IsTerminal(int(os.Stdin.Fd())))

	w1, h1, _ := terminal.GetSize(int(ff.Fd()))
	w2, h2, _ := terminal.GetSize(int(os.Stdout.Fd()))
	w3, h3, _ := terminal.GetSize(int(os.Stderr.Fd()))
	w4, h4, _ := terminal.GetSize(int(os.Stdin.Fd()))

	println(w1, h1)
	println(w2, h2)
	println(w3, h3)
	println(w4, h4)

	println("---------------")

	println(ff.Fd())
	println(ff3.Fd())
	println(os.Stdout.Fd())
	println(os.Stderr.Fd())
	println(os.Stdin.Fd())
}
