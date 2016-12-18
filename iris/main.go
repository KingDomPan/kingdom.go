package main

import (
	"bufio"
	"fmt"
	"time"

	"github.com/kataras/iris"
)

func main() {
	iris.Any("/stream", func(ctx *iris.Context) {
		ctx.StreamWriter(stream)
	})
	iris.Listen(":7070")
}

func stream(w *bufio.Writer) {
	for i := 0; i < 10; i++ {
		if i == 9 {
			fmt.Fprintf(w, "0\r\n")
			fmt.Fprintf(w, "\r\n")
		} else {
			fmt.Fprintf(w, "5\r\n")
			fmt.Fprintf(w, "12345\r\n")
		}
		if err := w.Flush(); err != nil {
			return
		}
		time.Sleep(time.Second)
	}
}
