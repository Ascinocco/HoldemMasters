package main

import (
	"HoldemMasters/api/auth/controllers"
	"HoldemMasters/api/auth/middleware"
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.Use(middleware.JwtAuthentication)

	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}

	router.HandleFunc("/api/user/new", controllers.CreateAccount).Methods("POST")
	router.HandleFunc("/api/user/login", controllers.Authenticate).Methods("POST")

	router.Handle("/", router)

	err := http.ListenAndServe(":"+port, router)

	if err != nil {
		fmt.Println("Error booting up http server", err)
		os.Exit(1)
	}
}
