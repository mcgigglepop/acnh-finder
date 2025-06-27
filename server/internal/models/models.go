package models

type User struct {
	UserID     string `dynamodbav:"user_id"`
	Hemisphere string `dynamodbav:"hemisphere"`
}

type Fish struct {
	FishID            string                 `dynamodbav:"fish_id"`
	Name              string                 `dynamodbav:"name"`
	Icon              string                 `dynamodbav:"icon"`
	SellPrice         int                    `dynamodbav:"sell_price"`
	ShadowSize        string                 `dynamodbav:"shadow_size"`
	ShadowIcon        string                 `dynamodbav:"shadow_icon"`
	Location          string                 `dynamodbav:"location"`
	NorthAvailability []SeasonalAvailability `dynamodbav:"north_availability"`
	SouthAvailability []SeasonalAvailability `dynamodbav:"south_availability"`
}

type UserFish struct {
	UserID string `dynamodbav:"user_id"`
	FishID string `dynamodbav:"fish_id"`
	Owned  bool   `dynamodbav:"owned"`
}

type TimeRange struct {
	Start string `dynamodbav:"start"` // "16:00"
	End   string `dynamodbav:"end"`   // "09:00"
}

type SeasonalAvailability struct {
	Months     []int       `dynamodbav:"months"`      // e.g. [3, 4, 5, 6]
	TimeRanges []TimeRange `dynamodbav:"time_ranges"` // 1+ ranges
}
