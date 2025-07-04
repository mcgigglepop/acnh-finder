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
			NorthAvailability: []models.SeasonalAvailability{
				{
					Months: []int{1, 2, 3, 11, 12},
					TimeRanges: []models.TimeRange{
						{Start: "00:00", End: "23:59"},
					},
				},
			},
			SouthAvailability: []models.SeasonalAvailability{
				{
					Months: []int{5, 6, 7, 8, 9},
					TimeRanges: []models.TimeRange{
						{Start: "00:00", End: "23:59"},
					},
				},
			},
		},
		{
			FishID:     "2-pale-chub",
			Name:       "Pale Chub",
			Icon:       "/static/images/fish/icons/pale-chub.png",
			SellPrice:  200,
			ShadowSize: "Tiny",
			ShadowIcon: "/static/images/fish/icons/shadow-tiny.png",
			Location:   "River",
			NorthAvailability: []models.SeasonalAvailability{
				{
					Months: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12},
					TimeRanges: []models.TimeRange{
						{Start: "09:00", End: "16:00"},
					},
				},
			},
			SouthAvailability: []models.SeasonalAvailability{
				{
					Months: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12},
					TimeRanges: []models.TimeRange{
						{Start: "09:00", End: "16:00"},
					},
				},
			},
		},
		{
			FishID:     "3-crucian-carp",
			Name:       "Crucian Carp",
			Icon:       "/static/images/fish/icons/crucian-carp.png",
			SellPrice:  160,
			ShadowSize: "Small",
			ShadowIcon: "/static/images/fish/icons/shadow-small.png",
			Location:   "River",
			NorthAvailability: []models.SeasonalAvailability{
				{
					Months: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12},
					TimeRanges: []models.TimeRange{
						{Start: "00:00", End: "23:59"},
					},
				},
			},
			SouthAvailability: []models.SeasonalAvailability{
				{
					Months: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12},
					TimeRanges: []models.TimeRange{
						{Start: "00:00", End: "23:59"},
					},
				},
			},
		},
		{
			FishID:     "4-dace",
			Name:       "Dace",
			Icon:       "/static/images/fish/icons/dace.png",
			SellPrice:  240,
			ShadowSize: "Medium",
			ShadowIcon: "/static/images/fish/icons/shadow-medium.png",
			Location:   "River",
			NorthAvailability: []models.SeasonalAvailability{
				{
					Months: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12},
					TimeRanges: []models.TimeRange{
						{Start: "16:00", End: "09:00"},
					},
				},
			},
			SouthAvailability: []models.SeasonalAvailability{
				{
					Months: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12},
					TimeRanges: []models.TimeRange{
						{Start: "16:00", End: "09:00"},
					},
				},
			},
		},
		{
			FishID:     "5-carp",
			Name:       "Carp",
			Icon:       "/static/images/fish/icons/carp.png",
			SellPrice:  300,
			ShadowSize: "Large",
			ShadowIcon: "/static/images/fish/icons/shadow-large.png",
			Location:   "Pond",
			NorthAvailability: []models.SeasonalAvailability{
				{
					Months: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12},
					TimeRanges: []models.TimeRange{
						{Start: "00:00", End: "23:59"},
					},
				},
			},
			SouthAvailability: []models.SeasonalAvailability{
				{
					Months: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12},
					TimeRanges: []models.TimeRange{
						{Start: "00:00", End: "23:59"},
					},
				},
			},
		},
		{
			FishID:     "6-koi",
			Name:       "Koi",
			Icon:       "/static/images/fish/icons/koi.png",
			SellPrice:  4000,
			ShadowSize: "Large",
			ShadowIcon: "/static/images/fish/icons/shadow-large.png",
			Location:   "Pond",
			NorthAvailability: []models.SeasonalAvailability{
				{
					Months: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12},
					TimeRanges: []models.TimeRange{
						{Start: "16:00", End: "09:00"},
					},
				},
			},
			SouthAvailability: []models.SeasonalAvailability{
				{
					Months: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12},
					TimeRanges: []models.TimeRange{
						{Start: "16:00", End: "09:00"},
					},
				},
			},
		},
		{
			FishID:     "7-goldfish",
			Name:       "Goldfish",
			Icon:       "/static/images/fish/icons/goldfish.png",
			SellPrice:  1300,
			ShadowSize: "Tiny",
			ShadowIcon: "/static/images/fish/icons/shadow-tiny.png",
			Location:   "Pond",
			NorthAvailability: []models.SeasonalAvailability{
				{
					Months: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12},
					TimeRanges: []models.TimeRange{
						{Start: "00:00", End: "23:59"},
					},
				},
			},
			SouthAvailability: []models.SeasonalAvailability{
				{
					Months: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12},
					TimeRanges: []models.TimeRange{
						{Start: "00:00", End: "23:59"},
					},
				},
			},
		},
		{
			FishID:     "8-pop-eyed-goldfish",
			Name:       "Pop-eyed Goldfish",
			Icon:       "/static/images/fish/icons/pop-eyed-goldfish.png",
			SellPrice:  1300,
			ShadowSize: "Tiny",
			ShadowIcon: "/static/images/fish/icons/shadow-tiny.png",
			Location:   "Pond",
			NorthAvailability: []models.SeasonalAvailability{
				{
					Months: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12},
					TimeRanges: []models.TimeRange{
						{Start: "09:00", End: "16:00"},
					},
				},
			},
			SouthAvailability: []models.SeasonalAvailability{
				{
					Months: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12},
					TimeRanges: []models.TimeRange{
						{Start: "09:00", End: "16:00"},
					},
				},
			},
		},
		{
			FishID:     "9-ranchu-goldfish",
			Name:       "Ranchu Goldfish",
			Icon:       "/static/images/fish/icons/ranchu-goldfish.png",
			SellPrice:  4500,
			ShadowSize: "Small",
			ShadowIcon: "/static/images/fish/icons/shadow-small.png",
			Location:   "Pond",
			NorthAvailability: []models.SeasonalAvailability{
				{
					Months: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12},
					TimeRanges: []models.TimeRange{
						{Start: "09:00", End: "16:00"},
					},
				},
			},
			SouthAvailability: []models.SeasonalAvailability{
				{
					Months: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12},
					TimeRanges: []models.TimeRange{
						{Start: "09:00", End: "16:00"},
					},
				},
			},
		},
		{
			FishID:     "10-killifish",
			Name:       "Killifish",
			Icon:       "/static/images/fish/icons/killifish.png",
			SellPrice:  300,
			ShadowSize: "Tiny",
			ShadowIcon: "/static/images/fish/icons/shadow-tiny.png",
			Location:   "Pond",
			NorthAvailability: []models.SeasonalAvailability{
				{
					Months: []int{4, 5, 6, 7, 8},
					TimeRanges: []models.TimeRange{
						{Start: "00:00", End: "23:59"},
					},
				},
			},
			SouthAvailability: []models.SeasonalAvailability{
				{
					Months: []int{1, 2, 10, 11, 12},
					TimeRanges: []models.TimeRange{
						{Start: "00:00", End: "23:59"},
					},
				},
			},
		},
		{
			FishID:     "11-crawfish",
			Name:       "Crawfish",
			Icon:       "/static/images/fish/icons/crawfish.png",
			SellPrice:  200,
			ShadowSize: "Small",
			ShadowIcon: "/static/images/fish/icons/shadow-small.png",
			Location:   "Pond",
			NorthAvailability: []models.SeasonalAvailability{
				{
					Months: []int{4, 5, 6, 7, 8, 9},
					TimeRanges: []models.TimeRange{
						{Start: "00:00", End: "23:59"},
					},
				},
			},
			SouthAvailability: []models.SeasonalAvailability{
				{
					Months: []int{1, 2, 3, 10, 11, 12},
					TimeRanges: []models.TimeRange{
						{Start: "00:00", End: "23:59"},
					},
				},
			},
		},
		{
			FishID:     "12-soft-shelled-turtle",
			Name:       "Soft-shelled Turtle",
			Icon:       "/static/images/fish/icons/soft-shelled-turtle.png",
			SellPrice:  3750,
			ShadowSize: "Large",
			ShadowIcon: "/static/images/fish/icons/shadow-large.png",
			Location:   "River",
			NorthAvailability: []models.SeasonalAvailability{
				{
					Months: []int{8, 9},
					TimeRanges: []models.TimeRange{
						{Start: "16:00", End: "09:00"},
					},
				},
			},
			SouthAvailability: []models.SeasonalAvailability{
				{
					Months: []int{2, 3},
					TimeRanges: []models.TimeRange{
						{Start: "16:00", End: "09:00"},
					},
				},
			},
		},
		{
			FishID:     "13-snapping-turtle",
			Name:       "Snapping Turtle",
			Icon:       "/static/images/fish/icons/snapping-turtle.png",
			SellPrice:  5000,
			ShadowSize: "Large",
			ShadowIcon: "/static/images/fish/icons/shadow-large.png",
			Location:   "River",
			NorthAvailability: []models.SeasonalAvailability{
				{
					Months: []int{4, 5, 6, 7, 8, 9, 10},
					TimeRanges: []models.TimeRange{
						{Start: "21:00", End: "04:00"},
					},
				},
			},
			SouthAvailability: []models.SeasonalAvailability{
				{
					Months: []int{1, 2, 3, 4, 10, 11, 12},
					TimeRanges: []models.TimeRange{
						{Start: "21:00", End: "04:00"},
					},
				},
			},
		},
		{
			FishID:     "14-tadpole",
			Name:       "Tadpole",
			Icon:       "/static/images/fish/icons/tadpole.png",
			SellPrice:  100,
			ShadowSize: "Tiny",
			ShadowIcon: "/static/images/fish/icons/shadow-tiny.png",
			Location:   "Pond",
			NorthAvailability: []models.SeasonalAvailability{
				{
					Months: []int{3, 4, 5, 6, 7},
					TimeRanges: []models.TimeRange{
						{Start: "00:00", End: "23:59"},
					},
				},
			},
			SouthAvailability: []models.SeasonalAvailability{
				{
					Months: []int{1, 9, 10, 11, 12},
					TimeRanges: []models.TimeRange{
						{Start: "00:00", End: "23:59"},
					},
				},
			},
		},
		{
			FishID:     "15-frog",
			Name:       "Frog",
			Icon:       "/static/images/fish/icons/frog.png",
			SellPrice:  120,
			ShadowSize: "Small",
			ShadowIcon: "/static/images/fish/icons/shadow-small.png",
			Location:   "Pond",
			NorthAvailability: []models.SeasonalAvailability{
				{
					Months: []int{5, 6, 7, 8},
					TimeRanges: []models.TimeRange{
						{Start: "00:00", End: "23:59"},
					},
				},
			},
			SouthAvailability: []models.SeasonalAvailability{
				{
					Months: []int{1, 2, 11, 12},
					TimeRanges: []models.TimeRange{
						{Start: "00:00", End: "23:59"},
					},
				},
			},
		},
		{
			FishID:     "16-freshwater-goby",
			Name:       "Freshwater Goby",
			Icon:       "/static/images/fish/icons/freshwater-goby.png",
			SellPrice:  400,
			ShadowSize: "Small",
			ShadowIcon: "/static/images/fish/icons/shadow-small.png",
			Location:   "River",
			NorthAvailability: []models.SeasonalAvailability{
				{
					Months: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12},
					TimeRanges: []models.TimeRange{
						{Start: "16:00", End: "09:00"},
					},
				},
			},
			SouthAvailability: []models.SeasonalAvailability{
				{
					Months: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12},
					TimeRanges: []models.TimeRange{
						{Start: "16:00", End: "09:00"},
					},
				},
			},
		},
		{
			FishID:     "17-loach",
			Name:       "Loach",
			Icon:       "/static/images/fish/icons/loach.png",
			SellPrice:  400,
			ShadowSize: "Small",
			ShadowIcon: "/static/images/fish/icons/shadow-small.png",
			Location:   "River",
			NorthAvailability: []models.SeasonalAvailability{
				{
					Months: []int{3, 4, 5},
					TimeRanges: []models.TimeRange{
						{Start: "00:00", End: "23:59"},
					},
				},
			},
			SouthAvailability: []models.SeasonalAvailability{
				{
					Months: []int{9, 10, 11},
					TimeRanges: []models.TimeRange{
						{Start: "00:00", End: "23:59"},
					},
				},
			},
		},
		{
			FishID:     "18-catfish",
			Name:       "Catfish",
			Icon:       "/static/images/fish/icons/catfish.png",
			SellPrice:  800,
			ShadowSize: "Large",
			ShadowIcon: "/static/images/fish/icons/shadow-large.png",
			Location:   "Pond",
			NorthAvailability: []models.SeasonalAvailability{
				{
					Months: []int{5, 6, 7, 8, 9, 10},
					TimeRanges: []models.TimeRange{
						{Start: "16:00", End: "09:00"},
					},
				},
			},
			SouthAvailability: []models.SeasonalAvailability{
				{
					Months: []int{1, 2, 3, 4, 11, 12},
					TimeRanges: []models.TimeRange{
						{Start: "16:00", End: "09:00"},
					},
				},
			},
		},
		{
			FishID:     "19-giant-snakehead",
			Name:       "Giant Snakehead",
			Icon:       "/static/images/fish/icons/giant-snakehead.png",
			SellPrice:  5500,
			ShadowSize: "Large",
			ShadowIcon: "/static/images/fish/icons/shadow-large.png",
			Location:   "Pond",
			NorthAvailability: []models.SeasonalAvailability{
				{
					Months: []int{6, 7, 8},
					TimeRanges: []models.TimeRange{
						{Start: "09:00", End: "16:00"},
					},
				},
			},
			SouthAvailability: []models.SeasonalAvailability{
				{
					Months: []int{1, 2, 12},
					TimeRanges: []models.TimeRange{
						{Start: "09:00", End: "16:00"},
					},
				},
			},
		},
		{
			FishID:     "20-bluegill",
			Name:       "Bluegill",
			Icon:       "/static/images/fish/icons/bluegill.png",
			SellPrice:  180,
			ShadowSize: "Small",
			ShadowIcon: "/static/images/fish/icons/shadow-small.png",
			Location:   "River",
			NorthAvailability: []models.SeasonalAvailability{
				{
					Months: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12},
					TimeRanges: []models.TimeRange{
						{Start: "09:00", End: "16:00"},
					},
				},
			},
			SouthAvailability: []models.SeasonalAvailability{
				{
					Months: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12},
					TimeRanges: []models.TimeRange{
						{Start: "09:00", End: "16:00"},
					},
				},
			},
		},
		{
			FishID:     "21-yellow-perch",
			Name:       "Yellow Perch",
			Icon:       "/static/images/fish/icons/yellow-perch.png",
			SellPrice:  300,
			ShadowSize: "Medium",
			ShadowIcon: "/static/images/fish/icons/shadow-medium.png",
			Location:   "River",
			NorthAvailability: []models.SeasonalAvailability{
				{
					Months: []int{1, 2, 3, 10, 11, 12},
					TimeRanges: []models.TimeRange{
						{Start: "00:00", End: "23:59"},
					},
				},
			},
			SouthAvailability: []models.SeasonalAvailability{
				{
					Months: []int{4, 5, 6, 7, 8, 9},
					TimeRanges: []models.TimeRange{
						{Start: "00:00", End: "23:59"},
					},
				},
			},
		},
		{
			FishID:     "22-black-bass",
			Name:       "Black Bass",
			Icon:       "/static/images/fish/icons/black-bass.png",
			SellPrice:  400,
			ShadowSize: "Large",
			ShadowIcon: "/static/images/fish/icons/shadow-large.png",
			Location:   "River",
			NorthAvailability: []models.SeasonalAvailability{
				{
					Months: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12},
					TimeRanges: []models.TimeRange{
						{Start: "00:00", End: "23:59"},
					},
				},
			},
			SouthAvailability: []models.SeasonalAvailability{
				{
					Months: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12},
					TimeRanges: []models.TimeRange{
						{Start: "00:00", End: "23:59"},
					},
				},
			},
		},
		{
			FishID:     "23-tilapia",
			Name:       "Tilapia",
			Icon:       "/static/images/fish/icons/tilapia.png",
			SellPrice:  800,
			ShadowSize: "Medium",
			ShadowIcon: "/static/images/fish/icons/shadow-medium.png",
			Location:   "River",
			NorthAvailability: []models.SeasonalAvailability{
				{
					Months: []int{6, 7, 8, 9, 10},
					TimeRanges: []models.TimeRange{
						{Start: "00:00", End: "23:59"},
					},
				},
			},
			SouthAvailability: []models.SeasonalAvailability{
				{
					Months: []int{1, 2, 3, 4, 12},
					TimeRanges: []models.TimeRange{
						{Start: "00:00", End: "23:59"},
					},
				},
			},
		},
		{
			FishID:     "24-pike",
			Name:       "Pike",
			Icon:       "/static/images/fish/icons/pike.png",
			SellPrice:  1800,
			ShadowSize: "Very Large",
			ShadowIcon: "/static/images/fish/icons/shadow-very-large.png",
			Location:   "River",
			NorthAvailability: []models.SeasonalAvailability{
				{
					Months: []int{9, 10, 11, 12},
					TimeRanges: []models.TimeRange{
						{Start: "00:00", End: "23:59"},
					},
				},
			},
			SouthAvailability: []models.SeasonalAvailability{
				{
					Months: []int{3, 4, 5, 6},
					TimeRanges: []models.TimeRange{
						{Start: "00:00", End: "23:59"},
					},
				},
			},
		},
		{
			FishID:     "25-pond-smelt",
			Name:       "Pond Smelt",
			Icon:       "/static/images/fish/icons/pond-smelt.png",
			SellPrice:  400,
			ShadowSize: "Small",
			ShadowIcon: "/static/images/fish/icons/shadow-small.png",
			Location:   "River",
			NorthAvailability: []models.SeasonalAvailability{
				{
					Months: []int{1, 2, 12},
					TimeRanges: []models.TimeRange{
						{Start: "00:00", End: "23:59"},
					},
				},
			},
			SouthAvailability: []models.SeasonalAvailability{
				{
					Months: []int{6, 7, 8},
					TimeRanges: []models.TimeRange{
						{Start: "00:00", End: "23:59"},
					},
				},
			},
		},
		{
			FishID:     "26-sweetfish",
			Name:       "Sweetfish",
			Icon:       "/static/images/fish/icons/sweetfish.png",
			SellPrice:  900,
			ShadowSize: "Medium",
			ShadowIcon: "/static/images/fish/icons/shadow-medium.png",
			Location:   "River",
			NorthAvailability: []models.SeasonalAvailability{
				{
					Months: []int{7, 8, 9},
					TimeRanges: []models.TimeRange{
						{Start: "00:00", End: "23:59"},
					},
				},
			},
			SouthAvailability: []models.SeasonalAvailability{
				{
					Months: []int{1, 2, 3},
					TimeRanges: []models.TimeRange{
						{Start: "00:00", End: "23:59"},
					},
				},
			},
		},
		{
			FishID:     "27-cherry-salmon",
			Name:       "Cherry Salmon",
			Icon:       "/static/images/fish/icons/cherry-salmon.png",
			SellPrice:  1000,
			ShadowSize: "Medium",
			ShadowIcon: "/static/images/fish/icons/shadow-medium.png",
			Location:   "River (Clifftop)",
			NorthAvailability: []models.SeasonalAvailability{
				{
					Months: []int{3, 4, 5, 6},
					TimeRanges: []models.TimeRange{
						{Start: "16:00", End: "09:00"},
					},
				},
				{
					Months: []int{9, 10, 11},
					TimeRanges: []models.TimeRange{
						{Start: "00:00", End: "23:59"},
					},
				},
			},
			SouthAvailability: []models.SeasonalAvailability{
				{
					Months: []int{3, 4, 5},
					TimeRanges: []models.TimeRange{
						{Start: "00:00", End: "23:59"},
					},
				},
				{
					Months: []int{9, 10, 11},
					TimeRanges: []models.TimeRange{
						{Start: "16:00", End: "09:00"},
					},
				},
			},
		},
		{
			FishID:     "28-char",
			Name:       "Char",
			Icon:       "/static/images/fish/icons/char.png",
			SellPrice:  3800,
			ShadowSize: "Medium",
			ShadowIcon: "/static/images/fish/icons/shadow-medium.png",
			Location:   "River (Clifftop)",
			NorthAvailability: []models.SeasonalAvailability{
				{
					Months: []int{3, 4, 5, 6},
					TimeRanges: []models.TimeRange{
						{Start: "16:00", End: "09:00"},
					},
				},
				{
					Months: []int{9, 10, 11},
					TimeRanges: []models.TimeRange{
						{Start: "00:00", End: "23:59"},
					},
				},
			},
			SouthAvailability: []models.SeasonalAvailability{
				{
					Months: []int{3, 4, 5},
					TimeRanges: []models.TimeRange{
						{Start: "00:00", End: "23:59"},
					},
				},
				{
					Months: []int{9, 10, 11},
					TimeRanges: []models.TimeRange{
						{Start: "16:00", End: "09:00"},
					},
				},
			},
		},
		{
			FishID:     "29-golden-trout",
			Name:       "Golden Trout",
			Icon:       "/static/images/fish/icons/golden-trout.png",
			SellPrice:  15000,
			ShadowSize: "Medium",
			ShadowIcon: "/static/images/fish/icons/shadow-medium.png",
			Location:   "River (Clifftop)",
			NorthAvailability: []models.SeasonalAvailability{
				{
					Months: []int{3, 4, 5, 9, 10, 11},
					TimeRanges: []models.TimeRange{
						{Start: "16:00", End: "09:00"},
					},
				},
			},
			SouthAvailability: []models.SeasonalAvailability{
				{
					Months: []int{3, 4, 5, 9, 10, 11},
					TimeRanges: []models.TimeRange{
						{Start: "16:00", End: "09:00"},
					},
				},
			},
		},
		{
			FishID:     "30-stringfish",
			Name:       "Stringfish",
			Icon:       "/static/images/fish/icons/stringfish.png",
			SellPrice:  15000,
			ShadowSize: "Very Large",
			ShadowIcon: "/static/images/fish/icons/shadow-very-large.png",
			Location:   "River (Clifftop)",
			NorthAvailability: []models.SeasonalAvailability{
				{
					Months: []int{1, 2, 3, 12},
					TimeRanges: []models.TimeRange{
						{Start: "16:00", End: "09:00"},
					},
				},
			},
			SouthAvailability: []models.SeasonalAvailability{
				{
					Months: []int{6, 7, 8, 9},
					TimeRanges: []models.TimeRange{
						{Start: "16:00", End: "09:00"},
					},
				},
			},
		},
		{
			FishID:     "31-salmon",
			Name:       "Salmon",
			Icon:       "/static/images/fish/icons/salmon.png",
			SellPrice:  700,
			ShadowSize: "Large",
			ShadowIcon: "/static/images/fish/icons/shadow-large.png",
			Location:   "River (Mouth)",
			NorthAvailability: []models.SeasonalAvailability{
				{
					Months: []int{9},
					TimeRanges: []models.TimeRange{
						{Start: "00:00", End: "23:59"},
					},
				},
			},
			SouthAvailability: []models.SeasonalAvailability{
				{
					Months: []int{3},
					TimeRanges: []models.TimeRange{
						{Start: "00:00", End: "23:59"},
					},
				},
			},
		},
		{
			FishID:     "32-king-salmon",
			Name:       "King Salmon",
			Icon:       "/static/images/fish/icons/king-salmon.png",
			SellPrice:  1800,
			ShadowSize: "Very Large",
			ShadowIcon: "/static/images/fish/icons/shadow-very-large.png",
			Location:   "River (Mouth)",
			NorthAvailability: []models.SeasonalAvailability{
				{
					Months: []int{9},
					TimeRanges: []models.TimeRange{
						{Start: "00:00", End: "23:59"},
					},
				},
			},
			SouthAvailability: []models.SeasonalAvailability{
				{
					Months: []int{3},
					TimeRanges: []models.TimeRange{
						{Start: "00:00", End: "23:59"},
					},
				},
			},
		},
		{
			FishID:     "33-mitten-crab",
			Name:       "Mitten Crab",
			Icon:       "/static/images/fish/icons/mitten-crab.png",
			SellPrice:  2000,
			ShadowSize: "Small",
			ShadowIcon: "/static/images/fish/icons/shadow-small.png",
			Location:   "River",
			NorthAvailability: []models.SeasonalAvailability{
				{
					Months: []int{9, 10, 11},
					TimeRanges: []models.TimeRange{
						{Start: "16:00", End: "09:00"},
					},
				},
			},
			SouthAvailability: []models.SeasonalAvailability{
				{
					Months: []int{3, 4, 5},
					TimeRanges: []models.TimeRange{
						{Start: "16:00", End: "09:00"},
					},
				},
			},
		},
		{
			FishID:     "34-guppy",
			Name:       "Guppy",
			Icon:       "/static/images/fish/icons/guppy.png",
			SellPrice:  1300,
			ShadowSize: "Tiny",
			ShadowIcon: "/static/images/fish/icons/shadow-tiny.png",
			Location:   "River",
			NorthAvailability: []models.SeasonalAvailability{
				{
					Months: []int{4, 5, 6, 7, 8, 9, 10, 11},
					TimeRanges: []models.TimeRange{
						{Start: "09:00", End: "16:00"},
					},
				},
			},
			SouthAvailability: []models.SeasonalAvailability{
				{
					Months: []int{1, 2, 3, 4, 5, 10, 11, 12},
					TimeRanges: []models.TimeRange{
						{Start: "09:00", End: "16:00"},
					},
				},
			},
		},
		{
			FishID:     "35-nibble-fish",
			Name:       "Nibble Fish",
			Icon:       "/static/images/fish/icons/nibble-fish.png",
			SellPrice:  1500,
			ShadowSize: "Tiny",
			ShadowIcon: "/static/images/fish/icons/shadow-tiny.png",
			Location:   "River",
			NorthAvailability: []models.SeasonalAvailability{
				{
					Months: []int{5, 6, 7, 8, 9},
					TimeRanges: []models.TimeRange{
						{Start: "09:00", End: "16:00"},
					},
				},
			},
			SouthAvailability: []models.SeasonalAvailability{
				{
					Months: []int{1, 2, 3, 11, 12},
					TimeRanges: []models.TimeRange{
						{Start: "09:00", End: "16:00"},
					},
				},
			},
		},
		{
			FishID:     "36-angelfish",
			Name:       "Angelfish",
			Icon:       "/static/images/fish/icons/angelfish.png",
			SellPrice:  3000,
			ShadowSize: "Small",
			ShadowIcon: "/static/images/fish/icons/shadow-small.png",
			Location:   "River",
			NorthAvailability: []models.SeasonalAvailability{
				{
					Months: []int{5, 6, 7, 8, 9, 10},
					TimeRanges: []models.TimeRange{
						{Start: "16:00", End: "09:00"},
					},
				},
			},
			SouthAvailability: []models.SeasonalAvailability{
				{
					Months: []int{1, 2, 3, 4, 11, 12},
					TimeRanges: []models.TimeRange{
						{Start: "16:00", End: "09:00"},
					},
				},
			},
		},
		{
			FishID:     "37-betta",
			Name:       "Betta",
			Icon:       "/static/images/fish/icons/betta.png",
			SellPrice:  2500,
			ShadowSize: "Small",
			ShadowIcon: "/static/images/fish/icons/shadow-small.png",
			Location:   "River",
			NorthAvailability: []models.SeasonalAvailability{
				{
					Months: []int{5, 6, 7, 8, 9, 10},
					TimeRanges: []models.TimeRange{
						{Start: "09:00", End: "16:00"},
					},
				},
			},
			SouthAvailability: []models.SeasonalAvailability{
				{
					Months: []int{1, 2, 3, 4, 11, 12},
					TimeRanges: []models.TimeRange{
						{Start: "09:00", End: "16:00"},
					},
				},
			},
		},
		{
			FishID:     "38-neon-tetra",
			Name:       "Neon Tetra",
			Icon:       "/static/images/fish/icons/neon-tetra.png",
			SellPrice:  500,
			ShadowSize: "Tiny",
			ShadowIcon: "/static/images/fish/icons/shadow-tiny.png",
			Location:   "River",
			NorthAvailability: []models.SeasonalAvailability{
				{
					Months: []int{4, 5, 6, 7, 8, 9, 10, 11},
					TimeRanges: []models.TimeRange{
						{Start: "09:00", End: "16:00"},
					},
				},
			},
			SouthAvailability: []models.SeasonalAvailability{
				{
					Months: []int{1, 2, 3, 4, 5, 10, 11, 12},
					TimeRanges: []models.TimeRange{
						{Start: "09:00", End: "16:00"},
					},
				},
			},
		},
		{
			FishID:     "39-rainbowfish",
			Name:       "Rainbowfish",
			Icon:       "/static/images/fish/icons/rainbowfish.png",
			SellPrice:  800,
			ShadowSize: "Tiny",
			ShadowIcon: "/static/images/fish/icons/shadow-tiny.png",
			Location:   "River",
			NorthAvailability: []models.SeasonalAvailability{
				{
					Months: []int{5, 6, 7, 8, 9, 10},
					TimeRanges: []models.TimeRange{
						{Start: "09:00", End: "16:00"},
					},
				},
			},
			SouthAvailability: []models.SeasonalAvailability{
				{
					Months: []int{1, 2, 3, 4, 11, 12},
					TimeRanges: []models.TimeRange{
						{Start: "09:00", End: "16:00"},
					},
				},
			},
		},
		{
			FishID:     "40-piranha",
			Name:       "Piranha",
			Icon:       "/static/images/fish/icons/piranha.png",
			SellPrice:  2500,
			ShadowSize: "Small",
			ShadowIcon: "/static/images/fish/icons/shadow-small.png",
			Location:   "River",
			NorthAvailability: []models.SeasonalAvailability{
				{
					Months: []int{6, 7, 8, 9},
					TimeRanges: []models.TimeRange{
						{Start: "09:00", End: "16:00"},
						{Start: "21:00", End: "04:00"},
					},
				},
			},
			SouthAvailability: []models.SeasonalAvailability{
				{
					Months: []int{1, 2, 3, 12},
					TimeRanges: []models.TimeRange{
						{Start: "09:00", End: "16:00"},
						{Start: "21:00", End: "04:00"},
					},
				},
			},
		},
		{
			FishID:     "41-arowana",
			Name:       "Arowana",
			Icon:       "/static/images/fish/icons/arowana.png",
			SellPrice:  10000,
			ShadowSize: "Large",
			ShadowIcon: "/static/images/fish/icons/shadow-large.png",
			Location:   "River",
			NorthAvailability: []models.SeasonalAvailability{
				{
					Months: []int{6, 7, 8, 9},
					TimeRanges: []models.TimeRange{
						{Start: "16:00", End: "09:00"},
					},
				},
			},
			SouthAvailability: []models.SeasonalAvailability{
				{
					Months: []int{1, 2, 3, 12},
					TimeRanges: []models.TimeRange{
						{Start: "16:00", End: "09:00"},
					},
				},
			},
		},
		{
			FishID:     "42-dorado",
			Name:       "Dorado",
			Icon:       "/static/images/fish/icons/dorado.png",
			SellPrice:  15000,
			ShadowSize: "Very Large",
			ShadowIcon: "/static/images/fish/icons/shadow-very-large.png",
			Location:   "River",
			NorthAvailability: []models.SeasonalAvailability{
				{
					Months: []int{6, 7, 8, 9},
					TimeRanges: []models.TimeRange{
						{Start: "04:00", End: "21:00"},
					},
				},
			},
			SouthAvailability: []models.SeasonalAvailability{
				{
					Months: []int{1, 2, 3, 12},
					TimeRanges: []models.TimeRange{
						{Start: "04:00", End: "21:00"},
					},
				},
			},
		},
		{
			FishID:     "43-gar",
			Name:       "Gar",
			Icon:       "/static/images/fish/icons/gar.png",
			SellPrice:  6000,
			ShadowSize: "Very Large",
			ShadowIcon: "/static/images/fish/icons/shadow-very-large.png",
			Location:   "Pond",
			NorthAvailability: []models.SeasonalAvailability{
				{
					Months: []int{6, 7, 8, 9},
					TimeRanges: []models.TimeRange{
						{Start: "16:00", End: "09:00"},
					},
				},
			},
			SouthAvailability: []models.SeasonalAvailability{
				{
					Months: []int{1, 2, 3, 12},
					TimeRanges: []models.TimeRange{
						{Start: "16:00", End: "09:00"},
					},
				},
			},
		},
		{
			FishID:     "44-arapaima",
			Name:       "Arapaima",
			Icon:       "/static/images/fish/icons/arapaima.png",
			SellPrice:  10000,
			ShadowSize: "Huge",
			ShadowIcon: "/static/images/fish/icons/shadow-huge.png",
			Location:   "River",
			NorthAvailability: []models.SeasonalAvailability{
				{
					Months: []int{6, 7, 8, 9},
					TimeRanges: []models.TimeRange{
						{Start: "16:00", End: "09:00"},
					},
				},
			},
			SouthAvailability: []models.SeasonalAvailability{
				{
					Months: []int{1, 2, 3, 12},
					TimeRanges: []models.TimeRange{
						{Start: "16:00", End: "09:00"},
					},
				},
			},
		},
		{
			FishID:     "45-saddled-bichir",
			Name:       "Saddled Bichir",
			Icon:       "/static/images/fish/icons/saddled-bichir.png",
			SellPrice:  4000,
			ShadowSize: "Large",
			ShadowIcon: "/static/images/fish/icons/shadow-large.png",
			Location:   "River",
			NorthAvailability: []models.SeasonalAvailability{
				{
					Months: []int{6, 7, 8, 9},
					TimeRanges: []models.TimeRange{
						{Start: "21:00", End: "04:00"},
					},
				},
			},
			SouthAvailability: []models.SeasonalAvailability{
				{
					Months: []int{1, 2, 3, 12},
					TimeRanges: []models.TimeRange{
						{Start: "21:00", End: "04:00"},
					},
				},
			},
		},
		{
			FishID:     "46-sturgeon",
			Name:       "Sturgeon",
			Icon:       "/static/images/fish/icons/sturgeon.png",
			SellPrice:  10000,
			ShadowSize: "Huge",
			ShadowIcon: "/static/images/fish/icons/shadow-huge.png",
			Location:   "River (Mouth)",
			NorthAvailability: []models.SeasonalAvailability{
				{
					Months: []int{1, 2, 3, 9, 10, 11, 12},
					TimeRanges: []models.TimeRange{
						{Start: "00:00", End: "23:59"},
					},
				},
			},
			SouthAvailability: []models.SeasonalAvailability{
				{
					Months: []int{3, 4, 5, 6, 7, 8, 9},
					TimeRanges: []models.TimeRange{
						{Start: "00:00", End: "23:59"},
					},
				},
			},
		},
		{
			FishID:     "47-sea-butterfly",
			Name:       "Sea Butterfly",
			Icon:       "/static/images/fish/icons/sea-butterfly.png",
			SellPrice:  1000,
			ShadowSize: "Tiny",
			ShadowIcon: "/static/images/fish/icons/shadow-tiny.png",
			Location:   "Sea",
			NorthAvailability: []models.SeasonalAvailability{
				{
					Months: []int{1, 2, 3, 12},
					TimeRanges: []models.TimeRange{
						{Start: "00:00", End: "23:59"},
					},
				},
			},
			SouthAvailability: []models.SeasonalAvailability{
				{
					Months: []int{6, 7, 8, 9},
					TimeRanges: []models.TimeRange{
						{Start: "00:00", End: "23:59"},
					},
				},
			},
		},
		{
			FishID:     "48-sea-horse",
			Name:       "Sea Horse",
			Icon:       "/static/images/fish/icons/sea-horse.png",
			SellPrice:  1100,
			ShadowSize: "Tiny",
			ShadowIcon: "/static/images/fish/icons/shadow-tiny.png",
			Location:   "Sea",
			NorthAvailability: []models.SeasonalAvailability{
				{
					Months: []int{4, 5, 6, 7, 8, 9, 10, 11},
					TimeRanges: []models.TimeRange{
						{Start: "00:00", End: "23:59"},
					},
				},
			},
			SouthAvailability: []models.SeasonalAvailability{
				{
					Months: []int{1, 2, 3, 4, 5, 10, 11, 12},
					TimeRanges: []models.TimeRange{
						{Start: "00:00", End: "23:59"},
					},
				},
			},
		},
		{
			FishID:     "49-clown-fish",
			Name:       "Clown Fish",
			Icon:       "/static/images/fish/icons/clown-fish.png",
			SellPrice:  650,
			ShadowSize: "Tiny",
			ShadowIcon: "/static/images/fish/icons/shadow-tiny.png",
			Location:   "Sea",
			NorthAvailability: []models.SeasonalAvailability{
				{
					Months: []int{4, 5, 6, 7, 8, 9},
					TimeRanges: []models.TimeRange{
						{Start: "00:00", End: "23:59"},
					},
				},
			},
			SouthAvailability: []models.SeasonalAvailability{
				{
					Months: []int{1, 2, 3, 10, 11, 12},
					TimeRanges: []models.TimeRange{
						{Start: "00:00", End: "23:59"},
					},
				},
			},
		},
		{
			FishID:     "50-surgeonfish",
			Name:       "Surgeonfish",
			Icon:       "/static/images/fish/icons/surgeonfish.png",
			SellPrice:  1000,
			ShadowSize: "Small",
			ShadowIcon: "/static/images/fish/icons/shadow-small.png",
			Location:   "Sea",
			NorthAvailability: []models.SeasonalAvailability{
				{
					Months: []int{4, 5, 6, 7, 8, 9},
					TimeRanges: []models.TimeRange{
						{Start: "00:00", End: "23:59"},
					},
				},
			},
			SouthAvailability: []models.SeasonalAvailability{
				{
					Months: []int{1, 2, 3, 10, 11, 12},
					TimeRanges: []models.TimeRange{
						{Start: "00:00", End: "23:59"},
					},
				},
			},
		},
		{
			FishID:     "51-butterfly-fish",
			Name:       "Butterfly Fish",
			Icon:       "/static/images/fish/icons/butterfly-fish.png",
			SellPrice:  1000,
			ShadowSize: "Small",
			ShadowIcon: "/static/images/fish/icons/shadow-small.png",
			Location:   "Sea",
			NorthAvailability: []models.SeasonalAvailability{
				{
					Months: []int{4, 5, 6, 7, 8, 9},
					TimeRanges: []models.TimeRange{
						{Start: "00:00", End: "23:59"},
					},
				},
			},
			SouthAvailability: []models.SeasonalAvailability{
				{
					Months: []int{1, 2, 3, 10, 11, 12},
					TimeRanges: []models.TimeRange{
						{Start: "00:00", End: "23:59"},
					},
				},
			},
		},
		{
			FishID:     "52-napoleonfish",
			Name:       "Napoleonfish",
			Icon:       "/static/images/fish/icons/napoleonfish.png",
			SellPrice:  10000,
			ShadowSize: "Huge",
			ShadowIcon: "/static/images/fish/icons/shadow-huge.png",
			Location:   "Sea",
			NorthAvailability: []models.SeasonalAvailability{
				{
					Months: []int{7, 8},
					TimeRanges: []models.TimeRange{
						{Start: "04:00", End: "21:00"},
					},
				},
			},
			SouthAvailability: []models.SeasonalAvailability{
				{
					Months: []int{1, 2},
					TimeRanges: []models.TimeRange{
						{Start: "04:00", End: "21:00"},
					},
				},
			},
		},
		{
			FishID:     "53-zebra-turkeyfish",
			Name:       "Zebra Turkeyfish",
			Icon:       "/static/images/fish/icons/zebra-turkeyfish.png",
			SellPrice:  500,
			ShadowSize: "Medium",
			ShadowIcon: "/static/images/fish/icons/shadow-medium.png",
			Location:   "Sea",
			NorthAvailability: []models.SeasonalAvailability{
				{
					Months: []int{4, 5, 6, 7, 8, 9, 10, 11},
					TimeRanges: []models.TimeRange{
						{Start: "00:00", End: "23:59"},
					},
				},
			},
			SouthAvailability: []models.SeasonalAvailability{
				{
					Months: []int{1, 2, 3, 4, 5, 10, 11, 12},
					TimeRanges: []models.TimeRange{
						{Start: "00:00", End: "23:59"},
					},
				},
			},
		},
		{
			FishID:     "54-blowfish",
			Name:       "Blowfish",
			Icon:       "/static/images/fish/icons/blowfish.png",
			SellPrice:  5000,
			ShadowSize: "Medium",
			ShadowIcon: "/static/images/fish/icons/shadow-medium.png",
			Location:   "Sea",
			NorthAvailability: []models.SeasonalAvailability{
				{
					Months: []int{1, 2, 11, 12},
					TimeRanges: []models.TimeRange{
						{Start: "21:00", End: "04:00"},
					},
				},
			},
			SouthAvailability: []models.SeasonalAvailability{
				{
					Months: []int{5, 6, 7, 8},
					TimeRanges: []models.TimeRange{
						{Start: "21:00", End: "04:00"},
					},
				},
			},
		},
		{
			FishID:     "55-puffer-fish",
			Name:       "Puffer Fish",
			Icon:       "/static/images/fish/icons/puffer-fish.png",
			SellPrice:  250,
			ShadowSize: "Medium",
			ShadowIcon: "/static/images/fish/icons/shadow-medium.png",
			Location:   "Sea",
			NorthAvailability: []models.SeasonalAvailability{
				{
					Months: []int{7, 8, 9},
					TimeRanges: []models.TimeRange{
						{Start: "00:00", End: "23:59"},
					},
				},
			},
			SouthAvailability: []models.SeasonalAvailability{
				{
					Months: []int{1, 2, 3},
					TimeRanges: []models.TimeRange{
						{Start: "00:00", End: "23:59"},
					},
				},
			},
		},
		{
			FishID:     "56-anchovy",
			Name:       "Anchovy",
			Icon:       "/static/images/fish/icons/anchovy.png",
			SellPrice:  200,
			ShadowSize: "Small",
			ShadowIcon: "/static/images/fish/icons/shadow-small.png",
			Location:   "Sea",
			NorthAvailability: []models.SeasonalAvailability{
				{
					Months: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12},
					TimeRanges: []models.TimeRange{
						{Start: "04:00", End: "21:00"},
					},
				},
			},
			SouthAvailability: []models.SeasonalAvailability{
				{
					Months: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12},
					TimeRanges: []models.TimeRange{
						{Start: "04:00", End: "21:00"},
					},
				},
			},
		},
		{
			FishID:     "57-horse-mackerel",
			Name:       "Horse Mackerel",
			Icon:       "/static/images/fish/icons/horse-mackerel.png",
			SellPrice:  150,
			ShadowSize: "Small",
			ShadowIcon: "/static/images/fish/icons/shadow-small.png",
			Location:   "Sea",
			NorthAvailability: []models.SeasonalAvailability{
				{
					Months: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12},
					TimeRanges: []models.TimeRange{
						{Start: "00:00", End: "23:59"},
					},
				},
			},
			SouthAvailability: []models.SeasonalAvailability{
				{
					Months: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12},
					TimeRanges: []models.TimeRange{
						{Start: "00:00", End: "23:59"},
					},
				},
			},
		},
		{
			FishID:     "58-barred-knifejaw",
			Name:       "Barred Knifejaw",
			Icon:       "/static/images/fish/icons/barred-knifejaw.png",
			SellPrice:  5000,
			ShadowSize: "Medium",
			ShadowIcon: "/static/images/fish/icons/shadow-medium.png",
			Location:   "Sea",
			NorthAvailability: []models.SeasonalAvailability{
				{
					Months: []int{3, 4, 5, 6, 7, 8, 9, 10, 11},
					TimeRanges: []models.TimeRange{
						{Start: "00:00", End: "23:59"},
					},
				},
			},
			SouthAvailability: []models.SeasonalAvailability{
				{
					Months: []int{1, 2, 3, 4, 5, 9, 10, 11, 12},
					TimeRanges: []models.TimeRange{
						{Start: "00:00", End: "23:59"},
					},
				},
			},
		},
		{
			FishID:     "59-sea-bass",
			Name:       "Sea Bass",
			Icon:       "/static/images/fish/icons/sea-bass.png",
			SellPrice:  400,
			ShadowSize: "Very Large",
			ShadowIcon: "/static/images/fish/icons/shadow-very-large.png",
			Location:   "Sea",
			NorthAvailability: []models.SeasonalAvailability{
				{
					Months: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12},
					TimeRanges: []models.TimeRange{
						{Start: "00:00", End: "23:59"},
					},
				},
			},
			SouthAvailability: []models.SeasonalAvailability{
				{
					Months: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12},
					TimeRanges: []models.TimeRange{
						{Start: "00:00", End: "23:59"},
					},
				},
			},
		},
		{
			FishID:     "60-red-snapper",
			Name:       "Red Snapper",
			Icon:       "/static/images/fish/icons/red-snapper.png",
			SellPrice:  3000,
			ShadowSize: "Large",
			ShadowIcon: "/static/images/fish/icons/shadow-large.png",
			Location:   "Sea",
			NorthAvailability: []models.SeasonalAvailability{
				{
					Months: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12},
					TimeRanges: []models.TimeRange{
						{Start: "00:00", End: "23:59"},
					},
				},
			},
			SouthAvailability: []models.SeasonalAvailability{
				{
					Months: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12},
					TimeRanges: []models.TimeRange{
						{Start: "00:00", End: "23:59"},
					},
				},
			},
		},
		{
			FishID:     "61-dab",
			Name:       "Dab",
			Icon:       "/static/images/fish/icons/dab.png",
			SellPrice:  300,
			ShadowSize: "Medium",
			ShadowIcon: "/static/images/fish/icons/shadow-medium.png",
			Location:   "Sea",
			NorthAvailability: []models.SeasonalAvailability{
				{
					Months: []int{1, 2, 3, 4, 10, 11, 12},
					TimeRanges: []models.TimeRange{
						{Start: "00:00", End: "23:59"},
					},
				},
			},
			SouthAvailability: []models.SeasonalAvailability{
				{
					Months: []int{4, 5, 6, 7, 8, 9, 10},
					TimeRanges: []models.TimeRange{
						{Start: "00:00", End: "23:59"},
					},
				},
			},
		},
		{
			FishID:     "62-olive-flounder",
			Name:       "Olive Flounder",
			Icon:       "/static/images/fish/icons/olive-flounder.png",
			SellPrice:  800,
			ShadowSize: "Very Large",
			ShadowIcon: "/static/images/fish/icons/shadow-very-large.png",
			Location:   "Sea",
			NorthAvailability: []models.SeasonalAvailability{
				{
					Months: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12},
					TimeRanges: []models.TimeRange{
						{Start: "00:00", End: "23:59"},
					},
				},
			},
			SouthAvailability: []models.SeasonalAvailability{
				{
					Months: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12},
					TimeRanges: []models.TimeRange{
						{Start: "00:00", End: "23:59"},
					},
				},
			},
		},
		{
			FishID:     "63-squid",
			Name:       "Squid",
			Icon:       "/static/images/fish/icons/squid.png",
			SellPrice:  500,
			ShadowSize: "Medium",
			ShadowIcon: "/static/images/fish/icons/shadow-medium.png",
			Location:   "Sea",
			NorthAvailability: []models.SeasonalAvailability{
				{
					Months: []int{1, 2, 3, 4, 5, 6, 7, 8, 12},
					TimeRanges: []models.TimeRange{
						{Start: "00:00", End: "23:59"},
					},
				},
			},
			SouthAvailability: []models.SeasonalAvailability{
				{
					Months: []int{1, 2, 6, 7, 8, 9, 10, 11, 12},
					TimeRanges: []models.TimeRange{
						{Start: "00:00", End: "23:59"},
					},
				},
			},
		},
		{
			FishID:     "64-moray-eel",
			Name:       "Moray Eel",
			Icon:       "/static/images/fish/icons/moray-eel.png",
			SellPrice:  2000,
			ShadowSize: "Long and Thin",
			ShadowIcon: "/static/images/fish/icons/shadow-long-and-thin.png",
			Location:   "Sea",
			NorthAvailability: []models.SeasonalAvailability{
				{
					Months: []int{8, 9, 10},
					TimeRanges: []models.TimeRange{
						{Start: "00:00", End: "23:59"},
					},
				},
			},
			SouthAvailability: []models.SeasonalAvailability{
				{
					Months: []int{2, 3, 4},
					TimeRanges: []models.TimeRange{
						{Start: "00:00", End: "23:59"},
					},
				},
			},
		},
		{
			FishID:     "65-ribbon-eel",
			Name:       "Ribbon Eel",
			Icon:       "/static/images/fish/icons/ribbon-eel.png",
			SellPrice:  600,
			ShadowSize: "Long and Thin",
			ShadowIcon: "/static/images/fish/icons/shadow-long-and-thin.png",
			Location:   "Sea",
			NorthAvailability: []models.SeasonalAvailability{
				{
					Months: []int{6, 7, 8, 9, 10},
					TimeRanges: []models.TimeRange{
						{Start: "00:00", End: "23:59"},
					},
				},
			},
			SouthAvailability: []models.SeasonalAvailability{
				{
					Months: []int{1, 2, 3, 4, 12},
					TimeRanges: []models.TimeRange{
						{Start: "00:00", End: "23:59"},
					},
				},
			},
		},
		{
			FishID:     "66-tuna",
			Name:       "Tuna",
			Icon:       "/static/images/fish/icons/tuna.png",
			SellPrice:  7000,
			ShadowSize: "Huge",
			ShadowIcon: "/static/images/fish/icons/shadow-huge.png",
			Location:   "Pier",
			NorthAvailability: []models.SeasonalAvailability{
				{
					Months: []int{1, 2, 3, 4, 11, 12},
					TimeRanges: []models.TimeRange{
						{Start: "00:00", End: "23:59"},
					},
				},
			},
			SouthAvailability: []models.SeasonalAvailability{
				{
					Months: []int{5, 6, 7, 8, 9, 10},
					TimeRanges: []models.TimeRange{
						{Start: "00:00", End: "23:59"},
					},
				},
			},
		},
		{
			FishID:     "67-blue-marlin",
			Name:       "Blue Marlin",
			Icon:       "/static/images/fish/icons/blue-marlin.png",
			SellPrice:  10000,
			ShadowSize: "Huge",
			ShadowIcon: "/static/images/fish/icons/shadow-huge.png",
			Location:   "Pier",
			NorthAvailability: []models.SeasonalAvailability{
				{
					Months: []int{1, 2, 3, 4, 7, 8, 9, 11, 12},
					TimeRanges: []models.TimeRange{
						{Start: "00:00", End: "23:59"},
					},
				},
			},
			SouthAvailability: []models.SeasonalAvailability{
				{
					Months: []int{1, 2, 3, 5, 6, 7, 8, 9, 10},
					TimeRanges: []models.TimeRange{
						{Start: "00:00", End: "23:59"},
					},
				},
			},
		},
		{
			FishID:     "68-giant-trevally",
			Name:       "Giant Trevally",
			Icon:       "/static/images/fish/icons/giant-trevally.png",
			SellPrice:  4500,
			ShadowSize: "Very Large",
			ShadowIcon: "/static/images/fish/icons/shadow-very-large.png",
			Location:   "Pier",
			NorthAvailability: []models.SeasonalAvailability{
				{
					Months: []int{5, 6, 7, 8, 9, 10},
					TimeRanges: []models.TimeRange{
						{Start: "00:00", End: "23:59"},
					},
				},
			},
			SouthAvailability: []models.SeasonalAvailability{
				{
					Months: []int{1, 2, 3, 4, 11, 12},
					TimeRanges: []models.TimeRange{
						{Start: "00:00", End: "23:59"},
					},
				},
			},
		},
		{
			FishID:     "69-mahi-mahi",
			Name:       "Mahi-mahi",
			Icon:       "/static/images/fish/icons/mahi-mahi.png",
			SellPrice:  6000,
			ShadowSize: "Very Large",
			ShadowIcon: "/static/images/fish/icons/shadow-very-large.png",
			Location:   "Pier",
			NorthAvailability: []models.SeasonalAvailability{
				{
					Months: []int{5, 6, 7, 8, 9, 10},
					TimeRanges: []models.TimeRange{
						{Start: "00:00", End: "23:59"},
					},
				},
			},
			SouthAvailability: []models.SeasonalAvailability{
				{
					Months: []int{1, 2, 3, 4, 11, 12},
					TimeRanges: []models.TimeRange{
						{Start: "00:00", End: "23:59"},
					},
				},
			},
		},
		{
			FishID:     "70-ocean-sunfish",
			Name:       "Ocean Sunfish",
			Icon:       "/static/images/fish/icons/ocean-sunfish.png",
			SellPrice:  4000,
			ShadowSize: "Very Large (Finned)",
			ShadowIcon: "/static/images/fish/icons/shadow-very-large.png",
			Location:   "Sea",
			NorthAvailability: []models.SeasonalAvailability{
				{
					Months: []int{7, 8, 9},
					TimeRanges: []models.TimeRange{
						{Start: "04:00", End: "21:00"},
					},
				},
			},
			SouthAvailability: []models.SeasonalAvailability{
				{
					Months: []int{1, 2, 3},
					TimeRanges: []models.TimeRange{
						{Start: "04:00", End: "21:00"},
					},
				},
			},
		},
		{
			FishID:     "71-ray",
			Name:       "Ray",
			Icon:       "/static/images/fish/icons/ray.png",
			SellPrice:  3000,
			ShadowSize: "Very Large",
			ShadowIcon: "/static/images/fish/icons/shadow-very-large.png",
			Location:   "Sea",
			NorthAvailability: []models.SeasonalAvailability{
				{
					Months: []int{8, 9, 10, 11},
					TimeRanges: []models.TimeRange{
						{Start: "04:00", End: "21:00"},
					},
				},
			},
			SouthAvailability: []models.SeasonalAvailability{
				{
					Months: []int{2, 3, 4, 5},
					TimeRanges: []models.TimeRange{
						{Start: "04:00", End: "21:00"},
					},
				},
			},
		},
		{
			FishID:     "72-saw-shark",
			Name:       "Saw Shark",
			Icon:       "/static/images/fish/icons/saw-shark.png",
			SellPrice:  12000,
			ShadowSize: "very Large (Finned)",
			ShadowIcon: "/static/images/fish/icons/shadow-very-large.png",
			Location:   "Sea",
			NorthAvailability: []models.SeasonalAvailability{
				{
					Months: []int{6, 7, 8, 9},
					TimeRanges: []models.TimeRange{
						{Start: "16:00", End: "09:00"},
					},
				},
			},
			SouthAvailability: []models.SeasonalAvailability{
				{
					Months: []int{1, 2, 3, 12},
					TimeRanges: []models.TimeRange{
						{Start: "16:00", End: "09:00"},
					},
				},
			},
		},
		{
			FishID:     "73-hammerhead-shark",
			Name:       "Hammerhead Shark",
			Icon:       "/static/images/fish/icons/hammerhead-shark.png",
			SellPrice:  8000,
			ShadowSize: "Very Large (Finned)",
			ShadowIcon: "/static/images/fish/icons/shadow-very-large.png",
			Location:   "Sea",
			NorthAvailability: []models.SeasonalAvailability{
				{
					Months: []int{6, 7, 8, 9},
					TimeRanges: []models.TimeRange{
						{Start: "16:00", End: "09:00"},
					},
				},
			},
			SouthAvailability: []models.SeasonalAvailability{
				{
					Months: []int{1, 2, 3, 12},
					TimeRanges: []models.TimeRange{
						{Start: "16:00", End: "09:00"},
					},
				},
			},
		},
		{
			FishID:     "74-great-white-shark",
			Name:       "Great White Shark",
			Icon:       "/static/images/fish/icons/great-white-shark.png",
			SellPrice:  15000,
			ShadowSize: "Very Large (Finned)",
			ShadowIcon: "/static/images/fish/icons/shadow-very-large.png",
			Location:   "Sea",
			NorthAvailability: []models.SeasonalAvailability{
				{
					Months: []int{6, 7, 8, 9},
					TimeRanges: []models.TimeRange{
						{Start: "16:00", End: "09:00"},
					},
				},
			},
			SouthAvailability: []models.SeasonalAvailability{
				{
					Months: []int{1, 2, 3, 12},
					TimeRanges: []models.TimeRange{
						{Start: "16:00", End: "09:00"},
					},
				},
			},
		},
		{
			FishID:     "75-whale-shark",
			Name:       "Whale Shark",
			Icon:       "/static/images/fish/icons/whale-shark.png",
			SellPrice:  13000,
			ShadowSize: "Very Large (Finned)",
			ShadowIcon: "/static/images/fish/icons/shadow-very-large.png",
			Location:   "Sea",
			NorthAvailability: []models.SeasonalAvailability{
				{
					Months: []int{6, 7, 8, 9},
					TimeRanges: []models.TimeRange{
						{Start: "00:00", End: "23:59"},
					},
				},
			},
			SouthAvailability: []models.SeasonalAvailability{
				{
					Months: []int{1, 2, 3, 12},
					TimeRanges: []models.TimeRange{
						{Start: "00:00", End: "23:59"},
					},
				},
			},
		},
		{
			FishID:     "76-suckerfish",
			Name:       "Suckerfish",
			Icon:       "/static/images/fish/icons/suckerfish.png",
			SellPrice:  1500,
			ShadowSize: "Very Large (Finned)",
			ShadowIcon: "/static/images/fish/icons/shadow-very-large.png",
			Location:   "Sea",
			NorthAvailability: []models.SeasonalAvailability{
				{
					Months: []int{6, 7, 8, 9},
					TimeRanges: []models.TimeRange{
						{Start: "00:00", End: "23:59"},
					},
				},
			},
			SouthAvailability: []models.SeasonalAvailability{
				{
					Months: []int{1, 2, 3, 12},
					TimeRanges: []models.TimeRange{
						{Start: "00:00", End: "23:59"},
					},
				},
			},
		},
		{
			FishID:     "77-football-fish",
			Name:       "Football Fish",
			Icon:       "/static/images/fish/icons/football-fish.png",
			SellPrice:  2500,
			ShadowSize: "Large",
			ShadowIcon: "/static/images/fish/icons/shadow-large.png",
			Location:   "Sea",
			NorthAvailability: []models.SeasonalAvailability{
				{
					Months: []int{1, 2, 3, 11, 12},
					TimeRanges: []models.TimeRange{
						{Start: "16:00", End: "09:00"},
					},
				},
			},
			SouthAvailability: []models.SeasonalAvailability{
				{
					Months: []int{5, 6, 7, 8, 9},
					TimeRanges: []models.TimeRange{
						{Start: "16:00", End: "09:00"},
					},
				},
			},
		},
		{
			FishID:     "78-oarfish",
			Name:       "Oarfish",
			Icon:       "/static/images/fish/icons/oarfish.png",
			SellPrice:  9000,
			ShadowSize: "Huge",
			ShadowIcon: "/static/images/fish/icons/shadow-huge.png",
			Location:   "Sea",
			NorthAvailability: []models.SeasonalAvailability{
				{
					Months: []int{1, 2, 3, 4, 5, 12},
					TimeRanges: []models.TimeRange{
						{Start: "00:00", End: "23:59"},
					},
				},
			},
			SouthAvailability: []models.SeasonalAvailability{
				{
					Months: []int{6, 7, 8, 9, 10, 11},
					TimeRanges: []models.TimeRange{
						{Start: "00:00", End: "23:59"},
					},
				},
			},
		},
		{
			FishID:     "79-barreleye",
			Name:       "Barreleye",
			Icon:       "/static/images/fish/icons/barreleye.png",
			SellPrice:  15000,
			ShadowSize: "Small",
			ShadowIcon: "/static/images/fish/icons/shadow-small.png",
			Location:   "Sea",
			NorthAvailability: []models.SeasonalAvailability{
				{
					Months: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12},
					TimeRanges: []models.TimeRange{
						{Start: "21:00", End: "04:00"},
					},
				},
			},
			SouthAvailability: []models.SeasonalAvailability{
				{
					Months: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12},
					TimeRanges: []models.TimeRange{
						{Start: "21:00", End: "04:00"},
					},
				},
			},
		},
		{
			FishID:     "80-coelacanth",
			Name:       "Coelacanth",
			Icon:       "/static/images/fish/icons/coelacanth.png",
			SellPrice:  15000,
			ShadowSize: "Huge",
			ShadowIcon: "/static/images/fish/icons/shadow-huge.png",
			Location:   "Sea (Raining)",
			NorthAvailability: []models.SeasonalAvailability{
				{
					Months: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12},
					TimeRanges: []models.TimeRange{
						{Start: "00:00", End: "23:59"},
					},
				},
			},
			SouthAvailability: []models.SeasonalAvailability{
				{
					Months: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12},
					TimeRanges: []models.TimeRange{
						{Start: "00:00", End: "23:59"},
					},
				},
			},
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
