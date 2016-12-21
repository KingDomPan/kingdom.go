package main

import (
	"fmt"
	"io/ioutil"
	"net"

	"golang.org/x/crypto/ssh"
	"golang.org/x/crypto/ssh/terminal"
)

func main() {

	config := &ssh.ServerConfig{
		PasswordCallback: func(c ssh.ConnMetadata, pass []byte) (*ssh.Permissions, error) {
			if c.User() == "panqd" && string(pass) == "panqd" {
				fmt.Printf("%+v", c)
				fmt.Println()
				fmt.Printf("%s", string(c.ClientVersion()))
				fmt.Printf("%s", string(c.ServerVersion()))
				return nil, nil
			}
			return nil, fmt.Errorf("password rejected for %q", c.User())
		},
	}

	privateBytes, err := ioutil.ReadFile("/Users/panqd/.ssh/id_rsa")
	if err != nil {
		panic("Failed to load private key")
	}

	private, err := ssh.ParsePrivateKey(privateBytes)
	if err != nil {
		panic("Failed to parse private key")
	}

	config.AddHostKey(private)

	listener, err := net.Listen("tcp", "0.0.0.0:2022")
	if err != nil {
		panic("failed to listen for connection")
	}
	for {
		nConn, err := listener.Accept()
		if err != nil {
			panic("failed to accept incoming connection")
		}
		fmt.Println("1. new tcp conn comes")
		go func(net.Conn) {
			_, chans, reqs, err := ssh.NewServerConn(nConn, config)
			if err != nil {
				panic("failed to handshake")
			}

			go ssh.DiscardRequests(reqs)

			// Service the incoming Channel channel.
			for newChannel := range chans {
				// Channels have a type, depending on the application level
				// protocol intended. In the case of a shell, the type is
				// "session" and ServerShell may be used to present a simple
				// terminal interface.
				fmt.Println("2. newChannel comes, channel type is", newChannel.ChannelType())
				if newChannel.ChannelType() != "session" {
					newChannel.Reject(ssh.UnknownChannelType, "unknown channel type")
					continue
				}
				channel, requests, err := newChannel.Accept()
				if err != nil {
					panic("could not accept channel.")
				}
				fmt.Println("3. channel from newChannel.Accept()")

				// Sessions have out-of-band requests such as "shell",
				// "pty-req" and "env".  Here we handle only the
				// "shell" request.
				go func(in <-chan *ssh.Request) {
					for req := range in {
						fmt.Println("4. requests from channel")
						ok := false
						switch req.Type {
						case "shell":
							ok = true
							if len(req.Payload) > 0 {
								ok = false
							}
						}
						req.Reply(ok, nil)
					}
				}(requests)

				term := terminal.NewTerminal(channel, "> ")

				go func() {
					defer channel.Close()
					for {
						line, err := term.ReadLine()
						if err != nil {
							break
						}
						fmt.Println(line)
					}
				}()
			}
		}(nConn)
	}
}
