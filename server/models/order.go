package models

import "gorm.io/gorm"

type Order struct {
	gorm.Model
	OrderID        uint      `json:"order_id" gorm:"primaryKey"`
	UserID         uint      `json:"user_id"`
	ProductID      uint      `json:"product_id"`
	ProductName    string    `json:"product_name"`
	Num            int       `json:"num"`
	TotalPrice     float64   `json:"total_price"`
	Status         string    `json:"status"`
}