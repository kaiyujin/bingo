package infrastructure

import (
	config2 "bingo/config"
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
)

const (
	region              string = "ap-northeast-1"
	localCustomEndpoint string = "http://localhost:8000"
)

var awsConfig aws.Config

func init() {
	var err error
	var ctx = context.Background()
	if config2.IsLocal() {
		resolver := aws.EndpointResolverWithOptionsFunc(func(service, region string, opts ...interface{}) (aws.Endpoint, error) {
			return aws.Endpoint{
				URL:               localCustomEndpoint, // custom endpoint
				HostnameImmutable: true,
			}, nil
		})
		awsConfig, err = config.LoadDefaultConfig(ctx, config.WithEndpointResolverWithOptions(resolver))
	} else {
		awsConfig, err = config.LoadDefaultConfig(ctx, config.WithRegion(region))
	}

	if err != nil {
		fmt.Printf("load aws config: %s\n", err.Error())
		panic(err)
	}
}

func AwsConfig() aws.Config {
	return awsConfig
}
