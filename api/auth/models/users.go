package models

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	Email    string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
	Token    string `json:"token" sql:"-"`
}
