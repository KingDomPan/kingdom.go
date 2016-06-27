package singlechannel

// 单向channel, 限制某个操作对其只能读或者写

func main() {

	var ch1 chan int       // 正常的channel, 可读可写
	var ch2 chan<- float64 // 单向channel, 只用于写float64数据
	var ch3 <-chan int     // 单向channel, 只用于读取int数据

	ch4 := make(chan int)  // 正常的channel
	ch5 := <-chan int(ch4) // 转换为单向的读channel
	ch6 := chan<- int(ch4) // 转换为单向的写channel
}
