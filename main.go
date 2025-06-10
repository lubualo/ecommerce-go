package main

import (
	"context"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"

	"github.com/ddessilvestri/ecommerce-go/awsgo"
	"github.com/ddessilvestri/ecommerce-go/db"
	"github.com/ddessilvestri/ecommerce-go/internal/config"
	"github.com/ddessilvestri/ecommerce-go/routers"
	"github.com/ddessilvestri/ecommerce-go/secretm"
)

func main() {
	lambda.Start(LambdaExec)
}

func LambdaExec(ctx context.Context, request events.APIGatewayV2HTTPRequest) (*events.APIGatewayProxyResponse, error) {
	// Initialize AWS SDK
	awsgo.AWSInit()

	// Check required environment variables
	if err := config.ValidateEnvVars("SecretName", "UrlPrefix"); err != nil {
		panic("Environment validation error: " + err.Error())
	}

	// Load secrets from AWS Secrets Manager
	secretName := os.Getenv("SecretName")
	secret, err := secretm.GetSecret(secretName)
	if err != nil {
		return &events.APIGatewayProxyResponse{
			StatusCode: 500,
			Body:       "Failed to read secret: " + err.Error(),
		}, nil
	}

	// Establish database connection using the retrieved secret
	sqlDB, err := db.DbConnectWithSecret(secret)
	if err != nil {
		return &events.APIGatewayProxyResponse{
			StatusCode: 500,
			Body:       "Database connection error: " + err.Error(),
		}, nil
	}
	defer sqlDB.Close()

	// Call the central router
	urlPrefix := os.Getenv("UrlPrefix")
	response := routers.Router(request, urlPrefix, sqlDB)

	return response, nil
}
