package awsconfig

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
)

func GetAWSConfig(ctx context.Context) (aws.Config, error) {
	//config.LoadDefaultConfig(ctx, config.WithSharedConfigProfile("PROFILE_NAME"))cfg, err := config.LoadDefaultConfig(ctx)
	cfg, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		return aws.Config{}, err
	}

	return cfg, nil
}
