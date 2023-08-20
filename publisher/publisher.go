package publisher

import (
	"Sa2/sqs-pubsub/awsconfig"
	"Sa2/sqs-pubsub/env"
	"context"
	"encoding/json"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/sqs"

	"github.com/google/uuid"
)

func Publisher(ctx context.Context, env env.Env) {
	log.Println("Running publisher...")

	// ここでsqsをpublishする

	awscfg, err := awsconfig.GetAWSConfig(ctx)
	if err != nil {
		log.Fatal(err)
	}

	client := sqs.NewFromConfig(awscfg)

	hellowWorldMessage, err := SetHelloWorldMessage()
	if err != nil {
		log.Fatal(err)
	}

	messageBody, err := json.Marshal(hellowWorldMessage)
	if err != nil {
		log.Fatal(err)
	}
	messageBodyStr := string(messageBody)

	// メッセージがどういう処理で使うものなのかは、MessageAttributesを使ってみても良いかもしれない
	_, err = client.SendMessage(ctx, &sqs.SendMessageInput{
		QueueUrl:               aws.String(env.QueueURL),
		MessageBody:            aws.String(messageBodyStr),
		MessageDeduplicationId: aws.String(uuid.New().String()),
		MessageGroupId:         aws.String(uuid.New().String()),
	})
	if err != nil {
		log.Fatal(err)
	}
}
