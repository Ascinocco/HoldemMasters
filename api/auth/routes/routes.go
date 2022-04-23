package routes

const API_VERSION = "/api/v1"

// @TODO: Add route to get user
var PublicRoutes = map[string]string{
	"CreateUser":    API_VERSION + "/user",
	"CreateSession": API_VERSION + "/session",
}
