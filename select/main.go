package main

import (
	"fmt"
	"io"
	"os"
	"time"

	"github.com/creack/goselect"
	"golang.org/x/crypto/ssh/terminal"
)

type exposeFd interface {
	Fd() uintptr
}

func main() {

	buff := make([]byte, 128)
	oldState, err := terminal.MakeRaw(0)
	if err == nil {
		defer terminal.Restore(0, oldState)
	}

	rdfs := &goselect.FDSet{}
	reader := io.Reader(os.Stdin)
	for {
		select {
		default:
		}

		rdfs.Zero()
		rdfs.Set(reader.(exposeFd).Fd())
		err := goselect.Select(1, rdfs, nil, nil, 50*time.Millisecond)
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(127)
		}
		if rdfs.IsSet(reader.(exposeFd).Fd()) {
			size, err := reader.Read(buff)
			if err != nil {
				if err == io.EOF {
					fmt.Print(string(append([]byte("0"), byte(4))))
					if err != nil {
						fmt.Println(err.Error())
						os.Exit(127)
					}
					continue
				} else {
					fmt.Println(err.Error())
					os.Exit(127)
				}
			}

			if size <= 0 {
				continue
			}

			data := buff[:size]
			fmt.Print(string(append([]byte(""), data...)))
			if err != nil {
				fmt.Println(err.Error())
				os.Exit(127)
			}
		}
	}
}
