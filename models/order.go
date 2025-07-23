package models

type OrdersDetails struct {
	Id        int     `json:"id"`
	OrderId   int     `json:"orderId"`
	ProductId int     `json:"productId"`
	Quantity  int     `json:"quantity"`
	Price     float64 `json:"price"`
}

type Order struct {
	Id        int     `json:"id"`
	UserUUID  string  `json:"userUUID"`
	AddressId int     `json:"addressId"`
	Date      string  `json:"date"`
	Total     float64 `json:"total"`
	Details   []OrdersDetails
}
