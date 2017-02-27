package main

import (
	"fmt"
	"os/exec"
)

func main() {

	out, _ := exec.Command("bash", "-c", "ps aux | grep '2201.*node' | grep -v grep | awk '{print $3,$6}'").Output()
	fmt.Println(string(out))

}
