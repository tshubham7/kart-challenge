package handler

type CreateOrderRequest struct {
	CouponCode string      `json:"couponCode"`
	Items      []OrderItem `json:"items" validate:"gt=0,dive"`
}

type OrderItem struct {
	ProductId string `json:"productId" validate:"required"`
	Quantity  int    `json:"quantity" validate:"min=1"`
}
