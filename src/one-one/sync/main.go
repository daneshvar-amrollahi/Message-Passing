package main

import (
	"fmt"
	"strconv"
	"time"
)

type Broker struct {
	ch chan string
}

func runClient(broker Broker) {
	fmt.Println("CLIENT: sending messages 1..5 on channel")
	for i := 1; i <= 5; i++ {
		message := strconv.Itoa(i)
		broker.ch <- message //blocks here until server reads
	}
	close(broker.ch)
}

func runServer(broker Broker) {
	for message := range broker.ch {
		fmt.Println("SERVER: received " + message + " from channel")
		time.Sleep(time.Second)
	}
}

func main() {
	var broker Broker
	broker.ch = make(chan string, 0)
	go runClient(broker)
	runServer(broker)
}
