package category

import "github.com/lubualo/ecommerce-go/models"

type Storage interface {
	Insert(c models.Category) (int64, error)
	Update(c models.Category) (error)
	Delete(id int) (error)
	GetById(id int) (models.Category, error)
	GetBySlug(slug string) ([]models.Category, error)
	GetAll() ([]models.Category, error)
}
