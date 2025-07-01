package routers

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"
	"strings"

	"github.com/aws/aws-lambda-go/events"
	"github.com/lubualo/ecommerce-go/auth"
	"github.com/lubualo/ecommerce-go/authctx"
	"github.com/lubualo/ecommerce-go/internal/category"
	"github.com/lubualo/ecommerce-go/internal/product"
	"github.com/lubualo/ecommerce-go/internal/stock"
	"github.com/lubualo/ecommerce-go/internal/user"
	"github.com/lubualo/ecommerce-go/models"
	"github.com/lubualo/ecommerce-go/tools"
)

const (
	GET    = "GET"
	POST   = "POST"
	PUT    = "PUT"
	DELETE = "DELETE"
)

func Router(request events.APIGatewayV2HTTPRequest, urlPrefix string, db *sql.DB) *events.APIGatewayProxyResponse {
	path := strings.Replace(request.RawPath, urlPrefix, "", 1)
	method := request.RequestContext.HTTP.Method
	header := request.Headers

	firstSegment := getFirstPathSegment(path)

	entityRouter, err := CreateRouter(firstSegment, db)
	if err != nil {
		return tools.CreateApiResponse(http.StatusBadRequest, "Unable to route request: "+err.Error())
	}

	var authUser *models.AuthUser

	if !(path == "product" && method == "GET") && !(path == "category" && method == "GET") {
		authUser, err = auth.ExtractAuthUser(header)
		if err != nil {
			return tools.CreateApiResponse(http.StatusUnauthorized, "Unable to authenticate user: "+err.Error())
		}
	}

	context := authctx.WithUser(context.Background(), authUser)
    requestWithContext := models.NewRequestWithContext(request, context)

	switch method {
	case GET:
		return entityRouter.Get(requestWithContext)
	case POST:
		return entityRouter.Post(requestWithContext)
	case PUT:
		return entityRouter.Put(requestWithContext)
	case DELETE:
		return entityRouter.Delete(requestWithContext)
	default:
		return tools.CreateApiResponse(http.StatusMethodNotAllowed, "Method not allowed")
	}
}

func CreateRouter(entity string, db *sql.DB) (EntityRouter, error) {
	switch entity {
	case "category":
		return category.NewRouter(db), nil
	case "product":
		return product.NewRouter(db), nil
	case "stock":
		return stock.NewRouter(db), nil
	case "user":
		return user.NewRouter(db), nil
	default:
		return nil, fmt.Errorf("entity '%s' not implemented", entity)
	}
}

func getFirstPathSegment(path string) string {
	path = strings.Trim(path, "/")
	parts := strings.Split(path, "/")
	if len(parts) > 0 {
		return parts[0]
	}
	return ""
}
