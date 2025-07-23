package order

import (
	"errors"

	"github.com/lubualo/ecommerce-go/models"
)

type Service struct {
	repo Storage
}

func NewService(repo Storage) *Service {
	return &Service{repo: repo}
}

func (s *Service) Create(o models.Order) (int64, error) {
	if o.UserUUID == "" {
		return 0, ErrMissingUserUUID
	}
	if o.AddressId < 1 {
		return 0, ErrMissingAddress
	}
	if o.Total < 0 {
		return 0, ErrMissingTotal
	}
	count := 0
	for _, detail := range o.Details {
		if detail.ProductId == 0 {
			return 0, ErrMissingProductID
		}
		if detail.Quantity == 0 {
			return 0, ErrMissingQuantity
		}
		if detail.Price == 0 {
			return 0, ErrMissingPrice
		}
		count++
	}
	if count < 1 {
		return 0, ErrNoItems
	}

	return s.repo.Insert(o)
}

func (s *Service) GetById(orderId int) (models.Order, error) {
	return s.repo.GetById(orderId)
}

func (s *Service) GetAllByUserUUID(userUUID string, page int, limit int, from_date string, to_date string) ([]models.Order, error) {
	offset := (page - 1) * limit
	return s.repo.GetAllByUserUUID(userUUID, offset, limit, from_date, to_date)
}

var ErrMissingDate = errors.New("invalid order: date is required")
var ErrMissingTotal = errors.New("invalid order: total is required")
var ErrMissingUserUUID = errors.New("invalid order: user uuid is required")
var ErrMissingAddress = errors.New("invalid order: address is required")
var ErrNoItems = errors.New("invalid order: no items found in order")
var ErrMissingProductID = errors.New("invalid order: product ID is required in order details")
var ErrMissingQuantity = errors.New("invalid order: quantity is required in order details")
var ErrMissingPrice = errors.New("invalid order: price is required in order details")
