package main

import (
	"fmt"

	"golang.org/x/crypto/ssh/terminal"
)

func main() {
	oldState, err := terminal.MakeRaw(0)
	if err != nil {
		panic(err)
	}
	fmt.Println(oldState)
	defer terminal.Restore(0, oldState)
}
