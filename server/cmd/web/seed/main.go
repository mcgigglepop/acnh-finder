package main

import (
	"context"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	sdkdynamodb "github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"

	"github.com/mcgigglepop/acnh-finder/server/internal/models"
)

func main() {
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	client := sdkdynamodb.NewFromConfig(cfg)
	tableName := "Fish"

	fishList := []models.Fish{
		{
			FishID:     "1-bitterling",
			Name:       "Bitterling",
			Icon:       "/static/images/fish/icons/bitterling.png",
			SellPrice:  900,
			ShadowSize: "Tiny",
			ShadowIcon: "/static/images/fish/icons/shadow-tiny.png",
			Location:   "River",
			TimeRanges: []models.TimeRange{
				{Start: "00:00", End: "23:59"}, // All day
			},
			MonthsNorth: []int{1, 2, 3, 11, 12},
			MonthsSouth: []int{5, 6, 7, 8, 9},
		},
	}

	for _, fish := range fishList {
		item, err := attributevalue.MarshalMap(fish)
		if err != nil {
			log.Printf("error marshaling fish %s: %v", fish.FishID, err)
			continue
		}

		_, err = client.PutItem(context.TODO(), &sdkdynamodb.PutItemInput{
			TableName: aws.String(tableName),
			Item:      item,
		})

		if err != nil {
			log.Printf("error inserting fish %s: %v", fish.FishID, err)
		} else {
			log.Printf("successfully inserted: %s", fish.Name)
		}
	}
}
