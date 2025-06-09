package handlers

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/aws/aws-lambda-go/events"
	"github.com/lubualo/ecommerce-go/auth"
	"github.com/lubualo/ecommerce-go/db"
	"github.com/lubualo/ecommerce-go/models"
	"github.com/lubualo/ecommerce-go/services"
)


type CategoryHandler struct{
	service services.CategoryService
}

func (handler *CategoryHandler) Get(request events.APIGatewayV2HTTPRequest) (*events.APIGatewayProxyResponse) {
	responseHeaders := map[string]string{
		"Content-Type": "application/json",
	}
	var response *events.APIGatewayProxyResponse
	response = &events.APIGatewayProxyResponse{
		StatusCode: 405,
		Body:       "method not implemented",
		Headers:    responseHeaders,
	}
	return response
}

func (handler *CategoryHandler) Put(request events.APIGatewayV2HTTPRequest) (*events.APIGatewayProxyResponse) {
	responseHeaders := map[string]string{
		"Content-Type": "application/json",
	}
	var response *events.APIGatewayProxyResponse
	response = &events.APIGatewayProxyResponse{
		StatusCode: 405,
		Body:       "method not implemented",
		Headers:    responseHeaders,
	}
	return response

}

func (handler *CategoryHandler) Delete(request events.APIGatewayV2HTTPRequest) (*events.APIGatewayProxyResponse) {
	responseHeaders := map[string]string{
		"Content-Type": "application/json",
	}
	var response *events.APIGatewayProxyResponse
	response = &events.APIGatewayProxyResponse{
		StatusCode: 405,
		Body:       "method not implemented",
		Headers:    responseHeaders,
	}
	return response
}

func (handler *CategoryHandler) Post(request events.APIGatewayV2HTTPRequest) (*events.APIGatewayProxyResponse) {
	headers := request.Headers
	isOk, userOrError := validatePostAuthorization(headers)
	if !isOk {
		responseHeaders := map[string]string{
			"Content-Type": "application/json",
		}
		var response *events.APIGatewayProxyResponse
		response = &events.APIGatewayProxyResponse{
			StatusCode: 401,
			Body:       userOrError,
			Headers:    responseHeaders,
		}
		return response
	}
	var newCategory models.Category
	err := json.Unmarshal([]byte(request.Body), &newCategory)
	if err != nil {
		responseHeaders := map[string]string{
			"Content-Type": "application/json",
		}
		var response *events.APIGatewayProxyResponse
		response = &events.APIGatewayProxyResponse{
			StatusCode: 400,
			Body:       "Error in received data " + err.Error(),
			Headers:    responseHeaders,
		}
		return response
	}
	categoryId, err := handler.service.Create(newCategory, userOrError)
	if err != nil {
		responseHeaders := map[string]string{
			"Content-Type": "application/json",
		}
		var response *events.APIGatewayProxyResponse
		response = &events.APIGatewayProxyResponse{
			StatusCode: 400,
			Body:       "Error inserting category " + err.Error(),
			Headers:    responseHeaders,
		}
		return response
	}
	data := map[string]int64{"Categ_Id": categoryId}
	jsonString, err := json.Marshal(data)
	if err != nil {
		responseHeaders := map[string]string{
			"Content-Type": "application/json",
		}
		var response *events.APIGatewayProxyResponse
		response = &events.APIGatewayProxyResponse{
			StatusCode: 400,
			Body:       "Error while generating JSON response: " + err.Error(),
			Headers:    responseHeaders,
		}
		return response
	}
	responseHeaders := map[string]string{
		"Content-Type": "application/json",
	}
	var response *events.APIGatewayProxyResponse
	response = &events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       string(jsonString),
		Headers:    responseHeaders,
	}
	return response
}

func validatePostAuthorization(headers map[string]string) (bool, string) {
	token := headers["authorization"]
	if len(token) == 0 {
		return false, "Token required"
	}
	isOk, msg, _ := auth.ValidateToken(token)
	if !isOk {
		fmt.Println("Error in token: " + msg)
		return false, msg
	}
	fmt.Println("Token OK")
	return true, msg
}

// func validateAuthorization(path string, method string, headers map[string]string) (bool, int, string) {
// 	if (path == "product" && method == "GET") || (path == "category" && method == "GET") {
// 		return true, 200, ""
// 	}

// 	token := headers["authorization"]
// 	if len(token) == 0 {
// 		return false, 401, "Token required"
// 	}

// 	isOk, msg, _ := auth.ValidateToken(token)

// 	if !isOk {
// 		fmt.Println("Error in token: " + msg)
// 		return false, 401, msg
// 	}
// 	fmt.Println("Token OK")
// 	return true, 200, msg
// }
