package category

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

func (s *Service) Create(c models.Category) (int64, error) {
	if c.CategName == "" || c.CategPath == "" {
		return 0, ErrMissingNameOrPath
	}

	return s.repo.Insert(c)
}

func (s *Service) Update(c models.Category) error {
	if c.CategName == "" || c.CategPath == "" {
		return ErrMissingNameOrPath
	}
	if c.CategID < 1 {
		return ErrInvalidId
	}

	return s.repo.Update(c)
}

func (s *Service) Delete(id int) error {
	if id < 1 {
		return ErrInvalidId
	}

	return s.repo.Delete(id)
}

func (s *Service) GetById(id int) (models.Category, error) {
	if id < 1 {
		return models.Category{}, ErrInvalidId
	}

	return s.repo.GetById(id)
}

func (s *Service) GetBySlug(slug string) ([]models.Category, error) {
	if slug == "" {
		return []models.Category{}, ErrEmptySlug
	}
	return s.repo.GetBySlug(slug)
}

func (s *Service) GetAll() ([]models.Category, error) {
	return s.repo.GetAll()
}

var ErrMissingNameOrPath = errors.New("invalid category: name and path are required")
var ErrInvalidId = errors.New("invalid category id")
var ErrEmptySlug = errors.New("empty category slug")
