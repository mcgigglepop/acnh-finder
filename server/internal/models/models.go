package models

type User struct {
	UserID     string `dynamodbav:"user_id"`
	Hemisphere string `dynamodbav:"hemisphere"`
}

type Fish struct {
	FishID      string      `dynamodbav:"fish_id"`
	Name        string      `dynamodbav:"name"`
	Icon        string      `dynamodbav:"icon"`
	SellPrice   int         `dynamodbav:"sell_price"`
	ShadowSize  string      `dynamodbav:"shadow_size"`
	ShadowIcon  string      `dynamodbav:"shadow_icon"`
	Location    string      `dynamodbav:"location"`
	TimeRanges  []TimeRange `dynamodbav:"time_ranges"`
	MonthsNorth []int       `dynamodbav:"months_north"`
	MonthsSouth []int       `dynamodbav:"months_south"`
}

type UserFish struct {
	UserID string `dynamodbav:"user_id"`
	FishID string `dynamodbav:"fish_id"`
	Owned  bool   `dynamodbav:"owned"`
}

type TimeRange struct {
	Start string `dynamodbav:"start"` // e.g. "16:00"
	End   string `dynamodbav:"end"`   // e.g. "09:00"
}
