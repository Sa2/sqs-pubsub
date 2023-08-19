package env

import (
	"context"
	"log"

	"github.com/sethvargo/go-envconfig"
)

type Env struct {
	QueueURL string `env:"QUEUE_URL"`
}

func Init(ctx context.Context) Env {
	var env Env
	err := envconfig.Process(ctx, &env)
	if err != nil {
		log.Fatal(err.Error())
	}
	return env
}
