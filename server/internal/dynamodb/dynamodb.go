package dynamodb

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	sdkdynamodb "github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

type DDBClient struct {
	db        *sdkdynamodb.Client
	tableName string
}

func NewAppClient(cfg aws.Config, tableName string) *DDBClient {
	return &DDBClient{
		db:        sdkdynamodb.NewFromConfig(cfg),
		tableName: tableName,
	}
}
