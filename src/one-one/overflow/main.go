package main

import (
	"log"
	"strconv"
	"time"
)

const BUFFER_SIZE = 3

type Broker struct {
	ch chan string
	sz int
}

func send(broker *Broker, message string) {
	log.Println("CLIENT: sending message on channel  " + strconv.Itoa(broker.sz))
	if broker.sz < BUFFER_SIZE {
		broker.sz += 1
		broker.ch <- message
		log.Println("CLIENT: sent message on channel     " + strconv.Itoa(broker.sz))

	} else {
		log.Println("BUFFER OVERFLOW! CANNOT SEND NEW MESSAGE FOR NOW")
	}
}

func recv(broker *Broker) string {
	if broker.sz > 0 {
		broker.sz -= 1
		return <-broker.ch
	} else {
		log.Fatal("BUFFER UNDERFLOW!")
		return ""
	}
}

func main() {
	var broker Broker
	broker.ch = make(chan string, BUFFER_SIZE)

	go func() {
		for i := 1; i <= 10; i++ {
			send(&broker, "message")
			time.Sleep(time.Millisecond * 40)
		}
	}()

	go func() {
		for {
			time.Sleep(time.Millisecond * 80)
			message := recv(&broker)
			if len(message) > 0 {
				log.Println("SERVER: received message from channel")
			}
		}
	}()

	time.Sleep(time.Second * 100)
}
