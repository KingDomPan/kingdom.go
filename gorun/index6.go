package main

import "fmt"

type IMessage interface {
	Print()
}
type BaseMessage struct {
	msg string
}

func (message *BaseMessage) Print() {
	fmt.Println("baseMessage:msg", message.msg)
}

type SubMessage struct {
	BaseMessage
}

func (message *SubMessage) Print() {
	fmt.Println("subMessage:msg", message.msg)
}

func interface_use(i IMessage) {
	i.Print()
}

func main() {
	baseMessage := new(BaseMessage)
	baseMessage.msg = "a"
	interface_use(baseMessage)

	SubMessage := new(SubMessage)
	SubMessage.msg = "b"
	interface_use(SubMessage)
}
