package routers

import (
	"fmt"
	"strings"

	"github.com/aws/aws-lambda-go/events"
	"github.com/lubualo/ecommerce-go/auth"
)

func Router(path string, method string, body string, headers map[string]string, request events.APIGatewayV2HTTPRequest) (int, string) {
	fmt.Println("Processing " + path + " > " + method)
	id := request.PathParameters["id"]

	isOk, statusCode, user := validateAuthorization(path, method, headers)

	if !isOk {
		return statusCode, user
	}

	firstSegment := getFirstPathSegment(path)
	entityRouter, err := CreateRouter(firstSegment)
	if err != nil {
		return 400, "unable to create router: " + err.Error()
	}
	return entityRouter.Route(body, path, method, user, id, request)
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
