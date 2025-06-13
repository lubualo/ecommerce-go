package category

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

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
		return tools.CreateApiResponse(http.StatusBadRequest, "Invalid JSON body: "+err.Error()), nil
	}

	id, err := h.service.Create(c)
	if err != nil {
		return tools.CreateApiResponse(http.StatusBadRequest, "Error: "+err.Error()), nil
	}

	return tools.CreateApiResponse(http.StatusOK, fmt.Sprintf(`{"CategID": %d}`, id)), nil
}

func (h *Handler) Put(request events.APIGatewayV2HTTPRequest) (*events.APIGatewayProxyResponse, error) {
	body := request.Body

	var c models.Category

	err := json.Unmarshal([]byte(body), &c)
	if err != nil {
		return tools.CreateApiResponse(http.StatusBadRequest, "Invalid JSON body: "+err.Error()), nil
	}

	id, err := strconv.Atoi(request.PathParameters["id"])
	if err != nil {
		return tools.CreateApiResponse(http.StatusBadRequest, "Invalid ID: "+err.Error()), nil
	}

	c.CategID = id
	err = h.service.Update(c)
	if err != nil {
		return tools.CreateApiResponse(http.StatusBadRequest, "Error: "+err.Error()), nil
	}

	return tools.CreateApiResponse(http.StatusOK, fmt.Sprintf("Category updated: %d", id)), nil
}

func (h *Handler) Delete(request events.APIGatewayV2HTTPRequest) (*events.APIGatewayProxyResponse, error) {
	id, err := strconv.Atoi(request.PathParameters["id"])
	if err != nil {
		return tools.CreateApiResponse(http.StatusBadRequest, "Invalid ID: "+err.Error()), nil
	}

	err = h.service.Delete(id)
	if err != nil {
		return tools.CreateApiResponse(http.StatusBadRequest, "Error: "+err.Error()), nil
	}

	return tools.CreateApiResponse(http.StatusOK, fmt.Sprintf("Category deleted: %d", id)), nil
}

func (h *Handler) Get(request events.APIGatewayV2HTTPRequest) (*events.APIGatewayProxyResponse, error) {
	idStr := request.QueryStringParameters["id"]
	if idStr != "" {
		id, err := strconv.Atoi(idStr)
		if err != nil {
			return tools.CreateApiResponse(http.StatusBadRequest, "Invalid ID: "+err.Error()), nil
		}

		c, err := h.service.GetById(id)
		if err != nil {
			return tools.CreateApiResponse(http.StatusBadRequest, "Error: "+err.Error()), nil
		}

		body, err := json.Marshal(c)
		if err != nil {
			return tools.CreateApiResponse(http.StatusBadRequest, "Error converting model to JSON: "+err.Error()), nil
		}

		return tools.CreateApiResponse(http.StatusOK, string(body)), nil
	}
	categories, err := h.service.GetAll()
	if err != nil {
		return tools.CreateApiResponse(http.StatusBadRequest, "Error: "+err.Error()), nil
	}

	body, err := json.Marshal(categories)
	if err != nil {
		return tools.CreateApiResponse(http.StatusInternalServerError, "Error converting models to JSON: "+err.Error()), nil
	}

	return tools.CreateApiResponse(http.StatusOK, string(body)), nil
}
