package db

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/Masterminds/squirrel"
	_ "github.com/go-sql-driver/mysql"
	"github.com/lubualo/ecommerce-go/models"
)

func DbConnectAndReturn(secret models.RDSCredentials, dbName string) (*sql.DB, error) {
	db, err := sql.Open("mysql", ConnStr(secret, dbName))
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	fmt.Println("Successfully database connection")
	return db, nil
}

func ConnStr(json models.RDSCredentials, dbName string) string {
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?allowCleartextPasswords=true",
		json.Username,
		json.Password,
		json.Host,
		dbName,
	)
	fmt.Println(dsn)
	return dsn
}

func UserIsAdmin(db *sql.DB, userUUID string) (bool, string) {
	fmt.Println("Checking if user is admin:", userUUID)

	query, args, err := squirrel.
		Select("1").
		From("users").
		Where(squirrel.Eq{
			"User_UUID":   userUUID,
			"User_Status": 0,
		}).ToSql()

	if err != nil {
		return false, err.Error()
	}

	rows, err := db.Query(query, args...)
	if err != nil {
		return false, err.Error()
	}
	defer rows.Close()

	var result int
	if rows.Next() {
		if err := rows.Scan(&result); err != nil {
			return false, err.Error()
		}
		fmt.Println("User is admin")
		return true, ""
	}

	return false, "User is not Admin"
}

func EscapeString(s string) string {
	escaped := strings.ReplaceAll(s, "'", "")
	escaped = strings.ReplaceAll(escaped, "\"", "")
	return escaped
}
