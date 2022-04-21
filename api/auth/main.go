package main

import (
	"HoldemMasters/api/auth/app"
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

	err := http.ListenAndServe(":"+port, router)

	if err != nil {
		fmt.Println("Error booting up http server", err)
		os.Exit(1)
	}

	fmt.Println("listening on ", port)
}
