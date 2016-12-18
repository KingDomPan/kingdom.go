package main

import (
	"log"
	"net/http"
	"time"

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
	m.Group("/panqd", func(r martini.Router) {
		r.Get("/xx", func(res http.ResponseWriter) {
			log.Println("xxxxx")
			res.WriteHeader(400)
			res.Write([]byte("xx"))
		})
		r.Get("/:id", func() string {
			log.Println("panqd")
			return "panqd"
		})
	}, func(c martini.Context, req *http.Request) {
		start := time.Now()
		c.Next()
		log.Println(req.Method, req.URL.Path, time.Since(start))
	})
	m.RunOnAddr(":9090")
}
