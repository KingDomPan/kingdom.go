package main

import (
	"fmt"
	"sync"
)

func main() {

	data := []int{1}
	base := 2
	length := len(data)
	wg := sync.WaitGroup{}

	if length <= base {
		base = length
	}

	n := length / base // 1
	m := length % base
	if m != 0 {
		n = n + 1
	}
	index := 0

	for index < n {
		start := index * base
		end := (index + 1) * base
		if index == n-1 {
			end = length
		}
		subSlice := data[start:end]
		wg.Add(1)
		go func(datas []int) {
			defer wg.Done()
			fmt.Println(datas)
		}(subSlice)
		index = index + 1
	}

	wg.Wait()
}
