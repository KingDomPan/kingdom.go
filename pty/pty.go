package main

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"time"

	"github.com/kr/pty"
)

func main() {
	c := exec.Command("ping", "127.0.0.1")
	f, err := pty.Start(c)
	if err != nil {
		panic(err)
	}

	ff, _ := os.OpenFile("/dev/ttys001", os.O_RDWR, 0)
	defer ff.Close()

	row, col, _ := pty.Getsize(ff)
	fmt.Println(row, col)

	go func() {
		f.Write([]byte("foo\n"))
		f.Write([]byte("bar\n"))
		f.Write([]byte("baz\n"))
		f.Write([]byte{4})
	}()
	go func() {
		io.Copy(os.Stdout, f)
	}()
	time.Sleep(time.Second * 1000)
}
