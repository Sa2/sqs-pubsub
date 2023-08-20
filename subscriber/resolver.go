package subscriber

import (
	"Sa2/sqs-pubsub/sqsinterface"
	"Sa2/sqs-pubsub/subscriber/handler"
)

func resolver(sqsMessage sqsinterface.SQSMessage) {
	switch sqsMessage.FunctionName {
	case sqsinterface.HelloWorldFuncName:
		param, _ := sqsinterface.ResolveHelloWorld(sqsMessage.ParameterJSON)
		handler.HelloWorldHandler(param)
	}
}
