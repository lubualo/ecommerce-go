package category

import (
	"errors"

	"github.com/lubualo/ecommerce-go/models"
)

// Service provides methods for business logic related to category.
type Service struct {
	repo Storage // This is the interface, so it's decoupled from repositorySQL
}

// NewCategoryService acts like a constructor.
func NewCategoryService(repo Storage) *Service {
	return &Service{repo: repo}
}

// CreateCategory performs validation and delegates to the repository.
func (s *Service) CreateCategory(c models.Category) (int64, error) {
	// Simple validation (can be more elaborate in real use cases)
	if c.CategName == "" || c.CategPath == "" {
		return 0, ErrInvalidCategory
	}

	// Business logic: could include auditing, formatting, etc.
	return s.repo.InsertCategory(c)
}

// ErrInvalidCategory represents a validation error.
var ErrInvalidCategory = errors.New("invalid category: name and path are required")
