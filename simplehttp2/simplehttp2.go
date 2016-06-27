package main

import (
	"fmt"
	"io/ioutil"
	"net"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		os.Exit(1)
	}
	service := os.Args[1]

	conn, err := net.Dial("tcp", service)
	checkErr(err)

	_, err = conn.Write([]byte("GET / HTTP/1.0\r\n\r\n"))
	checkErr(err)

	result, err := ioutil.ReadAll(conn)
	checkErr(err)

	fmt.Println(string(result))
}

func checkErr(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}
