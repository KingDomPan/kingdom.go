package main

import (
	"fmt"
	"time"

	"github.com/olivere/elastic"
)

func main() {
	GetClient()
}
func GetClient() (client *elastic.Client, err error) {
	urls := []string{
		"http://127.0.0.1:9200",
	}
	options := []elastic.ClientOptionFunc{
		elastic.SetURL(urls...),
		elastic.SetSnifferTimeoutStartup(time.Second * 300),
	}

	client, err = elastic.NewClient(options...)

	if err != nil {

		fmt.Println(err)
	}

	return
}
