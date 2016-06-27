package bufchannel

import "fmt"
import "time"

func main() {
	ch := make(chan int, 10)
	for i := 0; i < 10; i++ {
		go func(index int) {
			ch <- index
		}(i)
	}
	go func() {
		for j := range ch { // 会阻塞, 存在bug, 在主线中会无限制等在channel的数据, 如果在非协程环境下
			fmt.Printf("Index: %d\n", j)
		}
	}()
	time.Sleep(time.Second * 2) // 主线程必须等待子线程结束, 否则退出程序
}
