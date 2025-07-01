package adminusers

import "github.com/lubualo/ecommerce-go/models"

type Storage interface {
	Delete(uuid string) error
	GetAll(page, limit int, sortBy, order string) ([]models.User, error)
}
