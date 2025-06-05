package main

import (
	"context"
	"os"
	"strings"
	"github.com/aws/aws-lambda-go/events"
	"github.com/lubualo/ecommerce-go/awsgo"
	"github.com/lubualo/ecommerce-go/db"
	"github.com/lubualo/ecommerce-go/handlers"

	lambda "github.com/aws/aws-lambda-go/lambda"
)

func main()  {
	lambda.Start(LambdaExec)
}

func LambdaExec(ctx context.Context, request events.APIGatewayV2HTTPRequest)  (*events.APIGatewayProxyResponse, error) {
	awsgo.AWSInit()
	if !IsValid() {
		panic("Missing param: 'SecretName', 'UserPoolId', 'Region' and 'UrlPrefix' are required")
	}
	var response *events.APIGatewayProxyResponse
	path := strings.Replace(request.RawPath, os.Getenv("UrlPrefix"), "", -1)
	method := request.RequestContext.HTTP.Method
	body := request.Body
	headers := request.Headers
	
	db.ReadSecret()

	status, message := handlers.Handlers(path, method, body, headers, request)

	responseHeaders := map[string]string {
		"Content-Type": "application/json",
	}
	response = &events.APIGatewayProxyResponse{
		StatusCode: status,
		Body: message,
		Headers: responseHeaders,
	}

	return response, nil
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
