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

		w.WriteHeader(http.StatusBadRequest)
		utils.Respond(w, resp)
		return
	}

	resp, err := models.Authenticate(user.Email, user.Password)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		w.WriteHeader(http.StatusOK)
	}

	utils.Respond(w, resp)
}
