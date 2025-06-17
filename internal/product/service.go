package product

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

func (s *Service) Create(p models.Product) (int64, error) {
	if p.Title == "" {
		return 0, ErrMissingTitle
	}

	return s.repo.Insert(p)
}

func (s *Service) Update(p models.Product) error {
	if p.Title == "" {
		return ErrMissingTitle
	}
	if p.Id < 1 {
		return ErrInvalidId
	}

	return s.repo.Update(p)
}

func (s *Service) Delete(id int) error {
	return ErrInvalidId
	// if id < 1 {
	// 	return ErrInvalidId
	// }

	// return s.repo.Delete(id)
}

func (s *Service) GetById(id int) (models.Product, error) {
	return models.Product{}, ErrInvalidId
	// if id < 1 {
	// 	return models.Category{}, ErrInvalidId
	// }

	// return s.repo.GetById(id)
}

func (s *Service) GetBySlug(slug string) ([]models.Product, error) {
	return []models.Product{}, ErrInvalidId
	// if slug == "" {
	// 	return []models.Category{}, ErrEmptySlug
	// }
	// return s.repo.GetBySlug(slug)
}

func (s *Service) GetAll() ([]models.Product, error) {
	return []models.Product{}, ErrInvalidId
	// return s.repo.GetAll()
}

var ErrMissingTitle = errors.New("invalid product: title is required")
var ErrInvalidId = errors.New("invalid product id")
var ErrEmptySlug = errors.New("empty product slug")
