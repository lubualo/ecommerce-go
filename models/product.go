package models

type Product struct {
	Id           int     `json:"id"`
	Name         string  `json:"name"`
	Description  string  `json:"description"`
	CreatedAt    string  `json:"createdAt"`
	Updated      string  `json:"updated"`
	Price        float64 `json:"price,omitempty"`
	Path         string  `json:"path"`
	Stock        int     `json:"stock"`
	CategoryId   int     `json:"categoryId"`
	CategoryPath string  `json:"categoryPath,omitempty"`
}
