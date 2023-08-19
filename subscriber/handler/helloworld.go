package handler

import (
	"Sa2/sqs-pubsub/sqsinterface"
	"log"
)

func HelloWorldHandler(param sqsinterface.HelloWorldMessage) {
	log.Println("Hello World!: ", param.Message)
}
