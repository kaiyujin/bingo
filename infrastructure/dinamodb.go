package infrastructure

import (
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"time"
)

const DynamoDefaultTimeout = 3 * time.Second

func DynamoDBClient() *dynamodb.Client {
	return dynamodb.NewFromConfig(AwsConfig())
}
