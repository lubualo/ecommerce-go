package stock

import (
	"errors"
)

type Service struct {
	repo Storage
}

func NewService(repo Storage) *Service {
	return &Service{repo: repo}
}

func (s *Service) UpdateStock(productId, delta int) error {
	if productId < 1 {
		return ErrInvalidProductId
	}

	return s.repo.UpdateStock(productId, delta)
}

var ErrInvalidProductId = errors.New("invalid product id")
