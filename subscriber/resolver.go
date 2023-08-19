package subscriber

import (
	"Sa2/sqs-pubsub/sqsinterface"
	"Sa2/sqs-pubsub/subscriber/handler"
)

func resolver(funcName string, parameterJSON string) {
	switch funcName {
	case sqsinterface.HelloWorldFuncName:
		param, _ := sqsinterface.ResolveHelloWorld(parameterJSON)
		handler.HelloWorldHandler(param)
	}
}
