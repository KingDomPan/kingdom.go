package main

import (
	"log"

	"github.com/go-martini/martini"
)

func main() {
	m := martini.Classic()
	m.Get("/", func() string {
		return "hello panqd"
	})

	m.Get("/service", func(log *log.Logger, context martini.Context) {
		log.Printf("Get Service")
	})

	m.RunOnAddr(":9090")

}
