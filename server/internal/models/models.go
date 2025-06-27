package models

type User struct {
	UserID     string `dynamodbav:"user_id"`
	Hemisphere string `dynamodbav:"hemisphere"`
}

type Fish struct {
	Name        string   `dynamodbav:"name"`
	Icon        string   `dynamodbav:"icon"`
	SellPrice   int      `dynamodbav:"sell_price"`
	ShadowSize  string   `dynamodbav:"shadow_size"`
	ShadowIcon  string   `dynamodbav:"shadow_icon"`
	Location    string   `dynamodbav:"location"`
	TimeOfDay   []string `dynamodbav:"time_of_day"` // e.g. ["9AM", "4PM", "ALL_DAY"]
	MonthsNorth []int    `dynamodbav:"months_north"` // 1–12
	MonthsSouth []int    `dynamodbav:"months_south"` // 1–12
}

type UserFish struct {
	UserID string `dynamodbav:"user_id"`
	FishID string `dynamodbav:"fish_id"`
	Owned  bool   `dynamodbav:"owned"`
}
