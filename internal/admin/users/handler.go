package adminusers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/aws/aws-lambda-go/events"

	"github.com/lubualo/ecommerce-go/models"
	"github.com/lubualo/ecommerce-go/tools"
)

type Handler struct {
	service *Service
}

func NewHandler(service *Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) Get(requestWithContext models.RequestWithContext) *events.APIGatewayProxyResponse {
// Verify if the user is admin

	query := requestWithContext.RequestQueryStringParameters()
	page, limit, sortBy, order, err := tools.ParsePaginationAndSorting(query)
	if err != nil {
		return tools.CreateApiResponse(http.StatusBadRequest, err.Error())
	}

	users, err := h.service.GetAll(page, limit, sortBy, order)
	if err != nil {
		return tools.CreateApiResponse(http.StatusInternalServerError, err.Error())
	}
	body, err := json.Marshal(users)
	if err != nil {
		return tools.CreateApiResponse(http.StatusInternalServerError, "error converting to JSON: "+err.Error())
	}
	return tools.CreateApiResponse(http.StatusOK, string(body))
}

func (h *Handler) Delete(requestWithContext models.RequestWithContext) *events.APIGatewayProxyResponse {
// Verify if the user is admin

	uuid := requestWithContext.RequestPathParameters()["id"]
	if uuid == "" {
		return tools.CreateApiResponse(http.StatusBadRequest, "invalid UUID")
	}

	err := h.service.Delete(uuid)
	if err != nil {
		return tools.CreateApiResponse(http.StatusBadRequest, "error: "+err.Error())
	}

	return tools.CreateApiResponse(http.StatusOK, fmt.Sprintf("product deleted: %s", uuid))
}

