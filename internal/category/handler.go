package category

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/lubualo/ecommerce-go/models"
)

// Handler struct wires the service (depends on Service)
type Handler struct {
	service *Service
}

// NewCategoryHandler creates a new handler with injected service
func NewCategoryHandler(service *Service) *Handler {
	return &Handler{service: service}
}

// Post handles the HTTP POST request to create a category
func (h *Handler) Post(body string, user string) (*events.APIGatewayProxyResponse, error) {
	var c models.Category

	// 1. Try to parse the incoming JSON
	err := json.Unmarshal([]byte(body), &c)
	if err != nil {
		return apiResponse(http.StatusBadRequest, "Invalid JSON body: "+err.Error()), nil
	}

	// 2. Call service to create category
	id, err := h.service.CreateCategory(c)
	if err != nil {
		return apiResponse(http.StatusBadRequest, "Error: "+err.Error()), nil
	}

	// 3. Return success response
	return apiResponse(http.StatusOK, fmt.Sprintf(`{"CategID": %d}`, id)), nil
}

// Utility function to standardize API responses
func apiResponse(status int, body string) *events.APIGatewayProxyResponse {
	return &events.APIGatewayProxyResponse{
		StatusCode: status,
		Body:       body,
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
	}
}
