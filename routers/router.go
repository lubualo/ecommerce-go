package routers

import (
	"fmt"
	"strings"

	"github.com/aws/aws-lambda-go/events"
	"github.com/lubualo/ecommerce-go/auth"
)

func Router(request events.APIGatewayV2HTTPRequest, urlPrefix string) (int, string) {
	path := strings.Replace(request.RawPath, urlPrefix, "", -1)
	method := request.RequestContext.HTTP.Method
	body := request.Body
	headers := request.Headers
	id := request.PathParameters["id"]
	fmt.Println("Processing " + path + " > " + method)

	isOk, statusCode, user := validateAuthorization(path, method, headers)

	if !isOk {
		return statusCode, user
	}

	firstSegment := getFirstPathSegment(path)
	entityRouter, err := CreateRouter(firstSegment)
	if err != nil {
		return 400, "unable to create router: " + err.Error()
	}
	switch method {
	case Get:
		return entityRouter.Get(user, id, request.QueryStringParameters)
	case Post:
		return entityRouter.Post(body, user)
	case Put:
		return entityRouter.Put(body, user, id)
	case Delete:
		return entityRouter.Delete(user, id)
	default:
		return 405, "method not allowed"
	}
}

func validateAuthorization(path string, method string, headers map[string]string) (bool, int, string) {
	if (path == "product" && method == "GET") || (path == "category" && method == "GET") {
		return true, 200, ""
	}

	token := headers["authorization"]
	if len(token) == 0 {
		return false, 401, "Token required"
	}

	isOk, msg, _ := auth.ValidateToken(token)

	if !isOk {
		fmt.Println("Error in token: " + msg)
		return false, 401, msg
	}
	fmt.Println("Token OK")
	return true, 200, msg
}

func getFirstPathSegment(path string) string {
	// Remove leading/trailing slashes
	trimmed := strings.Trim(path, "/")
	segments := strings.Split(trimmed, "/")
	if len(segments) > 0 && segments[0] != "" {
		return segments[0]
	}
	return ""
}
