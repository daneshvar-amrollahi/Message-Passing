package main

import (
	"fmt"
	"strconv"
	"time"
)

type Broker struct {
	ch chan string
}

var broker Broker

func runClient() {
	for i := 1; i <= 5; i++ {
		message := strconv.Itoa(i)
		fmt.Println("CLIENT: sent " + message + " to channel")
		broker.ch <- message //blocks here until server reads from the channel (sync)
		time.Sleep(time.Second)
	}
	close(broker.ch)
}

func runServer() {
	for message := range broker.ch {
		fmt.Println("SERVER: received " + message + " from channel")
		time.Sleep(time.Second)
	}
}

func main() {
	broker.ch = make(chan string, 0)
	go runServer() //server is waiting in the background to read messages
	runClient()
}
