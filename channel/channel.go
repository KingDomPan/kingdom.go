package main

import (
	"fmt"
	"time"
)

func main() {
	messages := make(chan string)
	go func() { messages <- "ping" }()

	time.Sleep(time.Duration(2) * time.Second)

	msg := <-messages
	fmt.Println(msg)

	buf := make(chan string, 2)

	buf <- "buffer1"
	buf <- "channel"

	fmt.Println(<-buf)
	fmt.Println(<-buf)
}
