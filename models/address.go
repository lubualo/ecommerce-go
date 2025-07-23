package models

type Address struct {
	Id         int    `json:"id"`
	Title      string `json:"title"`
	Name       string `json:"name"`
	Address    string `json:"address"`
	City       string `json:"city"`
	State      string `json:"state"`
	PostalCode string `json:"postalCode"`
	Phone      string `json:"phone"`
}
