package main

import (
	"fmt"
	"log"
	"net/rpc"
	"os"
)

type Args struct {
	A, B int
}

type Quotient struct {
	Quo, Rem int
}

func main() {

	if len(os.Args) != 2 {
		fmt.Println("Usage:", os.Args[0], "server")
		os.Exit(10)
	}

	serverAddress := os.Args[1]

	client, err := rpc.Dial("tcp", serverAddress+":9000")

	if err != nil {
		log.Fatal(err)
	}

	args := Args{17, 8}

	var reply int

	err = client.Call("Arith.Multiply", args, &reply)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("*** = %d\n", reply)

	var quo Quotient
	err = client.Call("Arith.Divide", args, &quo)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%%%% = %d, %d\n", quo.Quo, quo.Rem)
}
