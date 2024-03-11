package routes

import (
	"log"
	"moneywaste/package/handlers"
)

func RegisterAuthRoutes() {
	authHandler := handlers.NewAuthHandler()

	log.Printf("GROUP AUTH ROUTES")
	RegisterRoute("/auth/sign-in/", "POST", authHandler.SignIn)
	RegisterRoute("/auth/sign-up/", "POST", authHandler.SignUp)
	RegisterRoute("/auth/logout/", "POST", authHandler.Logout)
	log.Printf("\n\n")
}
