package category

import (
	"database/sql"

	"github.com/aws/aws-lambda-go/events"
	"github.com/lubualo/ecommerce-go/models"
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

func (r *Router) Post(requestWithContext models.RequestWithContext) *events.APIGatewayProxyResponse {
	resp, _ := r.handler.Post(requestWithContext)
	return resp
}

func (r *Router) Get(requestWithContext models.RequestWithContext) *events.APIGatewayProxyResponse {
	resp, _ := r.handler.Get(requestWithContext)
	return resp
}

func (r *Router) Put(requestWithContext models.RequestWithContext) *events.APIGatewayProxyResponse {
	resp, _ := r.handler.Put(requestWithContext)
	return resp
}

func (r *Router) Delete(requestWithContext models.RequestWithContext) *events.APIGatewayProxyResponse {
	resp, _ := r.handler.Delete(requestWithContext)
	return resp
}
