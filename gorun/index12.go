package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sync"
	"time"

	"golang.org/x/crypto/ssh/terminal"
)

type cmdTerminal struct {
	reader    io.ReadCloser
	bufReader *bufio.Reader
	writer    io.WriteCloser
	sync.Mutex
}

func (ct *cmdTerminal) Read(p []byte) (int, error) {
	return ct.bufReader.Read(p)
}

func (ct *cmdTerminal) Write(p []byte) (int, error) {
	return ct.writer.Write(p)
}

func main() {

	r, w, _ := os.Pipe()
	ct := &cmdTerminal{
		reader:    r,
		writer:    w,
		bufReader: bufio.NewReader(r),
	}

	t := terminal.NewTerminal(ct, "")
	go func() {
		time.Sleep(1 * time.Second)
		t.Write([]byte("p"))
		t.Write([]byte("s"))
		t.Write([]byte("\n"))
	}()

	for {
		cmd, err := t.ReadLine()
		if err != nil && err == io.EOF {
			continue
		}
		fmt.Println(string(cmd))
	}
}
