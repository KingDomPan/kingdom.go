package main

import (
	"log"
	"os"

	"golang.org/x/crypto/ssh"
)

func main() {
	// Create client config
	config := &ssh.ClientConfig{
		User: "panqd",
		Auth: []ssh.AuthMethod{
			ssh.Password("panqd"),
		},
	}
	// Connect to ssh server
	conn, err := ssh.Dial("tcp", "localhost:22", config)
	if err != nil {
		log.Fatalf("unable to connect: %s", err)
	}
	defer conn.Close()

	// Create a session
	session, err := conn.NewSession()
	if err != nil {
		log.Fatalf("unable to create session: %s", err)
	}
	defer session.Close()

	// Set up terminal modes
	modes := ssh.TerminalModes{
		ssh.ECHO:          0,     // disable echoing
		ssh.TTY_OP_ISPEED: 14400, // input speed = 14.4kbaud
		ssh.TTY_OP_OSPEED: 14400, // output speed = 14.4kbaud
	}

	// Request pseudo terminal
	if err := session.RequestPty("xterm", 80, 40, modes); err != nil {
		log.Fatalf("request for pseudo terminal failed: %s", err)
	}

	// 输入输出重定向, 实现一个建议版本的Ssh Shell客户端
	session.Stdin = os.Stdin
	session.Stdout = os.Stdout
	// Start remote shell
	if err := session.Shell(); err != nil {
		log.Fatalf("failed to start shell: %s", err)
	}

	select {}
}
