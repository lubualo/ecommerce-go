package product

import "github.com/lubualo/ecommerce-go/models"

type Storage interface {
	Insert(p models.Product) (int64, error)
	Update(p models.Product) error
	Delete(id int) error
	GetById(id int) (models.Product, error)
	GetBySlug(slug string) (models.Product, error)
	GetByCategoryId(id int) ([]models.Product, error)
	GetByCategorySlug(slug string) ([]models.Product, error)
	SearchByText(text string, page, limit int, sortBy, order string) ([]models.Product, error)
	GetAll(page, limit int, sortBy, order string) ([]models.Product, error)
}
