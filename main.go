package main

import (
	"context"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/lubualo/ecommerce-go/awsgo"
	"github.com/lubualo/ecommerce-go/db"
	"github.com/lubualo/ecommerce-go/routers"

	lambda "github.com/aws/aws-lambda-go/lambda"
)

func main() {
	lambda.Start(LambdaExec)
}

func LambdaExec(ctx context.Context, request events.APIGatewayV2HTTPRequest) (*events.APIGatewayProxyResponse, error) {
	awsgo.AWSInit()
	if !IsValid() {
		panic("Missing param: 'SecretName', 'UserPoolId', 'Region' and 'UrlPrefix' are required")
	}
	// Repository
	// db.ReadSecret()
	httpResponse := routers.Router(request, os.Getenv("UrlPrefix"))
	return httpResponse, nil
}

func IsValid() bool {
	var hasParam bool
	_, hasParam = os.LookupEnv("SecretName")
	if !hasParam {
		return hasParam
	}
	_, hasParam = os.LookupEnv("UrlPrefix")
	return hasParam
}
