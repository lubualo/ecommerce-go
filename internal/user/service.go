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

func (s *Service) Update(u models.User) error {
	if u.FirstName == "" && u.LastName == "" {
		return ErrMissingNames
	}
	return s.repo.Update(u)
}

func (s *Service) Delete(id int) error {
	if id < 1 {
		return ErrInvalidId
	}

	return s.repo.Delete(id)
}

func (s *Service) GetByUUID(uuid string) (models.User, error) {
	if uuid == "" {
		return models.User{}, ErrInvalidId
	}

	return s.repo.GetByUUID(uuid)
}

var ErrMissingNames = errors.New("invalid user: first name or last name is required")
var ErrInvalidId = errors.New("invalid user id")
