package category

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

func NewCategoryHandler(service *Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) Post(request events.APIGatewayV2HTTPRequest) (*events.APIGatewayProxyResponse, error) {
	body := request.Body
	
	var c models.Category

	err := json.Unmarshal([]byte(body), &c)
	if err != nil {
		return tools.CreateApiResponse(http.StatusBadRequest, "Invalid JSON body: " + err.Error()), nil
	}

	id, err := h.service.CreateCategory(c)
	if err != nil {
		return tools.CreateApiResponse(http.StatusBadRequest, "Error: " + err.Error()), nil
	}

	return tools.CreateApiResponse(http.StatusOK, fmt.Sprintf(`{"CategID": %d}`, id)), nil
}
