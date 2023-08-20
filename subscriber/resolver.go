package subscriber

import (
	"Sa2/sqs-pubsub/sqsinterface"
	"Sa2/sqs-pubsub/subscriber/handler"
)

func resolver(sqsMessage sqsinterface.SQSMessage) {
	// error 伝達を実装する。handlerがエラーの場合エラーを伝播させてSQSのメッセージを削除しないようにする。
	switch sqsMessage.FunctionName {
	case sqsinterface.HelloWorldFuncName:
		param, _ := sqsinterface.ResolveHelloWorld(sqsMessage.ParameterJSON)
		handler.HelloWorldHandler(param)
	}
}
