package handlers

import (
	"fmt"
	"strconv"

	"github.com/aws/aws-lambda-go/events"
	"github.com/lubualo/ecommerce-go/auth"
	"github.com/lubualo/ecommerce-go/routers"
)

func Handlers(path string, method string, body string, headers map[string]string, request events.APIGatewayV2HTTPRequest) (int, string) {
	fmt.Println("Processing " + path + " > " + method)
	id := request.PathParameters["id"]
	idn, _ := strconv.Atoi(id)

	isOk, statusCode, user := validateAuthorization(path, method, headers)

	if (!isOk) {
		return statusCode, user
	}

	switch path[0:4] {
		case "user":
			return ProcessUser(body, path, method, user, id, request)
		case "prod":
			return ProcessProducts(body, path, method, user, idn, request)
		case "stoc":
			return ProcessStock(body, path, method, user, idn, request)
		case "addr":
			return ProcessAddress(body, path, method, user, idn, request)
		case "cate":
			return ProcessCategory(body, path, method, user, idn, request)
		case "orde":
			return ProcessOrder(body, path, method, user, idn, request)
	}

	return 400, "Invalid method"
}

func validateAuthorization (path string, method string, headers map[string]string) (bool, int, string) {
	if (path == "product" && method == "GET") || (path == "category" && method == "GET") {
		return true, 200, ""
	}

	token := headers["authorization"]
	if len(token) == 0 {
		return false, 401, "Token required"
	}

	isOk, msg, err := auth.ValidateToken(token)

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
	switch method {
		case "POST":
			return routers.InsertCategory(body, user)
	}
	return 400, "Invalid method"
}

func ProcessCategory(body string, path string, method string, user string, id int, request events.APIGatewayV2HTTPRequest) (int, string) {
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