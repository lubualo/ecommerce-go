package address

import "github.com/lubualo/ecommerce-go/models"

type Storage interface {
	Exists(id int) bool
	Insert(a models.Address, userUUID string) (int64, error)
	Update(a models.Address) error
	Delete(id int) error
	GetAllByUserUUID(userUUID string) ([]models.Address, error)
}
