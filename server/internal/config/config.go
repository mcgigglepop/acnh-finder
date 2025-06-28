package config

import (
	"html/template"
	"log"

	"github.com/alexedwards/scs/v2"
	"github.com/mcgigglepop/acnh-finder/server/internal/cognito"
	"github.com/mcgigglepop/acnh-finder/server/internal/dynamodb"
)

type DynamoService struct {
	UserProfile *dynamodb.DDBClient
	Fish        *dynamodb.DDBClient
	UserFish        *dynamodb.DDBClient
}

// AppConfig holds the application config
type AppConfig struct {
	UseCache      bool
	TemplateCache map[string]*template.Template
	InfoLog       *log.Logger
	ErrorLog      *log.Logger
	InProduction  bool
	Session       *scs.SessionManager
	CognitoClient *cognito.CognitoClient
	Dynamo        *DynamoService
}
