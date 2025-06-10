package main

import (
	"context"

	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/ddessilvestri/ecommerce-go/awsgo"
	"github.com/ddessilvestri/ecommerce-go/db"
	"github.com/ddessilvestri/ecommerce-go/internal/config"
	"github.com/ddessilvestri/ecommerce-go/routers"

	lambda "github.com/aws/aws-lambda-go/lambda"
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

	// Read secrets from AWS Secrets Manager
	err := db.ReadSecret()
	if err != nil {
		return &events.APIGatewayProxyResponse{
			StatusCode: 500,
			Body:       "Failed to read secret: " + err.Error(),
		}, nil
	}

	// Establish a database connection
	sqlDB, err := db.DbConnectAndReturn()
	if err != nil {
		return &events.APIGatewayProxyResponse{
			StatusCode: 500,
			Body:       "Database connection error: " + err.Error(),
		}, nil
	}
	defer sqlDB.Close()

	// Call the new router with dependencies injected
	response := routers.Router(request, os.Getenv("UrlPrefix"), sqlDB)

	return response, nil
}
