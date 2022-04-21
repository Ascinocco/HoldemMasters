package controllers

import (
	"HoldemMasters/api/auth/utils"
	"fmt"
	"net/http"
)

func AuthTest(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Running")
	response := utils.Message(true, "Healthy")
	w.WriteHeader(http.StatusOK)
	w.Header().Add("Content-Type", "application/json")
	utils.Respond(w, response)
}
