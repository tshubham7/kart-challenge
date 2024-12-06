package model

type Product struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	Price    int    `json:"price"`
	Category string `json:"category"`
}

type Order struct {
	Id             string      `json:"id"`
	CouponDiscount int         `json:"-"`
	Items          []OrderItem `json:"items"`
	Products       []Product   `json:"products"`
}

type OrderItem struct {
	ProductId string `json:"productId"`
	Quantity  int    `json:"quantity"`
}
