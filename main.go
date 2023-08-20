package main

import (
	"Sa2/sqs-pubsub/env"
	"Sa2/sqs-pubsub/publisher"
	"Sa2/sqs-pubsub/subscriber"
	"context"
	"fmt"
	"log"
	"os"
)

func main() {
	ctx := context.Background()
	fmt.Println("booting...")
	if len(os.Args) == 1 {
		log.Println("No arguments passed")
		os.Exit(1)
	}

	mode := os.Args[1]

	env := env.Init(ctx)

	if mode == "publisher" {
		publisher.Publisher(ctx, env)
	} else if mode == "subscriber" {
		subscriber.Subscriber(ctx, env)
	} else {
		log.Println("Invalid argument passed")
		os.Exit(1)
	}

	os.Exit(0)
}
