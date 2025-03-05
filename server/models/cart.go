package models

import "gorm.io/gorm"

type Cart struct {
	gorm.Model
	UserID         uint      `json:"user_id" gorm:"primaryKey"`
	ProductID      uint      `json:"product_id"`
	Num			   uint      `json:"num"`
	ItemPrice     float64   `json:"item_price"`
}