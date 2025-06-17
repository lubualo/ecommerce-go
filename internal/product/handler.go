package product

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

func (h *Handler) Post(request events.APIGatewayV2HTTPRequest) *events.APIGatewayProxyResponse {
	body := request.Body

	var p models.Product

	err := json.Unmarshal([]byte(body), &p)
	if err != nil {
		return tools.CreateApiResponse(http.StatusBadRequest, "Invalid JSON body: "+err.Error())
	}

	id, err := h.service.Create(p)
	if err != nil {
		return tools.CreateApiResponse(http.StatusBadRequest, "Error: "+err.Error())
	}

	return tools.CreateApiResponse(http.StatusOK, fmt.Sprintf(`{"ProdID": %d}`, id))
}

func (h *Handler) Put(request events.APIGatewayV2HTTPRequest) *events.APIGatewayProxyResponse {
	return tools.CreateApiResponse(http.StatusMethodNotAllowed, "Not implemented")
	// body := request.Body

	// var c models.Category

	// err := json.Unmarshal([]byte(body), &c)
	// if err != nil {
	// 	return tools.CreateApiResponse(http.StatusBadRequest, "Invalid JSON body: "+err.Error()), nil
	// }

	// id, err := strconv.Atoi(request.PathParameters["id"])
	// if err != nil {
	// 	return tools.CreateApiResponse(http.StatusBadRequest, "Invalid ID: "+err.Error()), nil
	// }

	// c.CategID = id
	// err = h.service.Update(c)
	// if err != nil {
	// 	return tools.CreateApiResponse(http.StatusBadRequest, "Error: "+err.Error()), nil
	// }

	// return tools.CreateApiResponse(http.StatusOK, fmt.Sprintf("Category updated: %d", id)), nil
}

func (h *Handler) Delete(request events.APIGatewayV2HTTPRequest) *events.APIGatewayProxyResponse {
	return tools.CreateApiResponse(http.StatusMethodNotAllowed, "Not implemented")
	// id, err := strconv.Atoi(request.PathParameters["id"])
	// if err != nil {
	// 	return tools.CreateApiResponse(http.StatusBadRequest, "Invalid ID: "+err.Error()), nil
	// }

	// err = h.service.Delete(id)
	// if err != nil {
	// 	return tools.CreateApiResponse(http.StatusBadRequest, "Error: "+err.Error()), nil
	// }

	// return tools.CreateApiResponse(http.StatusOK, fmt.Sprintf("Category deleted: %d", id)), nil
}

func (h *Handler) Get(request events.APIGatewayV2HTTPRequest) *events.APIGatewayProxyResponse {
	return tools.CreateApiResponse(http.StatusMethodNotAllowed, "Not implemented")

	// var categories []models.Category
	// var err error

	// idStr := request.QueryStringParameters["id"]
	// slug := request.QueryStringParameters["slug"]

	// if idStr != "" {
	// 	var id int
	// 	id, err = strconv.Atoi(idStr)
	// 	if err != nil {
	// 		return tools.CreateApiResponse(http.StatusBadRequest, "Invalid ID: "+err.Error()), nil
	// 	}
	// 	var c models.Category
	// 	c, err = h.service.GetById(id)
	// 	categories = []models.Category{c}
	// } else if slug != "" {
	// 	categories, err = h.service.GetBySlug(slug)
	// } else {
	// 	categories, err = h.service.GetAll()
	// }
	// if err != nil {
	// 	return tools.CreateApiResponse(http.StatusBadRequest, "Error: "+err.Error()), nil
	// }
	// body, err := json.Marshal(categories)
	// if err != nil {
	// 	return tools.CreateApiResponse(http.StatusBadRequest, "Error converting to JSON: "+err.Error()), nil
	// }
	// return tools.CreateApiResponse(http.StatusOK, string(body)), nil
}
