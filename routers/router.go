package routers

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/aws/aws-lambda-go/events"
	"github.com/ddessilvestri/ecommerce-go/auth"
	"github.com/ddessilvestri/ecommerce-go/internal/category"
)

// HTTP method constants
const (
	GET    = "GET"
	POST   = "POST"
	PUT    = "PUT"
	DELETE = "DELETE"
)

// Router determines which entity router should handle the request.
func Router(request events.APIGatewayV2HTTPRequest, urlPrefix string, db *sql.DB) *events.APIGatewayProxyResponse {
	// Extract path & method
	path := strings.Replace(request.RawPath, urlPrefix, "", 1)
	method := request.RequestContext.HTTP.Method
	id := request.PathParameters["id"]
	body := request.Body
	header := request.Headers
	query := request.QueryStringParameters

	// Extract main segment (e.g. /category/123 => category)
	firstSegment := getFirstPathSegment(path)

	// Find the corresponding entity router (e.g., category.Router)
	entityRouter, err := CreateRouter(firstSegment, db)
	if err != nil {
		return jsonResponse(400, "Unable to route request: "+err.Error())
	}

	// Authenticate
	isOk, statusCode, user := auth.AuthValidation(path, method, header)
	if !isOk {
		return &events.APIGatewayProxyResponse{
			StatusCode: statusCode,
			Body:       user,
		}
	}

	// Route to correct handler based on HTTP method
	switch method {
	case GET:
		return entityRouter.Get(user, id, query)
	case POST:
		return entityRouter.Post(body, user)
	case PUT:
		return entityRouter.Put(body, user, id)
	case DELETE:
		return entityRouter.Delete(user, id)
	default:
		return jsonResponse(405, "Method not allowed")
	}
}

// CreateRouter maps entity names to their router implementations
func CreateRouter(entity string, db *sql.DB) (EntityRouter, error) {
	switch entity {
	case "category":
		return category.NewCategoryRouter(db), nil
	// case "product":
	//     return product.NewProductRouter(), nil
	default:
		return nil, fmt.Errorf("entity '%s' not implemented", entity)
	}
}

// Extract first part of path: "/category/1" -> "category"
func getFirstPathSegment(path string) string {
	path = strings.Trim(path, "/")
	parts := strings.Split(path, "/")
	if len(parts) > 0 {
		return parts[0]
	}
	return ""
}
