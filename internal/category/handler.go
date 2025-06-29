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

func NewHandler(service *Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) Post(requestWithContext models.RequestWithContext) (*events.APIGatewayProxyResponse, error) {
	body := requestWithContext.RequestBody()

	var c models.Category

	err := json.Unmarshal([]byte(body), &c)
	if err != nil {
		return tools.CreateApiResponse(http.StatusBadRequest, "Invalid JSON body: "+err.Error()), nil
	}

	id, err := h.service.Create(c)
	if err != nil {
		return tools.CreateApiResponse(http.StatusBadRequest, "Error: "+err.Error()), nil
	}

	return tools.CreateApiResponse(http.StatusOK, fmt.Sprintf(`{"CategoryId": %d}`, id)), nil
}

func (h *Handler) Put(requestWithContext models.RequestWithContext) (*events.APIGatewayProxyResponse, error) {
	body := requestWithContext.RequestBody()

	var c models.Category

	err := json.Unmarshal([]byte(body), &c)
	if err != nil {
		return tools.CreateApiResponse(http.StatusBadRequest, "Invalid JSON body: "+err.Error()), nil
	}

	id, err := strconv.Atoi(requestWithContext.RequestPathParameters()["id"])
	if err != nil {
		return tools.CreateApiResponse(http.StatusBadRequest, "Invalid ID: "+err.Error()), nil
	}

	c.Id = id
	err = h.service.Update(c)
	if err != nil {
		return tools.CreateApiResponse(http.StatusBadRequest, "Error: "+err.Error()), nil
	}

	return tools.CreateApiResponse(http.StatusOK, fmt.Sprintf("Category updated: %d", id)), nil
}

func (h *Handler) Delete(requestWithContext models.RequestWithContext) (*events.APIGatewayProxyResponse, error) {
	id, err := strconv.Atoi(requestWithContext.RequestPathParameters()["id"])
	if err != nil {
		return tools.CreateApiResponse(http.StatusBadRequest, "Invalid ID: "+err.Error()), nil
	}

	err = h.service.Delete(id)
	if err != nil {
		return tools.CreateApiResponse(http.StatusBadRequest, "Error: "+err.Error()), nil
	}

	return tools.CreateApiResponse(http.StatusOK, fmt.Sprintf("Category deleted: %d", id)), nil
}

func (h *Handler) Get(requestWithContext models.RequestWithContext) (*events.APIGatewayProxyResponse, error) {
	var categories []models.Category
	var err error

	idStr := requestWithContext.RequestQueryStringParameters()["id"]
	slug := requestWithContext.RequestQueryStringParameters()["slug"]

	if idStr != "" {
		var id int
		id, err = strconv.Atoi(idStr)
		if err != nil {
			return tools.CreateApiResponse(http.StatusBadRequest, "Invalid ID: "+err.Error()), nil
		}
		var c models.Category
		c, err = h.service.GetById(id)
		categories = []models.Category{c}
	} else if slug != "" {
		categories, err = h.service.GetBySlug(slug)
	} else {
		categories, err = h.service.GetAll()
	}
	if err != nil {
		return tools.CreateApiResponse(http.StatusBadRequest, "Error: "+err.Error()), nil
	}
	body, err := json.Marshal(categories)
	if err != nil {
		return tools.CreateApiResponse(http.StatusBadRequest, "Error converting to JSON: "+err.Error()), nil
	}
	return tools.CreateApiResponse(http.StatusOK, string(body)), nil
}
