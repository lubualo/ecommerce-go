package routers

import (
	"encoding/json"

	"github.com/aws/aws-lambda-go/events"
)

func jsonResponse(status int, message string) *events.APIGatewayProxyResponse {
	body, _ := json.Marshal(map[string]string{
		"message": message,
	})
	return &events.APIGatewayProxyResponse{
		StatusCode: status,
		Body:       string(body),
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
	}
}
