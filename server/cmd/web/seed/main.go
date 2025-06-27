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
				{Start: "00:00", End: "23:59"},
			},
			MonthsNorth: []int{1, 2, 3, 11, 12},
			MonthsSouth: []int{5, 6, 7, 8, 9},
		},
		{
			FishID:     "2-pale-chub",
			Name:       "Pale Chub",
			Icon:       "/static/images/fish/icons/pale-chub.png",
			SellPrice:  200,
			ShadowSize: "Tiny",
			ShadowIcon: "/static/images/fish/icons/shadow-tiny.png",
			Location:   "River",
			TimeRanges: []models.TimeRange{
				{Start: "09:00", End: "16:00"},
			},
			MonthsNorth: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12},
			MonthsSouth: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12},
		},
		{
			FishID:     "3-crucian-carp",
			Name:       "Crucian Carp",
			Icon:       "/static/images/fish/icons/crucian-carp.png",
			SellPrice:  160,
			ShadowSize: "Small",
			ShadowIcon: "/static/images/fish/icons/shadow-small.png",
			Location:   "River",
			TimeRanges: []models.TimeRange{
				{Start: "00:00", End: "23:59"},
			},
			MonthsNorth: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12},
			MonthsSouth: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12},
		},
		{
			FishID:     "4-dace",
			Name:       "Dace",
			Icon:       "/static/images/fish/icons/dace.png",
			SellPrice:  240,
			ShadowSize: "Medium",
			ShadowIcon: "/static/images/fish/icons/shadow-medium.png",
			Location:   "River",
			TimeRanges: []models.TimeRange{
				{Start: "16:00", End: "09:00"},
			},
			MonthsNorth: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12},
			MonthsSouth: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12},
		},
		{
			FishID:     "5-carp",
			Name:       "Carp",
			Icon:       "/static/images/fish/icons/carp.png",
			SellPrice:  300,
			ShadowSize: "Large",
			ShadowIcon: "/static/images/fish/icons/shadow-large.png",
			Location:   "Pond",
			TimeRanges: []models.TimeRange{
				{Start: "00:00", End: "23:59"},
			},
			MonthsNorth: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12},
			MonthsSouth: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12},
		},
		{
			FishID:     "6-koi",
			Name:       "Koi",
			Icon:       "/static/images/fish/icons/koi.png",
			SellPrice:  4000,
			ShadowSize: "Large",
			ShadowIcon: "/static/images/fish/icons/shadow-large.png",
			Location:   "Pond",
			TimeRanges: []models.TimeRange{
				{Start: "16:00", End: "09:00"},
			},
			MonthsNorth: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12},
			MonthsSouth: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12},
		},
		{
			FishID:     "7-goldfish",
			Name:       "Goldfish",
			Icon:       "/static/images/fish/icons/goldfish.png",
			SellPrice:  1300,
			ShadowSize: "Tiny",
			ShadowIcon: "/static/images/fish/icons/shadow-tiny.png",
			Location:   "Pond",
			TimeRanges: []models.TimeRange{
				{Start: "00:00", End: "23:59"},
			},
			MonthsNorth: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12},
			MonthsSouth: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12},
		},
		{
			FishID:     "8-pop-eyed-goldfish",
			Name:       "Pop-eyed Goldfish",
			Icon:       "/static/images/fish/icons/pop-eyed-goldfish.png",
			SellPrice:  1300,
			ShadowSize: "Tiny",
			ShadowIcon: "/static/images/fish/icons/shadow-tiny.png",
			Location:   "Pond",
			TimeRanges: []models.TimeRange{
				{Start: "09:00", End: "16:00"},
			},
			MonthsNorth: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12},
			MonthsSouth: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12},
		},
		{
			FishID:     "9-ranchu-goldfish",
			Name:       "Ranchu Goldfish",
			Icon:       "/static/images/fish/icons/ranchu-goldfish.png",
			SellPrice:  4500,
			ShadowSize: "Small",
			ShadowIcon: "/static/images/fish/icons/shadow-small.png",
			Location:   "Pond",
			TimeRanges: []models.TimeRange{
				{Start: "09:00", End: "16:00"},
			},
			MonthsNorth: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12},
			MonthsSouth: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12},
		},
		{
			FishID:     "10-killifish",
			Name:       "Killifish",
			Icon:       "/static/images/fish/icons/killifish.png",
			SellPrice:  300,
			ShadowSize: "Tiny",
			ShadowIcon: "/static/images/fish/icons/shadow-tiny.png",
			Location:   "Pond",
			TimeRanges: []models.TimeRange{
				{Start: "00:00", End: "23:59"},
			},
			MonthsNorth: []int{4, 5, 6, 7, 8},
			MonthsSouth: []int{1, 2, 10, 11, 12},
		},
		{
			FishID:     "11-crawfish",
			Name:       "Crawfish",
			Icon:       "/static/images/fish/icons/crawfish.png",
			SellPrice:  200,
			ShadowSize: "Small",
			ShadowIcon: "/static/images/fish/icons/shadow-small.png",
			Location:   "Pond",
			TimeRanges: []models.TimeRange{
				{Start: "00:00", End: "23:59"},
			},
			MonthsNorth: []int{4, 5, 6, 7, 8, 9},
			MonthsSouth: []int{1, 2, 3, 10, 11, 12},
		},
		{
			FishID:     "12-soft-shelled-turtle",
			Name:       "Soft-shelled Turtle",
			Icon:       "/static/images/fish/icons/soft-shelled-turtle.png",
			SellPrice:  3750,
			ShadowSize: "Large",
			ShadowIcon: "/static/images/fish/icons/shadow-large.png",
			Location:   "River",
			TimeRanges: []models.TimeRange{
				{Start: "16:00", End: "09:00"},
			},
			MonthsNorth: []int{8, 9},
			MonthsSouth: []int{2, 3},
		},
		{
			FishID:     "13-snapping-turtle",
			Name:       "Snapping Turtle",
			Icon:       "/static/images/fish/icons/snapping-turtle.png",
			SellPrice:  5000,
			ShadowSize: "Large",
			ShadowIcon: "/static/images/fish/icons/shadow-large.png",
			Location:   "River",
			TimeRanges: []models.TimeRange{
				{Start: "21:00", End: "04:00"},
			},
			MonthsNorth: []int{4, 5, 6, 7, 8, 9, 10},
			MonthsSouth: []int{1, 2, 3, 4, 10, 11, 12},
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
