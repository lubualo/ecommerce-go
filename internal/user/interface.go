package user

import "github.com/lubualo/ecommerce-go/models"

type Storage interface {
	Update(p models.User) error
	GetByUUID(uuid string) (models.User, error)
}
