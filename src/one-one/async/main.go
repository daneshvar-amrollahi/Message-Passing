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
	log.Println("CLIENT: sending " + message + " on channel")
	broker.ch <- message //does not block here unless buffer is full (async)
}

func runClient(broker *Broker) {

	for i := 1; i <= 60; i++ {
		message := strconv.Itoa(i)
		send(broker, message)
		time.Sleep(time.Nanosecond)
	}
	close(broker.ch)
}

func runServer(broker *Broker) {
	for message := range broker.ch {
		log.Println("SERVER: received " + message + " from channel")
	}
}

func main() {
	var broker Broker
	broker.ch = make(chan string, 1024) //increased buffer size from 0 to 1024 to avoid blocking
	go runClient(&broker)
	runServer(&broker)
}
