package main

import (
	"log"
	"time"

	"github.com/nats-io/nats.go"
)

func main() {
	// Connect to a local NATS server
	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		log.Fatal(err)
	}
	defer nc.Close()

	// Simple publisher that publishes a message every second
	for {
		err = nc.Publish("updates", []byte("New Update"))
		if err != nil {
			log.Println("Error publishing message:", err)
		}
		time.Sleep(time.Second)
	}
}
