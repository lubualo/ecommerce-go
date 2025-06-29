package user

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

func (h *Handler) Put(requestWithContext models.RequestWithContext) *events.APIGatewayProxyResponse {
	body := requestWithContext.RequestBody()

	var p models.Product

	err := json.Unmarshal([]byte(body), &p)
	if err != nil {
		return tools.CreateApiResponse(http.StatusBadRequest, "Invalid JSON body: "+err.Error())
	}

	id, err := strconv.Atoi(requestWithContext.RequestPathParameters()["id"])
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
func (h *Handler) Get(requestWithContext models.RequestWithContext) *events.APIGatewayProxyResponse {
	query := requestWithContext.RequestQueryStringParameters()

	// === 1. Lookup by ID ===
	if idStr := query["id"]; idStr != "" {
		id, err := strconv.Atoi(idStr)
		if err != nil || id <= 0 {
			return tools.CreateApiResponse(http.StatusBadRequest, "invalid 'id' parameter")
		}
		product, err := h.service.GetById(id)
		if err != nil {
			return tools.CreateApiResponse(http.StatusNotFound, "product not found: "+err.Error())
		}
		body, err := json.Marshal(product)
		if err != nil {
			return tools.CreateApiResponse(http.StatusInternalServerError, "error converting to JSON: "+err.Error())
		}
		return tools.CreateApiResponse(http.StatusOK, string(body))
	}

	// === 2. Lookup by Slug ===
	if slug := strings.TrimSpace(query["slug"]); slug != "" {
		product, err := h.service.GetBySlug(slug)
		if err != nil {
			return tools.CreateApiResponse(http.StatusNotFound, "product not found: "+err.Error())
		}
		body, err := json.Marshal(product)
		if err != nil {
			return tools.CreateApiResponse(http.StatusInternalServerError, "error converting to JSON: "+err.Error())
		}
		return tools.CreateApiResponse(http.StatusOK, string(body))
	}

	// === 3. Search full-text ===
	if search := strings.TrimSpace(query["search"]); search != "" {
		page, limit, sortBy, order, err := tools.ParsePaginationAndSorting(query)
		if err != nil {
			return tools.CreateApiResponse(http.StatusBadRequest, err.Error())
		}
		products, err := h.service.SearchByText(search, page, limit, sortBy, order)
		if err != nil {
			return tools.CreateApiResponse(http.StatusInternalServerError, err.Error())
		}
		body, err := json.Marshal(products)
		if err != nil {
			return tools.CreateApiResponse(http.StatusInternalServerError, "error converting to JSON: "+err.Error())
		}
		return tools.CreateApiResponse(http.StatusOK, string(body))
	}

	// === 4. Filter by category ID ===
	if catIdStr := strings.TrimSpace(query["categoryId"]); catIdStr != "" {
		categoryId, err := strconv.Atoi(catIdStr)
		if err != nil || categoryId <= 0 {
			return tools.CreateApiResponse(http.StatusBadRequest, "invalid 'categoryId' parameter")
		}
		products, err := h.service.GetByCategoryId(categoryId)
		if err != nil {
			return tools.CreateApiResponse(http.StatusInternalServerError, err.Error())
		}
		body, err := json.Marshal(products)
		if err != nil {
			return tools.CreateApiResponse(http.StatusInternalServerError, "error converting to JSON: "+err.Error())
		}
		return tools.CreateApiResponse(http.StatusOK, string(body))
	}

	// === 5. Filter by category slug ===
	if slugCateg := strings.TrimSpace(query["slugCateg"]); slugCateg != "" {
		products, err := h.service.GetByCategorySlug(slugCateg)
		if err != nil {
			return tools.CreateApiResponse(http.StatusInternalServerError, err.Error())
		}
		body, err := json.Marshal(products)
		if err != nil {
			return tools.CreateApiResponse(http.StatusInternalServerError, "error converting to JSON: "+err.Error())
		}
		return tools.CreateApiResponse(http.StatusOK, string(body))
	}

	// === 6. Default: Get all paginated ===
	page, limit, sortBy, order, err := tools.ParsePaginationAndSorting(query)
	if err != nil {
		return tools.CreateApiResponse(http.StatusBadRequest, err.Error())
	}

	products, err := h.service.GetAll(page, limit, sortBy, order)
	if err != nil {
		return tools.CreateApiResponse(http.StatusInternalServerError, err.Error())
	}
	body, err := json.Marshal(products)
	if err != nil {
		return tools.CreateApiResponse(http.StatusInternalServerError, "error converting to JSON: "+err.Error())
	}
	return tools.CreateApiResponse(http.StatusOK, string(body))
}
