package controllers

import (
	"HoldemMasters/api/auth/models"
	"HoldemMasters/api/auth/utils"
	"encoding/json"
	"net/http"
)

var CreateAccount = func(w http.ResponseWriter, r *http.Request) {
	user := &models.User{}

	err := json.NewDecoder(r.Body).Decode(user)

	if err != nil {
		utils.Respond(w, utils.Message(false, "Invalid request"))
		return
	}

	resp := user.Create()
	utils.Respond(w, resp)
}

var Authenticate = func(w http.ResponseWriter, r *http.Request) {
	user := &models.User{}
	err := json.NewDecoder(r.Body).Decode(user)
	if err != nil {
		utils.Respond(w, utils.Message(false, "Invalid request"))
		return
	}

	resp := models.Login(user.Email, user.Password)
	utils.Respond(w, resp)
}
