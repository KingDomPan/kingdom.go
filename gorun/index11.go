package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan struct{}, 1)

	go func() {
		select {
		case <-ch:
			fmt.Println("panqd")
		}
	}()
	go func() {
		select {
		case <-ch:
			fmt.Println("panqd2")
		}
	}()

	time.Sleep(time.Second * 1)
	close(ch)
	time.Sleep(time.Second * 2)
}
