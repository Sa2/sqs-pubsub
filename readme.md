# sqs-pubsub

SQSの実験のために作成しています。
そのためエラーを握り潰したりしてる部分があります。

## setup

設定に環境変数を設定する
direnvとか使って設定してください

```
export QUEUE_URL="<input your aws profile name>"
export AWS_PROFILE="<input your aws profile name>"
```

## subscriber

起動

```
$ go run ./... subscriber
```



## publisher

起動

```
$ go run ./... publisher
```
