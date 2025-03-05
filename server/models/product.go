package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Id    	 int32   `json:"id" gorm:"primaryKey"`
	Name  	 string  `json:"name"`
	Price 	 float64 `json:"price"`
	Stock 	 int     `json:"stock"`
	ImageUrl string  `json:"image_url"`
}