package handlers

import (
	"fmt"
	"strconv"

	"github.com/aws/aws-lambda-go/events"
	"github.com/ddessilvestri/ecommerce-go/auth"
	"github.com/ddessilvestri/ecommerce-go/routers"
)

func Handlers(path string, method string, body string, header map[string]string, request events.APIGatewayV2HTTPRequest) (int, string) {
	fmt.Println("Processing " + path + " > " + method)
	id := request.PathParameters["id"]
	idn, _ := strconv.Atoi(id)

	isOk, statusCode, user := authValidation(path, method, header)

	if !isOk {
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

	return 400, "Invalid Method"

}

func authValidation(path string, method string, header map[string]string) (bool, int, string) {
	if (path == "product" && method == "GET") ||
		(path == "category" && method == "GET") {
		return true, 200, ""
	}
	token := header["authorization"]
	if len(token) == 0 {
		return false, 401, "Required Token"
	}

	isOk, msg, err := auth.TokenValidation(token)
	if !isOk {
		if err != nil {
			fmt.Println("Token Error " + err.Error())
			return false, 401, err.Error()
		}
		fmt.Println("Token Error " + msg)
		return false, 401, msg
	}

	fmt.Println("Token OK")
	return true, 200, msg

}

func ProcessUser(body string, path string, method string, user string, id string, request events.APIGatewayV2HTTPRequest) (int, string) {
	return 400, "Invalid Method"
}
func ProcessProducts(body string, path string, method string, user string, id int, request events.APIGatewayV2HTTPRequest) (int, string) {
	return 400, "Invalid Method"
}
func ProcessCategory(body string, path string, method string, user string, id int, request events.APIGatewayV2HTTPRequest) (int, string) {

	switch method {
	case "POST":
		return routers.InsertCategory(body, user)
	}

	return 400, "Invalid Method"
}
func ProcessStock(body string, path string, method string, user string, id int, request events.APIGatewayV2HTTPRequest) (int, string) {
	return 400, "Invalid Method"
}
func ProcessAddress(body string, path string, method string, user string, id int, request events.APIGatewayV2HTTPRequest) (int, string) {
	return 400, "Invalid Method"
}
func ProcessOrder(body string, path string, method string, user string, id int, request events.APIGatewayV2HTTPRequest) (int, string) {
	return 400, "Invalid Method"
}
