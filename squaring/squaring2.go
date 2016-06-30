// 管道构建指南
// 1. 状态会在所有发送操作做完后, 关闭它们的流出channel
// 2. 状态会持续接收从流入channel输入的值, 直到channel关闭或者其发送者被释放
package main

import (
	"fmt"
	"sync"
)

func gen(nums ...int) <-chan int { // 当知道列表长度的时候可以使用缓冲channel优化
	out := make(chan int, len(nums))
	for _, n := range nums {
		out <- n
	}
	close(out)
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

func merge(done <-chan struct{}, cs ...<-chan int) <-chan int {
	var wg sync.WaitGroup
	out := make(chan int)
	output := func(c <-chan int) { // 在每一个channel读取数据的时候判断是否收到done的值
		defer wg.Done()
		for n := range c {
			select {
			case out <- n:
			case <-done: // 如果收到done的数据就代表要显示退出
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

	done := make(chan struct{})
	defer close(done)

	in := gen(2, 3)
	c1 := sq(in) // c1和c2会抢占从in读取到的值
	c2 := sq(in)

	out := merge(done, c1, c2)
	fmt.Println(<-out) // done中没有值, 所以会打印4 or 9

}
