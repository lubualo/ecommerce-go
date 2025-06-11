package category

import (
	"errors"

	"github.com/lubualo/ecommerce-go/models"
)

type Service struct {
	repo Storage
}

func NewCategoryService(repo Storage) *Service {
	return &Service{repo: repo}
}

func (s *Service) CreateCategory(c models.Category) (int64, error) {
	if c.CategName == "" || c.CategPath == "" {
		return 0, ErrMissingNameOrPath
	}

	return s.repo.InsertCategory(c)
}

func (s *Service) UpdateCategory(c models.Category) (error) {
	if c.CategName == "" || c.CategPath == "" {
		return ErrMissingNameOrPath
	}
	if c.CategID < 1 {
		return ErrInvalidId
	}

	return s.repo.UpdateCategory(c)
}

var ErrMissingNameOrPath = errors.New("invalid category: name and path are required")
var ErrInvalidId = errors.New("invalid category id")
