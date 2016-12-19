package main

import (
	"fmt"
	"io"
	"time"
)

func main() {
	r, w := io.Pipe()

	go func() {
		for i := 0; i < 10; i++ {
			w.Write([]byte("panqd"))
			time.Sleep(time.Second * 1)
		}
	}()

	go func() {
		buf := make([]byte, 32)
		for {
			r.Read(buf)
			fmt.Println(string(buf))
		}
	}()

	time.Sleep(time.Second * 1000)
}
