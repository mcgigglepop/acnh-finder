package models

type User struct {
	UserID     string `dynamodbav:"user_id"`
	Hemisphere string `dynamodbav:"hemisphere"`
}
