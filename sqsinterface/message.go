package sqsinterface

type SQSMessage struct {
	FunctionName  string `json:"functionName"`
	ParameterJSON string `json:"parameterJSON"`
}
