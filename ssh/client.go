package main

import (
	"bufio"
	"fmt"
	"io"

	"golang.org/x/crypto/ssh"
)

func main() {
	config := &ssh.ClientConfig{
		User: "panqd",
		Auth: []ssh.AuthMethod{
			ssh.Password("panqd"),
		},
	}
	client, err := ssh.Dial("tcp", "127.0.0.1:22", config)
	if err != nil {
		panic("Failed to dial: " + err.Error())
	}

	// Each ClientConn can support multiple interactive sessions,
	// represented by a Session.
	session, err := client.NewSession()
	if err != nil {
		panic("Failed to create session: " + err.Error())
	}
	defer session.Close()

	// Once a Session is created, you can execute a single command on
	// the remote side using the Run method.
	r, w := io.Pipe()
	session.Stdout = w
	session.Stderr = w
	defer r.Close()
	defer w.Close()
	if err := session.Start("ping -c 1 127.0.0.1"); err != nil {
		panic("Failed to run: " + err.Error())
	}

	ch := make(chan struct{}, 1)
	go func() {
		s := bufio.NewScanner(r)
		for s.Scan() {
			fmt.Println(s.Text())
		}
		ch <- struct{}{}
	}()
	if err := session.Wait(); err != nil {
		panic("Failed to run: " + err.Error())
	}
	w.Close()
	<-ch
}
