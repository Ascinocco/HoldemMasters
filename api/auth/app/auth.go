package app

import (
	"HoldemMasters/api/auth/utils"
	"context"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/dgrijalva/jwt-go"
)

type Token struct {
	UserId uint
	jwt.StandardClaims
}

var JwtAuthentication = func(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Endpoints that are not authenticated
		unauthenticatedEndpoints := []string{"/api/user/new", "/api/user/login", "/auth-test"}

		// current url
		requestPath := r.URL.Path

		// Check if the requested url is for an un-authed page, if so, forward the request to its handler
		for _, value := range unauthenticatedEndpoints {
			if value == requestPath {
				next.ServeHTTP(w, r)
				return
			}
		}

		response := make(map[string]interface{})
		tokenHeader := r.Header.Get("Authorization")

		// if no auth token is provided, 403 unauth'd
		if tokenHeader == "" {
			response = utils.Message(false, "Missing auth token")
			w.WriteHeader(http.StatusForbidden)
			w.Header().Add("Content-Type", "application/json")
			utils.Respond(w, response)
			return
		}

		// check for malformed token
		// split the token into an array, 0 = "Bearer", 1 = "the-actual-token"
		splitToken := strings.Split(tokenHeader, " ")
		if len(splitToken) != 2 {
			response = utils.Message(false, "Invalid/Malformed auth token")
			w.WriteHeader(http.StatusForbidden)
			w.Header().Add("Content-Type", "application/json")
			utils.Respond(w, response)
			return
		}

		tokenFromClient := splitToken[1]
		tk := &Token{}

		token, err := jwt.ParseWithClaims(tokenFromClient, tk, func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("token_secret")), nil
		})

		if err != nil {
			response = utils.Message(false, "Malformed auth token")
			w.WriteHeader(http.StatusForbidden)
			w.Header().Add("Content-Type", "application/json")
			utils.Respond(w, response)
			return
		}

		if !token.Valid { //Token is invalid, maybe not signed on this server
			response = utils.Message(false, "Token is not valid.")
			w.WriteHeader(http.StatusForbidden)
			w.Header().Add("Content-Type", "application/json")
			utils.Respond(w, response)
			return
		}

		fmt.Println("User", tk.UserId)
		ctx := context.WithValue(r.Context(), "user", tk.UserId)
		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)
	})
}
