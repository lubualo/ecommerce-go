package handlers

import (
	"fmt"
	// "strconv"

	"github.com/aws/aws-lambda-go/events"
)

func Handlers(path string, method string, body string, header map[string]string, request events.APIGatewayV2HTTPRequest) (int, string) {
	fmt.Println("Processing " + path + " > " + method)
	// id := request.PathParameters["id"]
	// idn, _ := strconv.aToi(id)

	return 400, "Invalid Method"

}
