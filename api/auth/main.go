package main

import (
	"HoldemMasters/api/auth/app"
	"HoldemMasters/api/auth/controllers"
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.Use(app.JwtAuthentication)

	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}

	router.HandleFunc("/auth-test", controllers.AuthTest).Methods("GET")

	router.Handle("/", router)

	err := http.ListenAndServe(":"+port, router)

	if err != nil {
		fmt.Println("Error booting up http server", err)
		os.Exit(1)
	}
}
