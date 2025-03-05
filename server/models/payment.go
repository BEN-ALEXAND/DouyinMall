package models

type PaymentRequest struct {
	OrderID   uint   `json:"order_id" binding:"required"`
	Amount    float64 `json:"amount" binding:"required"`
	PaymentMethod string `json:"payment_method" binding:"required"`
}