package category

import (
	"database/sql"

	"github.com/aws/aws-lambda-go/events"
)

type Router struct {
	handler *Handler
}

func NewCategoryRouter(db *sql.DB) *Router {
	repo := NewSQLRepository(db)
	service := NewCategoryService(repo)
	handler := NewCategoryHandler(service)
	return &Router{handler: handler}
}

func (r *Router) Post(request events.APIGatewayV2HTTPRequest) *events.APIGatewayProxyResponse {
	resp, _ := r.handler.Post(request)
	return resp
}

func (r *Router) Get(request events.APIGatewayV2HTTPRequest) *events.APIGatewayProxyResponse {
	return nil
}

func (r *Router) Put(request events.APIGatewayV2HTTPRequest) *events.APIGatewayProxyResponse {
	return nil
}

func (r *Router) Delete(request events.APIGatewayV2HTTPRequest) *events.APIGatewayProxyResponse {
	return nil
}
