package main

import "fmt"
import "runtime"

var ch3 chan int
var ch4 chan int

var chs = []chan int{ch3, ch4}
var numbers = []int{1, 2, 3, 4, 5}

func getChan(num int) chan int {
	fmt.Printf("chs[%d]\n", num)
	return chs[num]
}

func getNumer(num int) int {
	fmt.Printf("numbers[%d]\n", num)
	return numbers[num]
}

func main() {

	// 输出
	// chs[0]
	// numbers[2]
	// chs[1]
	// numbers[3]
	// default
	select { // 自上而下 自左而右的求表达式的值
	case getChan(0) <- getNumer(2):
		fmt.Println("1 th case is selected")
	case getChan(1) <- getNumer(3):
		fmt.Println("2 th case is selected")
	default:
		fmt.Println("default")
	}

	// nil的channel永远会被阻塞, 因此2个case都无法匹配, 输出default
	var ch5 chan int
	var ch6 chan string
	select {
	// 检测是否会被立即执行(协程不会因为此操作而阻塞, 根据缓冲或者非缓冲通道那一刻的具体情况)
	// 当发现第一个满足选择条件的case的时候, 会执行该case所包含的全部语句
	// 如果多个case满足条件, 会有一个伪随机的算法决定哪一个case将会被执行
	// 如果没有case满足并且没有default的话, 那么select就会被阻塞
	case ch5 <- 1:
		fmt.Println("3th case is selected")
	case ch6 <- "panqd":
		fmt.Println("4th case is selected")
	default:
		fmt.Println("default")
	}

	chanCap := 5
	ch7 := make(chan int, chanCap)
	for i := 0; i < chanCap; i++ {
		select {
		case ch7 <- 1:
		case ch7 <- 2:
		case ch7 <- 3:
		}
	}
	for i := 0; i < chanCap; i++ {
		fmt.Println(<-ch7)
	}
	runtime.Gosched()
}
