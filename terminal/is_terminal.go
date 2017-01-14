package main

import (
	"fmt"
	"os"

	"github.com/kr/pty"
	"github.com/mattn/go-isatty"
)

// 检测一个fd是不是一个terminal
// SYS_IOCTL 系统调用 TIOCGETA参数
// 底层是一个termio数据结构
func main() {
	// 检测标准输入是不是一个terminal
	if isatty.IsTerminal(os.Stdin.Fd()) {
		fmt.Println("Stdin Is Terminal")
	} else {
		fmt.Println("Stdin Is Not Terminal")
	}
	// 检测标准输出是不是一个terminal
	if isatty.IsTerminal(os.Stdout.Fd()) {
		fmt.Println("Stdout Is Terminal")
	} else {
		fmt.Println("Stdout Is Not Terminal")
	}
	// 检测错误输出是不是一个terminal
	if isatty.IsTerminal(os.Stderr.Fd()) {
		fmt.Println("Stderr Is Terminal")
	} else {
		fmt.Println("Stderr Is Not Terminal")
	}
	// 检测ptmx是不是terminal
	if isatty.IsTerminal(os.Stderr.Fd()) {
		fmt.Println("Ptmx Is Terminal")
	} else {
		fmt.Println("Ptmx Is Not Terminal")
	}

	ff, _ := os.OpenFile("/dev/ttys001", os.O_RDWR, 0)
	defer ff.Close()

	// 检测错误输出是不是一个terminal
	if isatty.IsTerminal(ff.Fd()) {
		fmt.Println("/dev/ttys001 Is Terminal")
	} else {
		fmt.Println("Is Not Terminal")
	}

	// TIOCGWINSZ 参数
	row, col, _ := pty.Getsize(ff)
	fmt.Println("/dev/ttys001", row, col)
}
