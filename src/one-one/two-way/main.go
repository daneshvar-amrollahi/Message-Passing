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
		var message string
		message = strconv.Itoa(i)
		log.Println("SERVER: sending " + message + " on channel")
		broker.ch <- message
		time.Sleep(time.Second * 2)
		message = <-broker.ch
		log.Println("SERVER: read \"" + message + "\" from channel")
	}
	close(broker.ch)
}

func runClient(broker *Broker) {
	for message := range broker.ch {
		if message[0] >= '0' && message[0] <= '9' { //message is from server
			log.Println("CLIENT: sending SERVER ACK " + message + " on channel")
			broker.ch <- "CLIENT ACK " + message
		}
	}
}

func main() {
	var broker Broker
	broker.ch = make(chan string, 0)
	go runServer(&broker)
	runClient(&broker)
}
