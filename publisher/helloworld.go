package publisher

import (
	"Sa2/sqs-pubsub/sqsinterface"
	"encoding/json"
	"fmt"
	"time"
)

func SetHelloWorldMessage() (sqsinterface.SQSMessage, error) {

	now := time.Now()
	param := sqsinterface.HelloWorldMessage{
		Message: fmt.Sprintf("from publisher %v", now.Format("2006-01-02 15:04:05")),
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
