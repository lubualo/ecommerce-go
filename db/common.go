package db

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/Masterminds/squirrel"
	_ "github.com/go-sql-driver/mysql"
	"github.com/lubualo/ecommerce-go/models"
	"github.com/lubualo/ecommerce-go/secretm"
)

var SecretModel models.SecretRDSJson
var err error
var Db *sql.DB

func ReadSecret() error {
	SecretModel, err = secretm.GetSecret(os.Getenv("SecretName"))
	/** escriba en la db usando el secreto y escribiendo la info del user q se registró*/
	return err
}

func DbConnect() error {
	Db, err = sql.Open("mysql", ConnStr(SecretModel))
	if err != nil {
		fmt.Print(err.Error())
		return err
	}

	err = Db.Ping()
	if err != nil {
		fmt.Print(err.Error())
		return err
	}

	fmt.Println("Connection to DB successful")
	return nil
}

func ConnStr(json models.SecretRDSJson) string {
	var dbUser, authToken, dbEndpoint, dbName string
	dbUser = json.Username
	authToken = json.Password
	dbEndpoint = json.Host
	dbName = "gambit"
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s)/%s?allowCleartextPasswords=true",
		dbUser,
		authToken,
		dbEndpoint,
		dbName,
	)
	fmt.Println(dsn)
	return dsn
}

func UserIsAdmin(UserUUID string) (bool, string) {
	fmt.Println("UserIsAdmin process begin")
	err := DbConnect()
	if err != nil {
		fmt.Println("DB connection failed: " + err.Error())
		return false, err.Error()
	}
	defer Db.Close()

	query, args, err := squirrel.Select("1").
		From("users").
		Where(squirrel.Eq{"User_UUID": UserUUID, "User_Status": 0}).
		ToSql()
	if err != nil {
		fmt.Println("Query build failed: " + err.Error())
		return false, err.Error()
	}
	fmt.Println("Query: " + query)

	rows := Db.QueryRow(query, args...)

	var value int
	if err := rows.Scan(&value); err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("User not found or inactive: " + err.Error())
			return false, err.Error()
		}
		fmt.Println("Query execution error: " + err.Error())
		return false, "Query execution error: " + err.Error()
	}
	fmt.Println("UserIsAdmin value:", value)
	return true, ""
}
