package main

import (
	"os"
	"path/filepath"
)

func main() {

	root := "/Users/panqd/ssh-relay/data/record"
	name := "panqingdao"
	path := filepath.Join(root, name)

	resultCh := ListSubDirectory(ListSubDirectory(GenPath(path)))
	for f := range resultCh {
		println(f)
	}

}

func GenPath(paths ...string) <-chan string {
	out := make(chan string)

	go func() {
		defer close(out)
		for _, path := range paths {
			out <- path
		}
	}()

	return out
}

func ListSubDirectory(in <-chan string) <-chan string {

	out := make(chan string)

	go func() {

		defer close(out)

		for dir := range in {

			file, err := os.Open(dir)
			if err != nil {
				return
			}
			defer file.Close()

			dirs, err := file.Readdirnames(-1)
			if err != nil {
				return
			}

			for _, subdir := range dirs {
				out <- filepath.Join(dir, subdir)
			}
		}
	}()

	return out
}
