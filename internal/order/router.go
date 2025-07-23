package order

import (
	"database/sql"
	"net/http"
	
	"github.com/aws/aws-lambda-go/events"
	"github.com/lubualo/ecommerce-go/tools"
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
	return r.handler.Post(requestWithContext)
}

func (r *Router) Get(requestWithContext models.RequestWithContext) *events.APIGatewayProxyResponse {
	return r.handler.Get(requestWithContext)

}

func (r *Router) Put(requestWithContext models.RequestWithContext) *events.APIGatewayProxyResponse {
	return tools.CreateApiResponse(http.StatusMethodNotAllowed, "not implemented")

}

func (r *Router) Delete(requestWithContext models.RequestWithContext) *events.APIGatewayProxyResponse {
	return tools.CreateApiResponse(http.StatusMethodNotAllowed, "not implemented")
}