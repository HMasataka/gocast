package main

import (
	"fmt"
	"time"

	"github.com/HMasataka/gocast"
)

type client struct{}

func (c client) Send(message []byte) {
	fmt.Println(string(message))
}

func (c client) Close() {}

func main() {
	c := client{}

	hub := gocast.NewHub()
	go hub.Run()

	hub.Register(c)

	hub.Broadcast([]byte("Hello, World!"))

	time.Sleep(time.Second)

	hub.Unregister(c)
}
