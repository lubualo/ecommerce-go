package address

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

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

	var a models.Address

	err := json.Unmarshal([]byte(body), &a)
	if err != nil {
		return tools.CreateApiResponse(http.StatusBadRequest, "Invalid JSON body: "+err.Error())
	}

	userUUID, err := authctx.UserUUIDFromContext(requestWithContext.Context())
	if err != nil {
		return tools.CreateApiResponse(http.StatusUnauthorized, "user not found in context")
	}

	id, err := h.service.Create(a, userUUID)
	if err != nil {
		return tools.CreateApiResponse(http.StatusBadRequest, "Error: "+err.Error())
	}

	return tools.CreateApiResponse(http.StatusOK, fmt.Sprintf(`{"addressId": %d}`, id))
}

func (h *Handler) Put(requestWithContext models.RequestWithContext) *events.APIGatewayProxyResponse {
	id, err := strconv.Atoi(requestWithContext.RequestPathParameters()["id"])
	if err != nil {
		return tools.CreateApiResponse(http.StatusBadRequest, "invalid ID: "+err.Error())
	}

	body := requestWithContext.RequestBody()

	var a models.Address

	err = json.Unmarshal([]byte(body), &a)
	if err != nil {
		return tools.CreateApiResponse(http.StatusBadRequest, "invalid JSON body: "+err.Error())
	}

	a.Id = id
	err = h.service.Update(a)
	if err != nil {
		return tools.CreateApiResponse(http.StatusBadRequest, "error: "+err.Error())
	}

	return tools.CreateApiResponse(http.StatusOK, fmt.Sprintf("Address updated: %d", id))
}

func (h *Handler) Delete(requestWithContext models.RequestWithContext) *events.APIGatewayProxyResponse {
	id, err := strconv.Atoi(requestWithContext.RequestPathParameters()["id"])
	if err != nil {
		return tools.CreateApiResponse(http.StatusBadRequest, "invalid ID: "+err.Error())
	}

	err = h.service.Delete(id)
	if err != nil {
		return tools.CreateApiResponse(http.StatusBadRequest, "error: "+err.Error())
	}

	return tools.CreateApiResponse(http.StatusOK, fmt.Sprintf("address deleted: %d", id))
}

func (h *Handler) Get(requestWithContext models.RequestWithContext) *events.APIGatewayProxyResponse {
	userUUID, err := authctx.UserUUIDFromContext(requestWithContext.Context())
	if err != nil {
		return tools.CreateApiResponse(http.StatusUnauthorized, "user not found in context")
	}
	addresses, err := h.service.GetAllByUserUUID(userUUID)
	if err != nil {
		return tools.CreateApiResponse(http.StatusInternalServerError, err.Error())
	}
	body, err := json.Marshal(addresses)
	if err != nil {
		return tools.CreateApiResponse(http.StatusInternalServerError, "error converting to JSON: "+err.Error())
	}
	return tools.CreateApiResponse(http.StatusOK, string(body))
}


