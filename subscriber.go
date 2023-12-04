package main

import (
	"log"

	"github.com/nats-io/nats.go"
)

func main() {
	// Connect to a local NATS server
	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		log.Fatal(err)
	}
	defer nc.Close()

	// Subscribe to updates subject
	_, err = nc.Subscribe("updates", func(msg *nats.Msg) {
		log.Printf("Received message: %s", msg.Data)
	})
	if err != nil {
		log.Fatal(err)
	}

	// Keep the subscriber running
	select {}
}
