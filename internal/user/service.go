package user

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
	if p.Name == "" {
		return 0, ErrMissingTitle
	}

	return s.repo.Insert(p)
}

func (s *Service) Update(p models.Product) error {
	if p.Name == "" {
		return ErrMissingTitle
	}
	if p.Id < 1 {
		return ErrInvalidId
	}

	return s.repo.Update(p)
}

func (s *Service) Delete(id int) error {
	if id < 1 {
		return ErrInvalidId
	}

	return s.repo.Delete(id)
}

func (s *Service) GetById(id int) (models.Product, error) {
	if id < 1 {
		return models.Product{}, ErrInvalidId
	}

	return s.repo.GetById(id)
}

func (s *Service) GetBySlug(slug string) (models.Product, error) {
	if slug == "" {
		return models.Product{}, ErrEmptySlug
	}
	return s.repo.GetBySlug(slug)
}

func (s *Service) GetByCategoryId(id int) ([]models.Product, error) {
	if id < 1 {
		return []models.Product{}, ErrInvalidId
	}
	return s.repo.GetByCategoryId(id)
}

func (s *Service) GetByCategorySlug(slug string) ([]models.Product, error) {
	if slug == "" {
		return []models.Product{}, ErrEmptySlug
	}
	return s.repo.GetByCategorySlug(slug)
}

func (s *Service) SearchByText(text string, page, limit int, sortBy, order string) ([]models.Product, error) {
	offset := (page - 1) * limit
	return s.repo.SearchByText(text, offset, limit, sortBy, order)
}

func (s *Service) GetAll(page, limit int, sortBy, order string) ([]models.Product, error) {
	offset := (page - 1) * limit
	return s.repo.GetAll(offset, limit, sortBy, order)
}

var ErrMissingTitle = errors.New("invalid product: title is required")
var ErrInvalidId = errors.New("invalid product id")
var ErrEmptySlug = errors.New("empty product slug")
