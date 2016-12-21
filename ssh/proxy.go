package main

import (
	"fmt"
	"log"
	"net/http"

	"golang.org/x/crypto/ssh"
)

/**
1. 创建ssh连接
2. 在ssh的连接上监听http连接(要求远程服务器开启http server)
*/
func main() {
	config := &ssh.ClientConfig{
		User: "panqd",
		Auth: []ssh.AuthMethod{
			ssh.Password("panqd"),
		},
	}
	// Dial your ssh server.
	conn, err := ssh.Dial("tcp", "localhost:22", config)
	if err != nil {
		log.Fatalf("unable to connect: %s", err)
	}
	defer conn.Close()

	// Request the remote side to open port 9090 on all interfaces.
	l, err := conn.Listen("tcp", "0.0.0.0:9090")
	if err != nil {
		log.Fatalf("unable to register tcp forward: %v", err)
	}
	defer l.Close()

	// Serve HTTP with your SSH server acting as a reverse proxy.
	http.Serve(l, http.HandlerFunc(func(resp http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(resp, "Hello world!\n")
	}))
}
