package publisher

import (
	"Sa2/sqs-pubsub/sqsinterface"
	"encoding/json"
)

func SetHelloWorldMessage() (sqsinterface.SQSMessage, error) {
	param := sqsinterface.HelloWorldMessage{
		Message: "from publisher",
	}

	paramJSON, err := json.Marshal(param)
	if err != nil {
		return sqsinterface.SQSMessage{}, err
	}

	return sqsinterface.SQSMessage{
		FunctionName:  sqsinterface.HelloWorldFuncName,
		ParameterJSON: string(paramJSON),
	}, nil
}
