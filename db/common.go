package db

import (
	"database/sql"
	"fmt"

	"github.com/Masterminds/squirrel"
	"github.com/ddessilvestri/ecommerce-go/models"
	_ "github.com/go-sql-driver/mysql"
)

// Builds and returns a new DB connection based on the given secret
func DbConnectWithSecret(secret models.SecretRDSJson) (*sql.DB, error) {
	connStr := ConnStr(secret)

	db, err := sql.Open("mysql", connStr)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	fmt.Println("Successfully connected to database")
	return db, nil
}

// Builds a MySQL connection string
func ConnStr(json models.SecretRDSJson) string {
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/gambit?allowCleartextPasswords=true",
		json.Username, json.Password, json.Host)
	fmt.Println("DSN:", dsn)
	return dsn
}

// Verifies if a user is admin
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
