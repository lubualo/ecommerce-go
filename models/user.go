package models

type User struct {
	UUID      string `json:"uuid"`
	Email     string `json:"email"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Status    int    `json:"status"`
	DateAdd   string `json:"dateAdd"`
	DateUpg   string `json:"dateUpg"`
}
