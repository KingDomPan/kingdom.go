package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()

	d, err := time.ParseDuration("85s")
	if err == nil {
		if d < time.Minute {
			d = time.Minute
		}
	}
	fmt.Println(now.Add(d))
}
