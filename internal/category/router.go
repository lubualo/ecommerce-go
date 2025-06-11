package category

import (
	"database/sql"

	"github.com/aws/aws-lambda-go/events"
)

// Router struct contains all dependencies
type Router struct {
	handler *Handler
}

// NewCategoryRouter sets up the repository, service, and handler
func NewCategoryRouter(db *sql.DB) *Router {
	repo := NewSQLRepository(db)
	service := NewCategoryService(repo)
	handler := NewCategoryHandler(service)
	return &Router{handler: handler}
}

// Implements the EntityRouter interface

func (r *Router) Post(body string, user string) *events.APIGatewayProxyResponse {
	resp, _ := r.handler.Post(body, user)
	return resp
}

// Future implementations (stubs for now)
func (r *Router) Get(user, id string, query map[string]string) *events.APIGatewayProxyResponse {
	return nil
}

func (r *Router) Put(body, user, id string) *events.APIGatewayProxyResponse {
	return nil
}

func (r *Router) Delete(user, id string) *events.APIGatewayProxyResponse {
	return nil
}
