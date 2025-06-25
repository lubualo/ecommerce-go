package product

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"

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
	body := request.Body

	var p models.Product

	err := json.Unmarshal([]byte(body), &p)
	if err != nil {
		return tools.CreateApiResponse(http.StatusBadRequest, "Invalid JSON body: "+err.Error())
	}

	id, err := strconv.Atoi(request.PathParameters["id"])
	if err != nil {
		return tools.CreateApiResponse(http.StatusBadRequest, "Invalid ID: "+err.Error())
	}

	p.Id = id
	err = h.service.Update(p)
	if err != nil {
		return tools.CreateApiResponse(http.StatusBadRequest, "Error: "+err.Error())
	}

	return tools.CreateApiResponse(http.StatusOK, fmt.Sprintf("Product updated: %d", id))
}

func (h *Handler) Delete(request events.APIGatewayV2HTTPRequest) *events.APIGatewayProxyResponse {
	id, err := strconv.Atoi(request.PathParameters["id"])
	if err != nil {
		return tools.CreateApiResponse(http.StatusBadRequest, "Invalid ID: "+err.Error())
	}

	err = h.service.Delete(id)
	if err != nil {
		return tools.CreateApiResponse(http.StatusBadRequest, "Error: "+err.Error())
	}

	return tools.CreateApiResponse(http.StatusOK, fmt.Sprintf("Product deleted: %d", id))
}

func (h *Handler) Get(request events.APIGatewayV2HTTPRequest) *events.APIGatewayProxyResponse {
		return tools.CreateApiResponse(http.StatusMethodNotAllowed, "Not implemented")
	// var categories []models.Category
	// var err error
    // query := request.QueryStringParameters

	// idStr := query["id"]
	// if idStr != "" {
    // 	id, err := strconv.Atoi(idStr)
	// 	if err != nil || id <= 0 {
	// 		return tools.CreateApiResponse(http.StatusBadRequest, "Invalid 'id' parameter")
	// 	}
	// 	product, err := h.service.GetById(id)
	// 	if err != nil {
	// 		return tools.CreateApiResponse(http.StatusNotFound, "Product not found: " + err.Error())
	// 	}
	// 	body, err := json.Marshal(product)
	// 	if err != nil {
	// 		return tools.CreateApiResponse(http.StatusBadRequest, "Error converting to JSON: "+err.Error())
	// 	}
	// 	return tools.CreateApiResponse(http.StatusOK, string(body))
	// }

    // page := 1 // Default value
	// value := query["page"]
	// if value != "" {
    //     parsed, err := strconv.Atoi(value)
    //     if err != nil || parsed < 1 {
    //         return tools.CreateApiResponse(http.StatusBadRequest, "Invalid 'page' parameter")
    //     }
	// 	page = parsed
	// }

    // limit := 10 // Default value
	// value = query["limit"]
	// if value != "" {
    //     parsed, err := strconv.Atoi(value)
    //     if err != nil || parsed < 1 {
    //         return tools.CreateApiResponse(http.StatusBadRequest, "Invalid 'limit' parameter")
    //     }
    //     limit = parsed
    // }

    // sortBy := "id" // Default value
	// value = query["sort_by"]
	// if value != "" {
    //     allowed := map[string]bool{
    //         "id": true,
    //         "title": true,
    //         "description": true,
    //         "price": true,
    //         "category_id": true,
    //         "stock": true,
    //         "created_at": true,
    //     }
    //     if !allowed[value] {
    //         return tools.CreateApiResponse(http.StatusBadRequest, "Invalid 'sort_by' parameter")
    //     }
    //     sortBy = value
    // }
	
    // order := "ASC" // Default value
	// value = query["order"]
	// if value != "" {
    //     upperCase := strings.ToUpper(value)
    //     if upperCase != "ASC" && upperCase != "DESC" {
    //         return tools.CreateApiResponse(http.StatusBadRequest, "Invalid 'order' parameter")
    //     }
    //     order = upperCase
    // }

    // products, err := h.service.GetAll(page, limit, sortBy, order)
    // if err != nil {
    //     return tools.CreateApiResponse(http.StatusInternalServerError, err.Error())
    // }

	// body, err := json.Marshal(products)
	// if err != nil {
	// 	return tools.CreateApiResponse(http.StatusBadRequest, "Error converting to JSON: "+err.Error())
	// }
	// return tools.CreateApiResponse(http.StatusOK, string(body))



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
