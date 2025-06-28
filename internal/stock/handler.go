package stock

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/aws/aws-lambda-go/events"
	"github.com/lubualo/ecommerce-go/tools"
)

type Handler struct {
	service *Service
}

func NewHandler(service *Service) *Handler {
	return &Handler{service: service}
}

type stockUpdate struct {
	Delta int `json:"delta"`
}

func (h *Handler) Put(request events.APIGatewayV2HTTPRequest) *events.APIGatewayProxyResponse {
	body := request.Body

	var s stockUpdate

	err := json.Unmarshal([]byte(body), &s)
	if err != nil {
		return tools.CreateApiResponse(http.StatusBadRequest, "invalid JSON body: "+err.Error())
	}

	productId, err := strconv.Atoi(request.PathParameters["productId"])
	if err != nil {
		return tools.CreateApiResponse(http.StatusBadRequest, "invalid productId: "+err.Error())
	}

	err = h.service.UpdateStock(productId, s.Delta)
	if err != nil {
		return tools.CreateApiResponse(http.StatusBadRequest, "error: "+err.Error())
	}

	return tools.CreateApiResponse(http.StatusOK, fmt.Sprintf("stock incremented %d for productId %d", s.Delta, productId))
}
