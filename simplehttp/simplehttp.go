package main

import (
	"bytes"
	"fmt"
	"io"
	"net"
	"os"
)

func checkErr(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}

func readFully(conn net.Conn) ([]byte, error) {
	defer conn.Close()
	result := bytes.NewBuffer(nil)
	var buf [512]byte
	for {
		n, err := conn.Read(buf[0:])
		result.Write(buf[0:n])
		if err != nil {
			if err == io.EOF {
				break
			}
			return nil, err
		}
	}
	return result.Bytes(), nil
}

func main() {

	if len(os.Args) != 2 {
		fmt.Fprint(os.Stderr, "Usage: %s host:port", os.Args[0])
		os.Exit(1)
	}

	service := os.Args[1]

	conn, err := net.Dial("tcp", service)
	checkErr(err)

	_, err = conn.Write([]byte("GET / HTTP/1.0\r\n\r\n"))
	checkErr(err)

	result, err := readFully(conn)
	checkErr(err)

	fmt.Println(string(result))

	os.Exit(0)
}
