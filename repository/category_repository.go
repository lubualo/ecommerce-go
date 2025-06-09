package repository

import (
	"fmt"

	"github.com/Masterminds/squirrel"
	_ "github.com/go-sql-driver/mysql"
	"github.com/lubualo/ecommerce-go/models"
	"github.com/lubualo/ecommerce-go/db"
)

type CategoryRepository struct{}


func (repo *CategoryRepository) Insert(category models.Category, user string) (int64, error) {
	fmt.Println("Category insert - start")

	err := db.DbConnect()
	if err != nil {
		return -1, err
	}
	defer db.Db.Close()

	isAdmin, msg := db.UserIsAdmin(user)
	if !isAdmin {
		return -1, fmt.Errorf("%s",msg)
	}

	query, args, err := squirrel.Insert("category").Columns("Categ_Name", "Categ_Path").Values(category.CategName, category.CategPath).ToSql()
	if err != nil {
		fmt.Println(err.Error())
		return -1, err
	}
	result, err := db.Db.Exec(query, args...)
	if err != nil {
		fmt.Println(err.Error())
		return -1, err
	}
	lastInsertId, err := result.LastInsertId()
	if err != nil {
		return -1, err
	}
	fmt.Println("Category inserted")
	return lastInsertId, nil
}