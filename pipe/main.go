package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	r, w, _ := os.Pipe()

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
