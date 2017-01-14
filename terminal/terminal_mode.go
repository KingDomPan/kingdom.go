package main

import (
	"fmt"

	"golang.org/x/crypto/ssh/terminal"
)

func main() {
	// 实际上就是读写termios数据结构
	// 修改I Mode L Mode
	oldState, err := terminal.MakeRaw(0)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%v", oldState)
	defer terminal.Restore(0, oldState)
}
