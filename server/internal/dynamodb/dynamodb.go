package dynamodb

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"time"

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

func (c *DDBClient) UpdateUserHemisphere(ctx context.Context, userSub string, hemisphere string) error {
	input := &sdkdynamodb.UpdateItemInput{
		TableName: aws.String(c.tableName),
		Key: map[string]types.AttributeValue{
			"user_id": &types.AttributeValueMemberS{Value: userSub},
		},
		UpdateExpression: aws.String("SET hemisphere = :h"),
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":h": &types.AttributeValueMemberS{Value: hemisphere},
		},
		ReturnValues: types.ReturnValueUpdatedNew,
	}

	_, err := c.db.UpdateItem(ctx, input)
	if err != nil {
		return fmt.Errorf("failed to update hemisphere: %w", err)
	}

	return nil
}

func containsInt(slice []int, val int) bool {
	for _, v := range slice {
		if v == val {
			return true
		}
	}
	return false
}

func isTimeInRange(current, start, end string) bool {
	layout := "15:04"
	now, _ := time.Parse(layout, current)
	from, _ := time.Parse(layout, start)
	to, _ := time.Parse(layout, end)

	if from.Before(to) {
		return now.After(from) && now.Before(to)
	}
	// Overnight wraparound
	return now.After(from) || now.Before(to)
}

func isFishAvailable(seasons []models.SeasonalAvailability, month int, hour string) bool {
	for _, s := range seasons {
		if containsInt(s.Months, month) {
			for _, tr := range s.TimeRanges {
				if isTimeInRange(hour, tr.Start, tr.End) {
					return true
				}
			}
		}
	}
	return false
}

func (c *DDBClient) getUserCaughtFishMap(ctx context.Context, userID string) (map[string]bool, error) {
	out, err := c.db.Query(ctx, &sdkdynamodb.QueryInput{
		TableName:              aws.String("UserFish"),
		KeyConditionExpression: aws.String("PK = :pk"),
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":pk": &types.AttributeValueMemberS{Value: fmt.Sprintf("USER#%s", userID)},
		},
	})
	if err != nil {
		return nil, err
	}

	caughtMap := make(map[string]bool)
	for _, item := range out.Items {
		var record struct {
			SK     string `dynamodbav:"SK"`
			Caught bool   `dynamodbav:"Caught"`
		}
		if err := attributevalue.UnmarshalMap(item, &record); err != nil {
			return nil, err
		}

		// Extract fish ID from SK like "FISH#abc123"
		parts := strings.Split(record.SK, "#")
		if len(parts) == 2 {
			fishID := parts[1]
			caughtMap[fishID] = record.Caught
		}
	}

	return caughtMap, nil
}

func (c *DDBClient) ListAvailableFish(ctx context.Context, userID string, month int, hour string, hemisphere string) ([]models.Fish, error) {
	// 1. Scan all fish 
	out, err := c.db.Scan(ctx, &sdkdynamodb.ScanInput{
		TableName: aws.String(c.tableName),
	})
	if err != nil {
		return nil, fmt.Errorf("scan failed: %w", err)
	}

	var allFish []models.Fish
	if err := attributevalue.UnmarshalListOfMaps(out.Items, &allFish); err != nil {
		return nil, fmt.Errorf("unmarshal failed: %w", err)
	}

	// 2. Fetch user caught fish map
	userCaughtMap, err := c.getUserCaughtFishMap(ctx, userID)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch caught fish map: %w", err)
	}

	// 3. Filter and merge
	var availableFish []models.Fish
	for _, fish := range allFish {
		var seasons []models.SeasonalAvailability
		if hemisphere == "north" {
			seasons = fish.NorthAvailability
		} else {
			seasons = fish.SouthAvailability
		}

		if isFishAvailable(seasons, month, hour) {
			// Check if user caught this fish
			fish.Caught = userCaughtMap[fish.FishID]
			availableFish = append(availableFish, fish)
		}
	}

	return availableFish, nil
}
