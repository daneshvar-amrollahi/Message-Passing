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
	broker.ch <- message //blocks here until server reads
}

func runClient(broker *Broker) {
	for i := 1; i <= 5; i++ {
		var message string

		message = strconv.Itoa(i)
		log.Println("CLIENT: sending " + message + " on channel")
		broker.ch <- message
		// log.Println("CLIENT: sent " + message + " on channel")

		time.Sleep(time.Second * 2)

		// log.Println("CLIENT: reading from channel")
		message = <-broker.ch
		log.Println("CLIENT: read \"" + message + "\" from channel")
	}
	close(broker.ch)
}

func runServer(broker *Broker) {
	for message := range broker.ch {
		if message[0] >= '0' && message[0] <= '9' { //message is from client
			log.Println("SERVER: sending SERVER ACK " + message + " on channel")
			broker.ch <- "SERVER ACK " + message
		}
	}
}

func main() {
	var broker Broker
	broker.ch = make(chan string, 0)
	go runClient(&broker)
	runServer(&broker)
}
