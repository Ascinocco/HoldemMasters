package models

import (
	"errors"
	"os"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

type Token struct {
	UserId uint
	jwt.StandardClaims
}

type SessionResponse struct {
	Token string `json:"token"`
	Error string `json:"error"`
}

var ErrorResponse = SessionResponse{
	Error: "Unable to create session, please try again.",
}

func Authenticate(email, password string) (SessionResponse, error) {
	user := &User{}
	err := GetDB().Table("users").Where("email = ?", email).First(user).Error

	if err != nil {
		return ErrorResponse, errors.New(ErrorResponse.Error)
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))

	if err != nil {
		return ErrorResponse, errors.New(ErrorResponse.Error)
	}

	tokenData := &Token{UserId: user.ID}
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tokenData)
	tokenString, err := token.SignedString([]byte(os.Getenv("token_password")))

	if err != nil {
		return ErrorResponse, errors.New(ErrorResponse.Error)
	}

	return SessionResponse{
		Token: tokenString,
	}, nil
}
