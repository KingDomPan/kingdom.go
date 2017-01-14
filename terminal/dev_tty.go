package main

import (
	"fmt"
	"os"
)

func main() {
	content := []byte("a")

	ff, _ := os.OpenFile("/dev/tty", os.O_RDWR, 0)
	defer ff.Close()

	fmt.Println(string(content)) // 输出重定向
	ff.Write(content)            // 写到tty
}
