package main

import (
	"fmt"

	"github.com/shirou/gopsutil/process"
)

func main() {
	p, _ := process.NewProcess(1)
	mem, _ := p.MemoryInfo()
	fmt.Println(mem.RSS)
	fmt.Println(p.Times())
}
