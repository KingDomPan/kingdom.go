package main

import (
	"fmt"
	"os"
	"syscall"
	"unsafe"
)

func main() {

	file, _ := os.Open("/dev/ttys003")

	var devOldState syscall.Termios

	if _, _, err := syscall.Syscall6(syscall.SYS_IOCTL, uintptr(file.Fd()), syscall.TIOCGETA, uintptr(unsafe.Pointer(&devOldState)), 0, 0, 0); err != 0 {
		panic(err)
	}

	fmt.Println(devOldState)
}
