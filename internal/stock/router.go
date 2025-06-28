package stock

import (
	"database/sql"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/lubualo/ecommerce-go/tools"
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
	return tools.CreateApiResponse(http.StatusMethodNotAllowed, "not implemented")
}

func (r *Router) Get(request events.APIGatewayV2HTTPRequest) *events.APIGatewayProxyResponse {
	return tools.CreateApiResponse(http.StatusMethodNotAllowed, "not implemented")
}

func (r *Router) Put(request events.APIGatewayV2HTTPRequest) *events.APIGatewayProxyResponse {
	return r.handler.Put(request)
}

func (r *Router) Delete(request events.APIGatewayV2HTTPRequest) *events.APIGatewayProxyResponse {
	return tools.CreateApiResponse(http.StatusMethodNotAllowed, "not implemented")
}
