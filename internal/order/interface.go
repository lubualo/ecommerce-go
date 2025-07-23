package order

import "github.com/lubualo/ecommerce-go/models"

type Storage interface {
	Insert(o models.Order) (int64, error)
	GetById(id int) (models.Order, error)
	GetAllByUserUUID(userUUID string, offset, limit int, sortBy, order string) ([]models.Order, error)
}
