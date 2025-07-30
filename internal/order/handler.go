package order

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/aws/aws-lambda-go/events"
	"github.com/lubualo/ecommerce-go/authctx"
	"github.com/lubualo/ecommerce-go/models"
	"github.com/lubualo/ecommerce-go/tools"
)

type Handler struct {
	service *Service
}

func NewHandler(service *Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) Post(requestWithContext models.RequestWithContext) *events.APIGatewayProxyResponse {
	body := requestWithContext.RequestBody()

	var o models.Order

	err := json.Unmarshal([]byte(body), &o)
	if err != nil {
		return tools.CreateApiResponse(http.StatusBadRequest, "Invalid JSON body: "+err.Error())
	}

	userUUID, err := authctx.UserUUIDFromContext(requestWithContext.Context())
	if err != nil {
		return tools.CreateApiResponse(http.StatusUnauthorized, "user not found in context")
	}
	o.UserUUID = userUUID
	id, err := h.service.Create(o)
	if err != nil {
		return tools.CreateApiResponse(http.StatusBadRequest, "Error: "+err.Error())
	}

	return tools.CreateApiResponse(http.StatusOK, fmt.Sprintf(`{"orderId": %d}`, id))
}

func (h *Handler) Get(requestWithContext models.RequestWithContext) *events.APIGatewayProxyResponse {
	userUUID, err := authctx.UserUUIDFromContext(requestWithContext.Context())
	if err != nil {
		return tools.CreateApiResponse(http.StatusUnauthorized, "user not found in context")
	}
	query := requestWithContext.RequestQueryStringParameters()

	// === 1. Lookup by ID ===
	if idStr := query["id"]; idStr != "" {
		id, err := strconv.Atoi(idStr)
		if err != nil || id <= 0 {
			return tools.CreateApiResponse(http.StatusBadRequest, "invalid 'id' parameter")
		}
		order, err := h.service.GetById(id)
		if err != nil {
			return tools.CreateApiResponse(http.StatusNotFound, "order not found: "+err.Error())
		}
		body, err := json.Marshal(order)
		if err != nil {
			return tools.CreateApiResponse(http.StatusInternalServerError, "error converting to JSON: "+err.Error())
		}
		return tools.CreateApiResponse(http.StatusOK, string(body))
	}

	// === 2. Default: Get all paginated ===
	page, limit := 1, 10
	from_date, to_date := "", ""

	if val := strings.TrimSpace(query["page"]); val != "" {
		p, err := strconv.Atoi(val)
		if err != nil || p < 1 {
			return tools.CreateApiResponse(http.StatusBadRequest, "invalid 'page' parameter")
		}
		page = p
	}

	if val := strings.TrimSpace(query["from_date"]); val != "" {
		from_date = val
	}

	if val := strings.TrimSpace(query["to_date"]); val != "" {
		to_date = val
	}

	products, err := h.service.GetAllByUserUUID(userUUID, page, limit, from_date, to_date)
	if err != nil {
		return tools.CreateApiResponse(http.StatusInternalServerError, err.Error())
	}
	body, err := json.Marshal(products)
	if err != nil {
		return tools.CreateApiResponse(http.StatusInternalServerError, "error converting to JSON: "+err.Error())
	}
	return tools.CreateApiResponse(http.StatusOK, string(body))
}
