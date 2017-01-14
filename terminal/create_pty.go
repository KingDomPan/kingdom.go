package main

import (
	"io"
	"os"
	"os/exec"

	"github.com/kr/pty"
)

// 1.根据ptmx创建1个tty
// 2.将命令的输入输出重定向到tty
// 3.往ptmx写就是让tty的stdin可以读
// 4.往tty写就是让ptmx可以读

func main() {
	c := exec.Command("grep", "--color=auto", "bar")
	f, err := pty.Start(c) // 返回pty, cmd的输入输出指向tts
	if err != nil {
		panic(err)
	}

	go func() {
		f.Write([]byte("foo\n"))
		f.Write([]byte("bar\n"))
		f.Write([]byte("baz\n"))
		f.Write([]byte{4}) // EOT end of transmission
	}()
	io.Copy(os.Stdout, f)
}
