package main

import (
	"log"

	"golang.org/x/crypto/ssh"
)

func main() {
	user := "panqd"
	NumberOfPrompts := 3

	// Normally this would be a callback that prompts the user to answer the
	// provided questions
	Cb := func(user, instruction string, questions []string, echos []bool) (answers []string, err error) {
		return []string{"panqd", "panqd"}, nil
	}

	config := &ssh.ClientConfig{
		User: user,
		Auth: []ssh.AuthMethod{
			ssh.RetryableAuthMethod(ssh.KeyboardInteractiveChallenge(Cb), NumberOfPrompts),
		},
	}

	client, err := ssh.Dial("tcp", "0.0.0.0:22", config)
	if err != nil {
		log.Fatalf("unable to connect: %v", err)
	}
	defer client.Close()
}
