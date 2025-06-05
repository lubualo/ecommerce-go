package db

import (
	"fmt"

	"github.com/Masterminds/squirrel"
	_ "github.com/go-sql-driver/mysql"
	"github.com/lubualo/ecommerce-go/models"
)

func InsertCategory(c models.Category) (int64, error) {
	fmt.Println("Category insert - start")

	err := DbConnect()
	if err != nil {
		return 0, err
	}
	defer Db.Close()

	query, args, err := squirrel.Insert("category").Columns("Categ_Name", "Categ_Path").Values(c.CategName, c.CategPath).ToSql()
	if err != nil {
		fmt.Println(err.Error())
		return 0, err
	}
	result, err := Db.Exec(query, args...)
	if err != nil {
		fmt.Println(err.Error())
		return 0, err
	}
	lastInsertId, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	fmt.Println("Category inserted")
	return lastInsertId, nil
}
