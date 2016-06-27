package main

import (
	"fmt"
	"sync"
)

func gen(nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		for _, n := range nums {
			out <- n
		}
		close(out)
	}()
	return out
}

func sq(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in { // 多个协程并发的情况下, 每个协程都只能从这个按顺序读取到一个值
			out <- n * n
		}
		close(out)
	}()
	return out
}

func merge(cs ...<-chan int) <-chan int {
	var wg sync.WaitGroup
	out := make(chan int)
	output := func(c <-chan int) {
		for n := range c {
			out <- n
		}
		wg.Done()
	}
	wg.Add(len(cs))
	for _, c := range cs {
		go output(c)
	}
	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}
func merge2(done <-chan struct{}, cs ...<-chan int) <-chan int {
	var wg sync.WaitGroup
	out := make(chan int)
	output := func(c <-chan int) {
		defer wg.Done()
		for n := range c {
			select {
			case out <- n:
			case <-done:
				return
			}
		}
		wg.Done()
	}
	wg.Add(len(cs))
	for _, c := range cs {
		go output(c)
	}
	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}

func printout(out <-chan int) {
	for n := range out {
		fmt.Println(n)
	}
}

func main() {
	//	printout(sq(sq(gen(1, 2, 3))))

	done := make(chan struct{})
	defer close(done)

	in := gen(2, 3)
	c1 := sq(in) // c1和c2会抢占从in读取到的值
	c2 := sq(in)

	out := merge2(done, c1, c2)
	fmt.Println(<-out) // 只会打印4 or 9

	/**
	for n := range merge(c1, c2) {
		fmt.Println(n)
	}
	**/

	/** 这样子的话会出现主线程无限制等待channel输入数据, 但实际上已经没有哪个协程为channel输入数据, 造成死锁
	fmt.Println("test read from channel")
	testChannel := make(chan int, 2)
	testChannel <- 1
	testChannel <- 2
	for n := range testChannel {
		fmt.Println(n)
	}
	for n := range testChannel {
		fmt.Println(n)
	}
	**/

	/**
	fmt.Println("test read from channel which read from goroutine")
	testChannel := make(chan int, 2)
	testChannel <- 1
	testChannel <- 2
	go func() {
		for n := range testChannel {
			fmt.Println(n)
		}
	}()

	go func() {
		for n := range testChannel {
			fmt.Println(n)
		}
	}()
	**/

}
