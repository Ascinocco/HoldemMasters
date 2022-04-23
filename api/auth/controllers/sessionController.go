package controllers

import (
	"HoldemMasters/api/auth/models"
	"HoldemMasters/api/auth/utils"
	"encoding/json"
	"net/http"
)

func CreateSession(w http.ResponseWriter, r *http.Request) {
	user := &models.User{}
	err := json.NewDecoder(r.Body).Decode(user)

	if err != nil {
		resp := models.SessionResponse{
			Error: "Invalid Request",
		}

		utils.Respond(w, resp)
		return
	}

	resp := models.Authenticate(user.Email, user.Password)
	utils.Respond(w, resp)
}
