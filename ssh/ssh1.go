package main

import (
	"bufio"
	"os"

	"golang.org/x/crypto/ssh"
)

func main() {
	SSH("panqd", "panqd", "PanKingDom.local:22")
}

func SSH(user, password, ip_port string) {
	PassWd := []ssh.AuthMethod{ssh.Password(password)}
	Conf := ssh.ClientConfig{User: user, Auth: PassWd}
	Client, _ := ssh.Dial("tcp", ip_port, &Conf)
	defer Client.Close()
	a := bufio.NewReader(os.Stdin)
	for {
		b, _, z := a.ReadLine()
		if z != nil {
			return
		}
		command := string(b)
		if session, err := Client.NewSession(); err == nil {
			defer session.Close()
			session.Stdout = os.Stdout
			session.Stderr = os.Stderr
			session.Run(command)
		}
	}
}
