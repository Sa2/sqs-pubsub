package sqsinterface

import (
	"context"
	"encoding/json"
)

type SQSMessage struct {
	FunctionName  string `json:"functionName"`
	ParameterJSON string `json:"parameterJSON"`
}

func GetSQSMessageFromSQSResBody(ctx context.Context, messageBody *string) (SQSMessage, error) {

	sqsMessage := SQSMessage{}
	err := json.Unmarshal([]byte(*messageBody), &sqsMessage)
	if err != nil {
		return SQSMessage{}, err
	}

	return sqsMessage, nil
}
