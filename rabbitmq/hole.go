package hole

import (
	"fmt"
	"log"

	"github.com/michaelklishin/rabbit-hole"
)

func main() {
	rmpc, _ := rabbithole.NewClient("http://192.168.99.100:15672", "admin", "panqd")
	x, err := rmpc.GetVhost("/")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(x.Name)
}
