package db

import (
	"fmt"

	"github.com/Masterminds/squirrel"
	"github.com/ddessilvestri/ecommerce-go/models"
	_ "github.com/go-sql-driver/mysql"
)

func InsertCategory(c models.Category) (int64, error) {
	fmt.Println("Category register starting... ")
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

	// var result sql.Result
	// result, err = Db.Exec(query, args)
	result, err := Db.Exec(query, args...)

	if err != nil {
		fmt.Println(err.Error())
		return 0, err
	}

	LastInsertId, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	fmt.Println("Insert Category > Successfull Execution")
	return LastInsertId, nil

}
