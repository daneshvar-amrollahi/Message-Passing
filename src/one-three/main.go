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
	log.Println("CLIENT: sending " + message)
	broker.ch <- message
}

func main() {
	var broker Broker
	broker.ch = make(chan string, 1024)

	go runServer(&broker, 1)
	go runServer(&broker, 2)
	go runServer(&broker, 3)

	for i := 1; i <= 100; i++ {
		send(&broker, strconv.Itoa(i))
		if i%4 == 0 {
			time.Sleep(time.Microsecond)
		}
	}

	time.Sleep(time.Second * 2)
	close(broker.ch)
}

func runServer(broker *Broker, id int) {
	for message := range broker.ch {
		if len(message) > 0 {
			log.Println("SERVER " + strconv.Itoa(id) + " received " + message)
			time.Sleep(time.Microsecond)
		}
	}
}
