package models

import "gorm.io/gorm"

type Cart struct {
	gorm.Model
	Id    	       int32     `json:"id" gorm:"primaryKey"`
	UserID         uint      `json:"user_id"`
	ProductID      uint      `json:"product_id"`
	Num			   uint      `json:"num"`
	ItemPrice      float64   `json:"item_price"`
}