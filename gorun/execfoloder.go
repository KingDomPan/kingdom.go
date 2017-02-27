package main

import "github.com/kardianos/osext"

func main() {

	path, err := osext.Executable()
	if err != nil {
		panic(err)
	}
	println(path)

}
