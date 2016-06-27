package passchannel

import "fmt"

// 管道类型数据
type PipeData struct {
	value   int
	handler func(int) int
	next    chan int
}

// 处理函数
func hanlde(queue chan *PipeData) {
	for data := range queue {
		data.next <- data.handler(data.value)
	}
}

func main() {
	data := &PipeData{10, func(value int) int {
		return value * 10
	}, make(chan int, 1)}

	queue := make(chan *PipeData, 1)

	queue <- data // 将数据放进第一个管道

	go func() {
		hanlde(queue) // 处理数据
	}()

	fmt.Println(<-data.next) // 从管道中拿出数据
}

// func pipeHandle() 定义第二个管道处理函数
