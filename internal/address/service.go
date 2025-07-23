package address

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

func (s *Service) Create(a models.Address, userUUID string) (int64, error) {
	if a.Address == "" {
		return 0, ErrMissingAddress
	}
	if a.Name == "" {
		return 0, ErrMissingName
	}
	if a.Title == "" {
		return 0, ErrMissingTitle
	}
	if a.City == "" {
		return 0, ErrMissingCity
	}
	if a.Phone == "" {
		return 0, ErrMissingPhone
	}
	if a.PostalCode == "" {
		return 0, ErrMissingPostalCode
	}
	if userUUID == "" {
		return 0, ErrInvalidUserUUID
	}
	return s.repo.Insert(a, userUUID)
}

func (s *Service) Update(a models.Address) error {
	if a.Id < 1 {
		return ErrInvalidId
	}
	if !s.repo.Exists(a.Id) {
		return ErrAddressNotFound
	}
	if a.Address == "" {
		return ErrMissingAddress
	}
	if a.Name == "" {
		return ErrMissingName
	}
	if a.Title == "" {
		return ErrMissingTitle
	}
	if a.City == "" {
		return ErrMissingCity
	}
	if a.Phone == "" {
		return ErrMissingPhone
	}
	if a.PostalCode == "" {
		return ErrMissingPostalCode
	}

	return s.repo.Update(a)
}

func (s *Service) Delete(id int) error {
	if id < 1 {
		return ErrInvalidId
	}
	if !s.repo.Exists(id) {
		return ErrAddressNotFound
	}
	return s.repo.Delete(id)
}

func (s *Service) GetAllByUserUUID(userUUID string) ([]models.Address, error) {
	return s.repo.GetAllByUserUUID(userUUID)
}

var ErrInvalidId = errors.New("invalid address: id is required")
var ErrMissingAddress = errors.New("invalid address: address is required")
var ErrMissingName = errors.New("invalid address: name is required")
var ErrMissingTitle = errors.New("invalid address: title is required")
var ErrMissingCity = errors.New("invalid address: city is required")
var ErrMissingPhone = errors.New("invalid address: phone is required")
var ErrMissingPostalCode = errors.New("invalid address: postal code is required")
var ErrInvalidUserUUID = errors.New("invalid user uuid")
var ErrAddressNotFound = errors.New("address not found")
