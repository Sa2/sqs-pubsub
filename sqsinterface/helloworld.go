package sqsinterface

import "encoding/json"

const HelloWorldFuncName = "helloworld"

type HelloWorldMessage struct {
	Message string `json:"message"`
}

func ResolveHelloWorld(parameterJSON string) (HelloWorldMessage, error) {
	var helloWorldMessage HelloWorldMessage
	err := json.Unmarshal([]byte(parameterJSON), &helloWorldMessage)
	if err != nil {
		return HelloWorldMessage{}, err
	}
	return helloWorldMessage, nil
}
