package main

import (
	"fmt"
	"time"

	"github.com/HMasataka/gocast"
)

func NewClient() gocast.Client {
	return client{}
}

type client struct{}

func (c client) Write(message []byte) (int, error) {
	fmt.Println(string(message))
	return len(message), nil
}

func (c client) Close() error {
	return nil
}

func (c client) Error(err error) {
	fmt.Println(err)
}

func main() {
	c := client{}

	hub := gocast.NewHub()
	go hub.Run()

	hub.Register(c)

	hub.Broadcast([]byte("Hello, World!"))

	time.Sleep(time.Second)

	hub.Unregister(c)
}
