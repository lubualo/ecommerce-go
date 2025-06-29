package models

type Product struct {
	Id           int     `json:"prodID"`
	Title        string  `json:"prodTitle"`
	Description  string  `json:"prodDescription"`
	CreatedAt    string  `json:"prodCreatedAt"`
	Updated      string  `json:"prodUpdated"`
	Price        float64 `json:"prodPrice,omitempty"`
	Path         string  `json:"prodPath"`
	Stock        int     `json:"prodStock"`
	CategoryId   int     `json:"prodCategId"`
	CategoryPath string  `json:"categPath,omitempty"`
}
