package routes

const API_VERSION = "/api/v1"

var PublicRoutes = map[string]string{
	"CreateUser":    API_VERSION + "/user",
	"CreateSession": API_VERSION + "/session",
}
