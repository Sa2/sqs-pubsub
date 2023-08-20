package subscriber

import (
	"Sa2/sqs-pubsub/awsconfig"
	"Sa2/sqs-pubsub/env"
	"Sa2/sqs-pubsub/sqsinterface"
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"runtime"
	"sync"
	"syscall"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
)

// https://docs.aws.amazon.com/code-library/latest/ug/go_2_sqs_code_examples.html
// TODO: 上記を参考にinterfaceを定義して実装していくように改変する

func run(ctx context.Context, goroutineID int, env env.Env) {
	awscfg, err := awsconfig.GetAWSConfig(ctx)
	if err != nil {
		log.Fatal(err)
		return
	}

	client := sqs.NewFromConfig(awscfg)

	sqsReceiveParam := &sqs.ReceiveMessageInput{
		QueueUrl:            aws.String(env.QueueURL),
		MaxNumberOfMessages: 1,
	}
	msgResult, err := client.ReceiveMessage(ctx, sqsReceiveParam)
	if err != nil {
		log.Fatal(err)
		return
	}

	if msgResult.Messages != nil {
		sqsMessage, err := sqsinterface.GetSQSMessageFromSQSResBody(ctx, msgResult.Messages[0].Body)
		if err != nil {
			log.Fatal(err)
			return
		}

		// 実処理へ
		resolver(sqsMessage)
		// 処理が完了時はここでメッセージを削除する

		_, err = deleteMessage(ctx, client, &sqs.DeleteMessageInput{
			QueueUrl:      aws.String(env.QueueURL),
			ReceiptHandle: msgResult.Messages[0].ReceiptHandle,
		})
		if err != nil {
			log.Fatal(err)
		}

	}
}

func Subscriber(ctx context.Context, env env.Env) {
	log.Println("Running subscriber...")

	quitChannel := make(chan os.Signal, 1) // need buffered channel
	signal.Notify(quitChannel, syscall.SIGINT, syscall.SIGTERM)

	shutdownChannel := make(chan struct{})
	waitGroup := &sync.WaitGroup{}

	// goroutineの並列度はここでコントロールできる
	for i := 0; i < 1; i++ {
		waitGroup.Add(1)

		go func(shutdownChannel chan struct{}, wg *sync.WaitGroup, i int) {
			log.Println("Starting goroutine: ", i)
			defer wg.Done()
			// ここでsqsをsubscribeする
			//run(ctx, i, env)

			// shutdown signal を受け取るまでループし続ける
			for {
				run(ctx, i, env)
				select {
				case <-shutdownChannel:
					log.Println("Shutdown goroutine: ", i)
					return
				default:
					runtime.Gosched()
				}
			}
		}(shutdownChannel, waitGroup, i)
	}

	<-quitChannel // received SIGINT or SIGTERM
	close(shutdownChannel)

	log.Println("Quit signal received, gracefully shutdown goroutines...")

	waitGroup.Wait() // wait for all goroutines

	/* you can do extra work here, goroutines are all stopped now */

	log.Println("Done!")
}

// sqsDeleteMessageAPI AWS SDK のサンプルコードを参考に。こういうinterfaceを引数の型にするの良いよね
type sqsDeleteMessageAPI interface {
	DeleteMessage(ctx context.Context,
		params *sqs.DeleteMessageInput,
		optFns ...func(*sqs.Options)) (*sqs.DeleteMessageOutput, error)
}

func deleteMessage(ctx context.Context, client sqsDeleteMessageAPI, dMInput *sqs.DeleteMessageInput) (*sqs.DeleteMessageOutput, error) {
	fmt.Println("Deleting message...", *dMInput.ReceiptHandle)
	return client.DeleteMessage(ctx, dMInput)
}
