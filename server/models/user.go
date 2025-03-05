package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Id 		 int32  `json:"id" gorm:"primaryKey"`
	Username string `json:"username" gorm:"unique"`
	Password string `json:"password"`
	Status   string `json:"status"`
	Role     string `json:"role"`
}