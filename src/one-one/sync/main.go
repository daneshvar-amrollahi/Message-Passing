package main

import (
	"log"
	"strconv"
	"time"
)

type Broker struct {
	ch chan string
}

func send(broker *Broker, message string) {
	log.Println("SERVER: sending " + message + " on channel")
	broker.ch <- message //blocks here until server reads
}

func runServer(broker *Broker) {

	for i := 1; i <= 5; i++ {
		message := strconv.Itoa(i)
		send(broker, message)
		time.Sleep(time.Millisecond * 500)
	}
	close(broker.ch)
}

func runClient(broker *Broker) {
	for message := range broker.ch {
		log.Println("SERVER: received " + message + " from channel")
	}
}

func main() {
	var broker Broker
	broker.ch = make(chan string, 0)
	go runServer(&broker)
	runClient(&broker)
}
