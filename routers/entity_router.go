// routers/entity_router.go
package routers

import "github.com/aws/aws-lambda-go/events"

type EntityRouter interface {
	Get(user, id string, query map[string]string) *events.APIGatewayProxyResponse
	Post(body, user string) *events.APIGatewayProxyResponse
	Put(body, user, id string) *events.APIGatewayProxyResponse
	Delete(user, id string) *events.APIGatewayProxyResponse
}
