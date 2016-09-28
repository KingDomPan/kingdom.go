package main

import (
	"fmt"
	"time"

	"golang.org/x/net/context"
)

func main() {
	TestDone("panqd")
}

func TestDone(name string) {
	ctx := context.Background()
	go func() {
		AsyncWork(ctx)
	}()
	time.Sleep(4 * time.Second)
}

func AsyncWork(ctx context.Context) {
	Ch := make(chan int, 1)
	go func() {
		time.Sleep(3 * time.Second)
		Ch <- 10000
	}()
	select {
	case <-ctx.Done():
		fmt.Println("Get From Done")
	case v := <-Ch:
		fmt.Println("Get From Ch", v)
	}
}
