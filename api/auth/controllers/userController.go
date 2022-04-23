package controllers

import (
	"HoldemMasters/api/auth/models"
	"HoldemMasters/api/auth/utils"
	"encoding/json"
	"net/http"
)

type CreateResponse struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	Token    string `json:"token"`
	Error    string `json:"error"`
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	user := &models.User{}

	err := json.NewDecoder(r.Body).Decode(user)
	if err != nil {
		resp := models.UserResponse{
			Error: "Invalid Request",
		}

		w.WriteHeader(http.StatusBadRequest)
		utils.Respond(w, resp)
		return
	}

	uResp, err := user.Create()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		utils.Respond(w, uResp)
		return
	}

	tResp, err := models.Authenticate(user.Email, user.PasswordConfirmation)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		utils.Respond(w, tResp)
		return
	}

	utils.Respond(w, CreateResponse{
		Email:    uResp.Email,
		Username: uResp.Username,
		Token:    tResp.Token,
	})
}
