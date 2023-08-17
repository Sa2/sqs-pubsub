package main

import (
	"Sa2/sqs-pubsub/publisher"
	"Sa2/sqs-pubsub/subscriber"
	"fmt"
	"log"
	"os"
)

func main() {
	fmt.Println("booting...")
	if len(os.Args) == 1 {
		log.Println("No arguments passed")
		os.Exit(1)
	}

	mode := os.Args[1]

	if mode == "publisher" {
		publisher.RunPublisher()
	} else if mode == "subscriber" {
		subscriber.Subscriber()
	} else {
		log.Println("Invalid argument passed")
		os.Exit(1)
	}

	os.Exit(0)
}
