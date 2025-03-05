package user

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email    string `gorm:"unique"`
	Password string
}

type UserService struct {
	db *gorm.DB
}

func NewUserService(db *gorm.DB) *UserService {
	return &UserService{db: db}
}

func (s *UserService) Register(email, password string) (int32, error) {
	user := User{Email: email, Password: password}
	result := s.db.Create(&user)
	return int32(user.ID), result.Error
}

func (s *UserService) Login(email, password string) (int32, error) {
	var user User
	result := s.db.Where("email = ? AND password = ?", email, password).First(&user)
	return int32(user.ID), result.Error
}