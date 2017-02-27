package main

import "os/exec"

func RunASession() {
	cmd := exec.Command("bash", "-c", "ssh panqingdao@127.0.0.1 -p 2022")
	if err := cmd.Start(); err != nil {
		panic(err)
	}
	if err := cmd.Wait(); err != nil {
		panic(err)
	}
}

func main() {

	for i := 0; i < 1000; i++ {
		go RunASession()
	}

	select {}
}
