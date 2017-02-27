package main

import (
	"fmt"
	"io"

	"golang.org/x/crypto/ssh/terminal"

	"github.com/kr/pty"
)

type ChannelSession struct {
	reader io.ReadCloser
	writer io.WriteCloser
}

func (cs *ChannelSession) Write(p []byte) (int, error) {
	return cs.writer.Write(p)
}

func (cs *ChannelSession) Read(p []byte) (int, error) {
	return cs.reader.Read(p)
}

func main() {

	pty, tty, err := pty.Open()
	if err != nil {
		panic(err)
	}
	defer pty.Close()
	defer tty.Close()

	//pty.Write([]byte("panqd"))
	//buf := make([]byte, 1024)
	//n, _ := pty.Read(buf)
	//fmt.Println(string(buf[:n]))

	cs := &ChannelSession{
		reader: tty,
		writer: pty,
	}

	t := terminal.NewTerminal(cs, "")

	pty.Write([]byte("panqd\n"))

	fmt.Println("p")

	//buf := make([]byte, 1024)
	//n, _ := reader.Read(buf)
	//fmt.Println(string(buf[:n]))

	//buf := make([]byte, 1024)
	//n, _ := reader.Read(buf)
	//fmt.Println(string(buf[:n]))

	//line, err := reader.ReadString('\n')
	//if err != nil {
	//panic(err)
	//}
	//fmt.Println(string(line))

	line, err := t.ReadLine()
	if err != nil {
		fmt.Println(line)
	}

	//select {}

}
