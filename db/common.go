package db

import (
	"database/sql"
	"fmt"
	"os"

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
