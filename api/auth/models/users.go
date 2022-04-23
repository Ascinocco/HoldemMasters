package models

import (
	"errors"
	"strings"

	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	gorm.Model
	Email                string `json:"email" gorm:"unique,index"`
	Username             string `json:"username" gorm:"unique"`
	Password             string `json:"password"`
	PasswordConfirmation string `json:"passwordConfirmation" sql:"-"`
	Token                string `json:"token" sql:"-"`
}

type UserResponse struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	Error    string `json:"error"`
}

func (user *User) Validate() (UserResponse, bool) {
	if !strings.Contains(user.Email, "@") {
		return UserResponse{Error: "Invalid email"}, false
	}

	// @TODO: Improve pw validation
	if len(user.Password) < 6 {
		return UserResponse{Error: "Password must be greater than 6 characters"}, false
	}

	if user.Password != user.PasswordConfirmation {
		return UserResponse{Error: "Passwords do not match"}, false
	}

	if len(user.Username) < 3 {
		return UserResponse{Error: "Username must be at least 3 characters"}, false
	}

	tu := &User{}

	err := GetDB().Table("users").Where("email = ?", user.Email).First(tu).Error

	if err != nil && err != gorm.ErrRecordNotFound {
		return UserResponse{Error: "Please try again"}, false
	}

	if tu.Email != "" {
		return UserResponse{Error: "Email cannot be registered"}, false
	}

	return UserResponse{}, true
}

func (user *User) Create() (UserResponse, error) {
	if resp, ok := user.Validate(); !ok {
		return resp, errors.New(resp.Error)
	}

	hashedPw, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	user.Password = string(hashedPw)

	GetDB().Create(user)

	if user.ID <= 0 {
		return UserResponse{Error: "Failed to create user"}, errors.New("Failed to create user")
	}

	return UserResponse{
		Email:    user.Email,
		Username: user.Username,
	}, nil
}
