package user

import "github.com/lubualo/ecommerce-go/models"

type Storage interface {
	Update(p models.User) error
	Delete(id int) error
	GetByUUID(uuid string) (models.User, error)
}
