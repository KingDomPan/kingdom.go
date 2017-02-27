package main

import "fmt"

type Q interface {
	Panqd() string
}

type S struct {
}

func (s *S) Panqd() string {
	return "Panqd"
}

func main() {
	q := S{}
	fmt.Println(q.Panqd())
}
