package models

import (
	"HoldemMasters/api/auth/utils"
	"os"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

type Token struct {
	UserId uint
	jwt.StandardClaims
}
type User struct {
	gorm.Model
	Email                string `json:"email" gorm:"unique,index"`
	Username             string `json:"username" gorm:"unique"`
	Password             string `json:"password"`
	PasswordConfirmation string `json:"passwordConfirmation" sql:"-"`
	Token                string `json:"token" sql:"-"`
}

func (user *User) Validate() (map[string]interface{}, bool) {
	if !strings.Contains(user.Email, "@") {
		return utils.Message(false, "Email address is required"), false
	}

	if len(user.Password) < 6 {
		return utils.Message(false, "Password must have length greater than 6, contain 1 number and 1 symbol"), false
	}

	if user.Password != user.PasswordConfirmation {
		return utils.Message(false, "Passwords do not match"), false
	}

	if len(user.Username) < 3 {
		return utils.Message(false, "Username must be at least 3 characters"), false
	}

	tu := &User{}

	err := GetDB().Table("users").Where("email = ?", user.Email).First(tu).Error

	if err != nil && err != gorm.ErrRecordNotFound {
		return utils.Message(false, "Connection error. Please retry."), false
	}

	if tu.Email != "" {
		return utils.Message(false, "Email address cannot be registered."), false
	}

	return utils.Message(false, "Success validating user"), true
}

func (user *User) Create() map[string]interface{} {
	if resp, ok := user.Validate(); !ok {
		return resp
	}

	hashedPw, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	user.Password = string(hashedPw)

	GetDB().Create(user)

	if user.ID <= 0 {
		return utils.Message(false, "Failed to create account, connection error.")
	}

	tk := &Token{UserId: user.ID}
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)
	tokenString, _ := token.SignedString([]byte(os.Getenv("token_secret")))
	user.Token = tokenString
	user.Password = ""
	response := utils.Message(true, "Account has been created")
	response["user"] = user

	return response
}

func Login(email, password string) map[string]interface{} {
	user := &User{}
	err := GetDB().Table("users").Where("email = ?", email).First(user).Error

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return utils.Message(false, "Email address not found")
		}
		return utils.Message(false, "Connection error. Please retry")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))

	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword { //Password does not match!
		return utils.Message(false, "Invalid login credentials. Please try again")
	}

	user.Password = ""

	tk := &Token{UserId: user.ID}
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)
	tokenString, _ := token.SignedString([]byte(os.Getenv("token_password")))
	user.Token = tokenString //Store the token in the response

	resp := utils.Message(true, "Logged In")
	resp["user"] = user
	return resp
}

func GetUser(id uint) *User {
	user := &User{}
	GetDB().Table("users").Where("id = ?", id).First(user)
	if user.Email == "" { //User not found!
		return nil
	}

	user.Password = ""
	return user
}
