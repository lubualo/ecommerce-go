package db

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/Masterminds/squirrel"
	"github.com/ddessilvestri/ecommerce-go/models"
	"github.com/ddessilvestri/ecommerce-go/secretm"
	_ "github.com/go-sql-driver/mysql"
)

var SecretModel models.SecretRDSJson
var err error
var Db *sql.DB

func ReadSecret() error {
	SecretModel, err = secretm.GetSecret(os.Getenv("SecretName"))
	return err
}

func DbConnect() error {
	Db, err = sql.Open("mysql", ConnStr(SecretModel))
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	err = Db.Ping()
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	fmt.Println("Successfully database connection")
	return nil
}

func ConnStr(json models.SecretRDSJson) string {
	var dbUser, authToken, dbEndpoint, dbName string
	dbUser = json.Username
	authToken = json.Password
	dbEndpoint = json.Host
	dbName = "gambit"
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?allowCleartextPasswords=true",
		dbUser, authToken, dbEndpoint, dbName)
	fmt.Println(dsn)
	return dsn
}

func UserIsAdmin(userUUID string) (bool, string) {
	fmt.Println("UserIsAdmin starts")

	err := DbConnect()
	if err != nil {
		return false, err.Error()
	}

	defer Db.Close()

	// sentence := "SELECT 1 FROM users WHERE User_UUID'"+userUUID+"' AND User_Status = 0"
	// rows, err := Db.Query(sentence)

	sentence, args, err := squirrel.
		Select("1").
		From("users").
		Where(squirrel.Eq{
			"User_UUID":   userUUID,
			"User_Status": 0,
		}).ToSql()

	if err != nil {
		return false, err.Error()
	}

	fmt.Println(sentence)

	rows, err := Db.Query(sentence, args...)

	if err != nil {
		return false, err.Error()
	}
	var value int

	defer rows.Close()
	if rows.Next() {
		if err := rows.Scan(&value); err != nil {
			return false, err.Error()
		}
		fmt.Println("UserIsAdmin > Successfull execution - value ", value)
		return true, ""
	}

	return false, "User is not Admin"

}
