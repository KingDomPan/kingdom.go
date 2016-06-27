package selecttimeout

import "time"
import "fmt"

// 利用select来实现channel的超时机制, 主要是因为select只要一个满足条件就会向下执行, 不会考虑其他case的情况
func main() {

	timeout := make(chan bool, 1)
	ch := make(chan int, 10)

	go func() {
		time.Sleep(1e9)
		timeout <- true
	}()

	select {
	case <-timeout:
		fmt.Println("Read From Timeout")
	case <-ch:
		fmt.Println("Read From ch")
	}

	fmt.Println("Main Thread Here")

}
