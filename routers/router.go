package routers

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/aws/aws-lambda-go/events"
	"github.com/lubualo/ecommerce-go/auth"
	"github.com/lubualo/ecommerce-go/handlers"
)

func Router(path string, method string, body string, headers map[string]string, request events.APIGatewayV2HTTPRequest) (int, string) {
	fmt.Println("Processing " + path + " > " + method)
	id := request.PathParameters["id"]
	idn, _ := strconv.Atoi(id)

	isOk, statusCode, user := validateAuthorization(path, method, headers)

	if !isOk {
		return statusCode, user
	}

	firstSegment := getFirstPathSegment(path)
	fmt.Println("First path segment: " + firstSegment)
	switch firstSegment {
	case "user":
		return ProcessUser(body, path, method, user, id, request)
	case "product":
		return ProcessProducts(body, path, method, user, idn, request)
	case "stock":
		return ProcessStock(body, path, method, user, idn, request)
	case "address":
		return ProcessAddress(body, path, method, user, idn, request)
	case "category":
		return ProcessCategory(body, path, method, user, idn, request)
	case "order":
		return ProcessOrder(body, path, method, user, idn, request)
	}

	return 400, "Invalid segment: " + firstSegment
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

func ProcessUser(body string, path string, method string, user string, id string, request events.APIGatewayV2HTTPRequest) (int, string) {
	return 400, "Invalid method"
}

func ProcessProducts(body string, path string, method string, user string, id int, request events.APIGatewayV2HTTPRequest) (int, string) {
	return 400, "Invalid method"
}

func ProcessCategory(body string, path string, method string, user string, id int, request events.APIGatewayV2HTTPRequest) (int, string) {
	switch method {
	case "POST":
		return handlers.PostCategory(body, user)
	}
	return 400, "Invalid method"
}

func ProcessStock(body string, path string, method string, user string, id int, request events.APIGatewayV2HTTPRequest) (int, string) {
	return 400, "Invalid method"
}

func ProcessAddress(body string, path string, method string, user string, id int, request events.APIGatewayV2HTTPRequest) (int, string) {
	return 400, "Invalid method"
}

func ProcessOrder(body string, path string, method string, user string, id int, request events.APIGatewayV2HTTPRequest) (int, string) {
	return 400, "Invalid method"
}
