package main

import (
	"fmt"
	"time"
)

func main() {

	runners := map[uint64]string{
		1: "panqd",
		2: "kingdom",
		3: "kingdom",
	}

	tasks := []uint64{2, 3}

	for key, _ := range runners {

		fmt.Println(runners)
		var x uint64 = 0
		for _, t := range tasks {
			if t == key {
				x = t
			}
		}
		if x == 0 {
			delete(runners, key)
		}
	}

	fmt.Println(runners)

	//now, _ := time.ParseDuration("2s")
	//timer := time.NewTimer(now)
	//go func() {
	//select {
	//case value := <-timer.C:
	//fmt.Println("______")
	//fmt.Println(value)
	//}
	//}()

	fmt.Println(time.Now().Unix())
	fmt.Println(time.Now().Weekday() == time.Monday)

	fmt.Println(time.Now().Add(time.Hour * 24 * 7))
}
