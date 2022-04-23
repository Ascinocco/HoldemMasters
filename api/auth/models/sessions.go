package models

import (
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

func Authenticate(email, password string) SessionResponse {
	user := &User{}
	err := GetDB().Table("users").Where("email = ?", email).First(user).Error

	if err != nil {
		return ErrorResponse
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))

	if err != nil {
		return ErrorResponse
	}

	tokenData := &Token{UserId: user.ID}
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tokenData)
	tokenString, err := token.SignedString([]byte(os.Getenv("token_password")))

	if err != nil {
		return ErrorResponse
	}

	return SessionResponse{
		Token: tokenString,
	}
}
