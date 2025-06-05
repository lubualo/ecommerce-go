package routers

import (
	"encoding/json"
	"fmt"

	"github.com/lubualo/ecommerce-go/db"
	"github.com/lubualo/ecommerce-go/models"
)

func InsertCategory(body string, user string) (int, string) {
	var t models.Category
	err := json.Unmarshal([]byte(body), &t)
	if err != nil {
		return 400, "Error in received data " + err.Error()
	}
	if len(t.CategName) == 0 {
		return 400, "Category name missing"
	}
	if len(t.CategPath) == 0 {
		return 400, "Category path missing"
	}

	isAdmin, msg := db.UserIsAdmin(user)
	if !isAdmin {
		return 400, msg
	}

	result, err := db.InsertCategory(t)
	if err != nil {
		fmt.Println("Error while inserting category " + t.CategName + ": " + err.Error())
		return 400, "Error while inserting category " + t.CategName + ": " + err.Error()
	}

	data := map[string]int64{"Categ_Id": result}
	jsonString, err := json.Marshal(data)
	if err != nil {
		fmt.Println("Error while generating JSON response: " + err.Error())
		return 400, "Error while generating JSON response: " + err.Error()
	}
	return 200, string(jsonString)
}
