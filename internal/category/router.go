package category

import (
	"database/sql"

	"github.com/aws/aws-lambda-go/events"
)

type Router struct {
	handler *Handler
}

func NewRouter(db *sql.DB) *Router {
	repo := NewSQLRepository(db)
	service := NewService(repo)
	handler := NewHandler(service)
	return &Router{handler: handler}
}

func (r *Router) Post(request events.APIGatewayV2HTTPRequest) *events.APIGatewayProxyResponse {
	resp, _ := r.handler.Post(request)
	return resp
}

func (r *Router) Get(request events.APIGatewayV2HTTPRequest) *events.APIGatewayProxyResponse {
	resp, _ := r.handler.Get(request)
	return resp
}

func (r *Router) Put(request events.APIGatewayV2HTTPRequest) *events.APIGatewayProxyResponse {
	resp, _ := r.handler.Put(request)
	return resp
}

func (r *Router) Delete(request events.APIGatewayV2HTTPRequest) *events.APIGatewayProxyResponse {
	resp, _ := r.handler.Delete(request)
	return resp
}
