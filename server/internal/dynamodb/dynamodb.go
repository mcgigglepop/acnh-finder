package dynamodb

import (
	"context"
	"errors"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	sdkdynamodb "github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/mcgigglepop/acnh-finder/server/internal/models"
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

func (c *DDBClient) GetUserProfile(ctx context.Context, userSub string) (*models.User, error) {
	input := &sdkdynamodb.GetItemInput{
		TableName: aws.String(c.tableName),
		Key: map[string]types.AttributeValue{
			"user_id": &types.AttributeValueMemberS{Value: userSub},
		},
	}

	result, err := c.db.GetItem(ctx, input)
	if err != nil {
		return nil, fmt.Errorf("failed to get item: %w", err)
	}

	if result.Item == nil {
		return nil, errors.New("user not found")
	}

	var user models.User
	if err := attributevalue.UnmarshalMap(result.Item, &user); err != nil {
		return nil, fmt.Errorf("failed to unmarshal item: %w", err)
	}

	return &user, nil
}
