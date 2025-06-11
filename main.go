package main

import (
	"context"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"

	"github.com/lubualo/ecommerce-go/awsgo"
	"github.com/lubualo/ecommerce-go/db"
	"github.com/lubualo/ecommerce-go/internal/config"
	"github.com/lubualo/ecommerce-go/routers"
	"github.com/lubualo/ecommerce-go/secretm"
)

func main() {
	lambda.Start(LambdaExec)
}

func LambdaExec(ctx context.Context, request events.APIGatewayV2HTTPRequest) (*events.APIGatewayProxyResponse, error) {
	awsgo.AWSInit()

	conf, err := config.LoadConfig()
	if err != nil {
		panic("Config load failed: " + err.Error())
	}

	secret, err := secretm.GetSecret(conf.SecretName)
	if err != nil {
		return &events.APIGatewayProxyResponse{
			StatusCode: 500,
			Body:       "Failed to read secret: " + err.Error(),
		}, nil
	}

	sqlDB, err := db.DbConnectAndReturn(secret, conf.DBName)
	if err != nil {
		return &events.APIGatewayProxyResponse{
			StatusCode: 500,
			Body:       "Database connection error: " + err.Error(),
		}, nil
	}
	defer sqlDB.Close()

	response := routers.Router(request, conf.UrlPrefix, sqlDB)
	return response, nil
}
