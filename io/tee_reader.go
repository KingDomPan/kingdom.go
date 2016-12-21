package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	r := strings.NewReader("some io.Reader stream to be read\n")
	tee := io.TeeReader(r, os.Stdout)

	buf := make([]byte, 32)
	b, err := tee.Read(buf)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d", b)
}
