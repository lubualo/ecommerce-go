package handlers


import (
	"fmt"

	"github.com/lubualo/ecommerce-go/common"
)

func CreateHandler(entity string) (EntityHandler, error) {
	switch entity {
	case common.User:
		return nil, fmt.Errorf("invalid entity: %s", entity)
	case common.Product:
		return nil, fmt.Errorf("invalid entity: %s", entity)
	case common.Stock:
		return nil, fmt.Errorf("invalid entity: %s", entity)
	case common.Address:
		return nil, fmt.Errorf("invalid entity: %s", entity)
	case common.Category:
		return &CategoryHandler{}, nil
	case common.Order:
		return nil, fmt.Errorf("invalid entity: %s", entity)
	default:
		return nil, fmt.Errorf("invalid entity: %s", entity)
	}

}